package server

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	config2 "studyRoomGo/common/config"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware/metadata"

	"io"
	http2 "net/http"
	"strings"
	myJwt "studyRoomGo/common/jwt"
	"studyRoomGo/common/wechat"
	"studyRoomGo/internal/conf"
	checkLoginEquipment "studyRoomGo/internal/middleware/checkLoginEquipment"
	"studyRoomGo/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
	jwtV4 "github.com/golang-jwt/jwt/v4"
	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, s *service.Service, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			metadata.Server(),
			selector.Server(
				jwt.Server(func(token *jwtV4.Token) (interface{}, error) {
					if config2.AppConf.Environment != "pro" {
						return myJwt.MySecretDev, nil
					}
					return myJwt.MySecret, nil
				}),
				checkLoginEquipment.Server(),
			).Match(NewCheckListMatcher()).Build(),
		),
		http.RequestDecoder(NewRequestDecoder),
		http.ErrorEncoder(NewErrorEncoder),

		//http.Filter(
		//	//跨域  推荐
		//	handlers.CORS(
		//		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		//		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"}),
		//		handlers.AllowedOrigins([]string{"*"}),
		//	)),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}

	srv := http.NewServer(opts...)

	h := openapiv2.NewHandler()
	//将/q/路由放在最前匹配
	srv.HandlePrefix("/q/", h)

	RegisterHTTPServer(srv, s)
	return srv
}

// NewCheckListMatcher 需要检查登录状态的接口
func NewCheckListMatcher() selector.MatchFunc {
	return func(ctx context.Context, operation string) bool {
		for _, prefix := range jwtApiPrefixCheckList {
			if strings.HasPrefix(operation, prefix) {
				if _, ok := jwtUrlIgnoreList[operation]; ok {
					return false
				}
				return true
			}
		}

		if _, ok := jwtUrlCheckList[operation]; ok {
			return true
		}
		return false
	}
}

// NewRequestDecoder decodes the request body to object.
func NewRequestDecoder(r *http.Request, v interface{}) error {
	wxSerial := r.Header.Get("Wechatpay-Serial")
	if wxSerial != "" { //微信支付回调通知解析
		err := WeXinPayDecoder(r, v)
		if err != nil {
			return errors.BadRequest("CODEC", err.Error())
		}
		return nil
	}
	//
	codec, ok := http.CodecForRequest(r, "Content-Type")
	if !ok {
		return errors.BadRequest("CODEC", fmt.Sprintf("unregister Content-Type: %s", r.Header.Get("Content-Type")))
	}
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return errors.BadRequest("CODEC", err.Error())
	}
	if len(data) == 0 {
		return nil
	}
	if err = codec.Unmarshal(data, v); err != nil {
		return errors.BadRequest("CODEC", fmt.Sprintf("body unmarshal %s", err.Error()))
	}
	return nil
}

func WeXinPayDecoder(r *http.Request, v interface{}) error {
	splitUrl := strings.Split(r.URL.Path, "/")
	if strings.Index(splitUrl[3], "m_") != 0 {
		return errors.BadRequest("CODEC", fmt.Sprintf("url error"))
	}
	m := strings.Split(splitUrl[3], "_")
	mId, _ := strconv.Atoi(m[1])
	merchantId := int64(mId)

	headerByte, _ := json.Marshal(r.Header)
	var noticeReply *notify.Request
	noticeReply, err := wechat.NoticeReply(merchantId, r) //微信支付解析
	bodyByte, _ := io.ReadAll(r.Body)
	wechat.RecordPayNoticeLog("Request", headerByte, bodyByte)
	if err != nil {
		wechat.RecordPayNoticeLog("NoticeReplyErr", headerByte, []byte(err.Error()))
		return err
	}
	replyByte, _ := json.Marshal(noticeReply)
	wechat.RecordPayNoticeLog("ParseRequest", headerByte, replyByte)

	if noticeReply != nil && noticeReply.Resource != nil { //微信支付解析
		data := []byte(noticeReply.Resource.Plaintext)
		wechat.RecordPayNoticeLog("ParsePlaintext", headerByte, data)
		if len(data) == 0 {
			return nil
		}
		codec, ok := http.CodecForRequest(r, "Content-Type")
		if !ok {
			return errors.BadRequest("CODEC", fmt.Sprintf("unregister Content-Type: %s", r.Header.Get("Content-Type")))
		}
		if err = codec.Unmarshal(data, v); err != nil {
			return errors.BadRequest("CODEC", fmt.Sprintf("body unmarshal %s", err.Error()))
		}
	}
	return nil
}

// NewErrorEncoder encodes the error to the HTTP response.
func NewErrorEncoder(w http.ResponseWriter, r *http.Request, err error) {
	wxSerial := r.Header.Get("Wechatpay-Serial")
	if wxSerial != "" { //微信支付回调通知解析
		WxPaymentNoticeErrorEncoder(w, r, err)
		return
	}

	se := errors.FromError(err)
	codec, _ := http.CodecForRequest(r, "Accept")
	body, err := codec.Marshal(se)
	if err != nil {
		w.WriteHeader(http2.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", ContentType(codec.Name()))
	w.WriteHeader(int(se.Code))
	_, _ = w.Write(body)
}

func WxPaymentNoticeErrorEncoder(w http.ResponseWriter, r *http.Request, err error) {
	se := errors.FromError(err)
	newError := &struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{
		Code:    "FAIL",
		Message: se.Message,
	}
	codec, _ := http.CodecForRequest(r, "Accept")
	body, err := codec.Marshal(newError)
	if err != nil {
		w.WriteHeader(http2.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", ContentType(codec.Name()))
	w.WriteHeader(int(se.Code))
	_, _ = w.Write(body)
}

const (
	baseContentType = "application"
)

// ContentType returns the content-type with base prefix.
func ContentType(subtype string) string {
	return strings.Join([]string{baseContentType, subtype}, "/")
}
