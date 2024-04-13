package wechat

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	config2 "studyRoomGo/common/config"
	"studyRoomGo/common/wechat/gen"
	"time"

	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/downloader"
	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	partnerpayments "github.com/wechatpay-apiv3/wechatpay-go/services/partnerpayments"
	partner_jsapi "github.com/wechatpay-apiv3/wechatpay-go/services/partnerpayments/jsapi"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
	"github.com/wechatpay-apiv3/wechatpay-go/services/profitsharing"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

var Merchant = make(map[int64]MerchantInfo)
var Shop = make(map[int64]ShopInfo)

type MerchantInfo struct {
	Appid           string
	Mchid           string
	v3Secret        string
	Cert            string
	IsProfitSharing bool
	pemPath         string
}
type ShopInfo struct {
	SubMchid string
	SubAppid string
}

func init() {
	LoadMerchantConf()
}

type WXPayClient struct {
	ctx    context.Context
	client *core.Client
}

func LoadMerchantConf() {
	Merchant[1] = MerchantInfo{ //微信支付商户号信息
		Appid:           "",
		Mchid:           "",
		v3Secret:        "",
		Cert:            "",
		IsProfitSharing: false, //是否分账开关,
		pemPath:         config2.AppConf.Config.Dir + "/m1_client_key.pem",
	}
	Merchant[3] = MerchantInfo{
		Appid:           "",
		Mchid:           "",
		v3Secret:        "",
		Cert:            "",
		IsProfitSharing: false, //是否分账开关,
		pemPath:         config2.AppConf.Config.Dir + "/partner_apiclient_key.pem",
	}
	Shop[3] = ShopInfo{
		SubAppid: "",
		SubMchid: "",
	}
	Shop[4] = ShopInfo{
		SubAppid: "",
		SubMchid: "",
	}
}
func WXPayClientInit(merchantId int64) *WXPayClient {
	LoadMerchantConf()
	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath(Merchant[merchantId].pemPath)
	if err != nil {
		log.Printf("load merchant private key error")
	}

	wxCtx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(Merchant[merchantId].Mchid, Merchant[merchantId].Cert, mchPrivateKey, Merchant[merchantId].v3Secret),
	}
	client, err := core.NewClient(wxCtx, opts...)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
	}
	return &WXPayClient{
		ctx:    wxCtx,
		client: client,
	}

}

type Attach struct {
	MerchantId int64
	ShopId     int64
}
type JSApi struct {
	MerchantId  int64
	ShopId      int64
	MemberId    int64
	Description string
	OutTradeNo  string
	Attach      string
	TotalPrice  int64
}

func (r *JSApi) Request() (*jsapi.PrepayWithRequestPaymentResponse, error) {
	wxInfo, err := GetUserInfo(r.MemberId)
	if err != nil {
		return nil, err
	}
	wxClient := WXPayClientInit(r.MerchantId)

	apiUrl := "https://api.fish-study.com/wx_api"
	if config2.AppConf.Environment != "pro" {
		apiUrl = "https://api.fish-study.com/wx_api_oplat"
	}

	svc := jsapi.JsapiApiService{Client: wxClient.client}
	// 得到prepay_id，以及调起支付所需的参数和签名
	resp, _, err := svc.PrepayWithRequestPayment(wxClient.ctx,
		jsapi.PrepayRequest{
			Appid:       core.String(Merchant[r.MerchantId].Appid),
			Mchid:       core.String(Merchant[r.MerchantId].Mchid),
			Description: core.String(r.Description),
			OutTradeNo:  core.String(r.OutTradeNo),
			Attach:      core.String(r.Attach),
			NotifyUrl:   core.String(fmt.Sprintf("%s/payment/v1/m_%d/wxpay_notice", apiUrl, r.MerchantId)),
			Amount: &jsapi.Amount{
				Total: core.Int64(r.TotalPrice),
			},
			Payer: &jsapi.Payer{
				Openid: core.String(wxInfo.Openid),
			},
			SettleInfo: &jsapi.SettleInfo{
				ProfitSharing: func() *bool {
					if Merchant[r.MerchantId].IsProfitSharing {
						return core.Bool(true)
					}
					return core.Bool(false)
				}(),
			},
		},
	)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *JSApi) PartnerRequest() (*jsapi.PrepayWithRequestPaymentResponse, error) {
	wxInfo, err := GetUserInfo(r.MemberId)
	if err != nil {
		return nil, err
	}
	wxClient := WXPayClientInit(r.MerchantId)

	apiUrl := "https://api.fish-study.com/wx_api"
	if config2.AppConf.Environment != "pro" {
		apiUrl = "https://api.fish-study.com/wx_api_oplat"
	}
	svc := partner_jsapi.JsapiApiService{Client: wxClient.client}
	// 得到prepay_id，以及调起支付所需的参数和签名

	resp, _, err := svc.Prepay(wxClient.ctx,
		partner_jsapi.PrepayRequest{
			SpAppid:     core.String(Merchant[r.MerchantId].Appid),
			SpMchid:     core.String(Merchant[r.MerchantId].Mchid),
			SubAppid:    core.String(Shop[r.ShopId].SubAppid),
			SubMchid:    core.String(Shop[r.ShopId].SubMchid),
			Description: core.String(r.Description),
			OutTradeNo:  core.String(r.OutTradeNo),
			Attach:      core.String(r.Attach),
			NotifyUrl:   core.String(fmt.Sprintf("%s/payment/v1/m_%d/wxpay_notice", apiUrl, r.MerchantId)),
			Amount: &partner_jsapi.Amount{
				Total: core.Int64(r.TotalPrice),
			},
			Payer: &partner_jsapi.Payer{
				SubOpenid: core.String(wxInfo.Openid),
			},
			SettleInfo: &partner_jsapi.SettleInfo{
				ProfitSharing: func() *bool {
					if Merchant[r.MerchantId].IsProfitSharing {
						return core.Bool(true)
					}
					return core.Bool(false)
				}(),
			},
		},
	)
	if err != nil {
		return nil, err
	}
	newResp := &jsapi.PrepayWithRequestPaymentResponse{}
	newResp.PrepayId = resp.PrepayId
	newResp.SignType = core.String("RSA")
	newResp.Appid = core.String(Shop[r.ShopId].SubAppid)
	newResp.TimeStamp = core.String(strconv.FormatInt(time.Now().Unix(), 10))
	nonce, err := utils.GenerateNonce()
	if err != nil {
		return nil, fmt.Errorf("generate request for payment err:%s", err.Error())
	}
	newResp.NonceStr = core.String(nonce)
	newResp.Package = core.String("prepay_id=" + *resp.PrepayId)
	message := fmt.Sprintf("%s\n%s\n%s\n%s\n", *newResp.Appid, *newResp.TimeStamp, *newResp.NonceStr, *newResp.Package)
	signatureResult, err := wxClient.client.Sign(context.Background(), message)
	if err != nil {
		return nil, fmt.Errorf("generate sign for payment err:%s", err.Error())
	}
	newResp.PaySign = core.String(signatureResult.Signature)
	return newResp, nil
}

// NoticeReply 微信支付 通知解析
func NoticeReply(merchantId int64, request *http.Request) (req *notify.Request, err error) {
	LoadMerchantConf()
	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath(Merchant[merchantId].pemPath)
	if err != nil {
		log.Printf("load merchant private key error")
	}
	ctx := context.Background()
	// 1. 使用 `RegisterDownloaderWithPrivateKey` 注册下载器
	err = downloader.MgrInstance().RegisterDownloaderWithPrivateKey(ctx, mchPrivateKey, Merchant[merchantId].Cert, Merchant[merchantId].Mchid, Merchant[merchantId].v3Secret)
	// 2. 获取商户号对应的微信支付平台证书访问器
	certificateVisitor := downloader.MgrInstance().GetCertificateVisitor(Merchant[merchantId].Mchid)
	// 3. 使用证书访问器初始化 `notify.Handler`
	handler, _ := notify.NewRSANotifyHandler(Merchant[merchantId].v3Secret, verifiers.NewSHA256WithRSAVerifier(certificateVisitor))
	notifyReq := &notify.Request{}
	if merchantId == 3 {
		transaction := new(partnerpayments.Transaction)
		notifyReq, err = handler.ParseNotifyRequest(context.Background(), request, transaction)

	} else {
		transaction := new(payments.Transaction)
		notifyReq, err = handler.ParseNotifyRequest(context.Background(), request, transaction)

	}
	// 如果验签未通过，或者解密失败
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return notifyReq, nil
}

type ProfitSharing struct {
	ShopId          int64
	MerchantId      int64
	OrderNo         string
	OrderDesc       string
	TransactionId   string
	UnfreezeUnsplit int8
	Amount          *ProfitSharingOrderAmount
	Receivers       *ProfitSharingOrderReceivers
}
type ProfitSharingOrderAmount struct {
	TotalAmount int
	Ratio       int8
	Amount      int
}

type ProfitSharingOrderReceivers struct {
	ReceiversType    string
	ReceiversAccount string
	ReceiversName    string
}

// 分账 创建订单
func (r *ProfitSharing) ProfitSharingCreateOrder() (resp *profitsharing.OrdersEntity, result *core.APIResult, err error) {
	wxClient := WXPayClientInit(r.MerchantId)
	unfreezeUnsplit := false
	if r.UnfreezeUnsplit == 1 {
		unfreezeUnsplit = true
	}
	svc := profitsharing.OrdersApiService{Client: wxClient.client}
	resp, result, err = svc.CreateOrder(wxClient.ctx,
		profitsharing.CreateOrderRequest{
			Appid:      core.String(Merchant[r.MerchantId].Appid),
			OutOrderNo: core.String(r.OrderNo),
			Receivers: []profitsharing.CreateOrderReceiver{profitsharing.CreateOrderReceiver{
				Account:     core.String(r.Receivers.ReceiversAccount),
				Amount:      core.Int64(int64(r.Amount.Amount)),
				Description: core.String(r.OrderDesc),
				Name:        core.String(r.Receivers.ReceiversName),
				Type:        core.String(r.Receivers.ReceiversType),
			}},
			TransactionId:   core.String(r.TransactionId),
			UnfreezeUnsplit: core.Bool(unfreezeUnsplit),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call CreateOrder err:%s", err)
		return resp, result, err
	}
	if result.Response.StatusCode != 200 {
		return resp, result, errors.New("分账失败")
	}
	return resp, result, err
}

func (r *ProfitSharing) ProfitSharingGetOrder(orderNo string, transactionId string) (resp *profitsharing.OrdersEntity, result *core.APIResult, err error) {
	wxClient := WXPayClientInit(r.MerchantId)
	svc := profitsharing.OrdersApiService{Client: wxClient.client}
	resp, result, err = svc.QueryOrder(wxClient.ctx, profitsharing.QueryOrderRequest{OutOrderNo: &orderNo, TransactionId: &transactionId})
	if err != nil {
		// 处理错误
		log.Printf("call CreateOrder err:%s", err)
		return resp, result, err
	}
	if result.Response.StatusCode != 200 {
		return resp, result, errors.New("分账查询失败")
	}
	return resp, result, err
}
func RecordPayNoticeLog(typ string, header, body []byte) {
	mgr := gen.WxPayNoticeLogMgr(config2.AppConf.Data.DB)

	h := string(header)
	b := string(body)
	mgr.Save(&gen.WxPayNoticeLog{
		Type:   typ,
		Header: h,
		Body:   b,
	})
}

func GetUserInfo(memberId int64) (info *gen.WxUserinfo, err error) {
	mgr := gen.WxUserinfoMgr(config2.AppConf.Data.DB)
	get, err := mgr.GetFromMemberID(memberId)
	if err != nil {
		return nil, err
	}
	return &get, nil
}
