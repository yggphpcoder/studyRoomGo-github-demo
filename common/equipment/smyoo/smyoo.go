package smyoo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	config2 "studyRoomGo/common/config"
)

// 闪优物联接口

var (
	clientID     string = "" //配置闪优物联信息
	clientSecret string = ""
	deviceid     string = ""

	phone    string = ""
	password string = ""

	url string = "https://auth.smyoo.com" // 闪优物联接口地址

	appId      int = 0
	endpointOS     = 1
	autoLogin      = "true"
	areaid         = 0
)

const (
	API_SYNLOGINOPEN     = "/v1/account/synloginopen"
	API_SYNLOGINTICKET   = "/api/gfriend/synloginticket"
	API_SETCHANNELDATA   = "/api/gfriend/setchanneldata"
	API_QUERYDEVICES     = "/api/gfriend/querydevices"
	API_SETMULTICHANNELS = "/api/gfriend/setmultichannels"
)

type Option func(o *Options)

type Options struct {
	Key string
}

func WithKey(key string) Option {
	return func(o *Options) {
		o.Key = key
	}
}

type Server struct {
	Option       Options
	initClient   *smyoo
	limitRefresh int
	refreshCount int
}

func NewServer(opts ...Option) *Server {
	options := Options{
		Key: "smyoo",
	}
	for _, o := range opts {
		o(&options)
	}
	return &Server{
		Option:       options,
		limitRefresh: 3,
	}
}
func (s *Server) GetKey() string {
	return s.Option.Key
}

func (s *Server) Setup() {
	smyoo := NewSmyoo()
	ticket, err := smyoo.GetTicket()
	if err != nil {
		fmt.Println(err)
	}
	bsId, err := smyoo.GetBpeSessionId(ticket)
	if err != nil {
		fmt.Println(err)
	}
	smyoo.BpeSessionId = bsId
	s.initClient = smyoo
}

func (s *Server) Close() {
}

func (s *Server) Automate(mac string, code string, status int) (ok bool, err error) {
	if config2.AppConf.Environment == "pro" {
		index, _ := strconv.Atoi(code)
		if statusCode, err := s.initClient.SetChannelData(mac, index, status); statusCode != 200 {
			if statusCode == -10297010 && s.refreshCount < s.limitRefresh {
				s.refreshCount++
				s.refreshBpeSessionId()

				s.Automate(mac, code, status)
			}
			fmt.Println(statusCode, err)
			return false, err
		}
		s.refreshCount = 0
	} else {
		fmt.Println("not send open door...")
		return true, nil
	}

	return true, nil
}

func (s *Server) AutomateMain(mac string, code []string, status int) (ok bool, err error) {
	if config2.AppConf.Environment == "pro" {
		for _, c := range code {
			index, _ := strconv.Atoi(c)
			if statusCode, err := s.initClient.SetChannelData(mac, index, status); statusCode != 200 {
				if statusCode == -10297010 && s.refreshCount < s.limitRefresh {
					s.refreshCount++
					s.refreshBpeSessionId()

					s.AutomateMain(mac, code, status)
				}
				fmt.Println(statusCode, err)
				return false, err
			}
			s.refreshCount = 0
		}

	} else {
		fmt.Println("not send open door...")
		return true, nil
	}

	return true, nil

}
func (s *Server) refreshBpeSessionId() (bool, error) {
	ticket, err := s.initClient.GetTicket()
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	bsId, err := s.initClient.GetBpeSessionId(ticket)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	s.initClient.BpeSessionId = bsId
	return true, nil
}

type smyoo struct {
	BpeSessionId string
	Ticket       string

	BaseReq BaseReq
}

type BaseReq struct {
	Appid      int    `json:"appId"`
	Areaid     int    `json:"areaId"`
	EndpointOS int    `json:"endpointOS"`
	Deviceid   string `json:"deviceId"`
}
type BaseRes struct {
	ResultCode int    `json:"resultCode"`
	ResultMsg  string `json:"resultMsg"`
}
type GetTicketReq struct {
	Clientid     string `json:"client_id"`
	Clientsecret string `json:"client_secret"`
	Phone        string `json:"phone"`
	Password     string `json:"password"`

	AutoLogin string `json:"autologin"`

	BaseReq
}

type GetTicketRep struct {
	BaseRes
	Data struct {
		CookieExpireTime int    `json:"cookieExpireTime"`
		NextAction       int    `json:"nextAction"`
		Ticket           string `json:"ticket"`
	} `json:"data"`
}

func NewSmyoo() *smyoo {
	return &smyoo{
		BaseReq: BaseReq{
			Appid:      appId,
			Areaid:     areaid,
			EndpointOS: endpointOS,
			Deviceid:   deviceid,
		},
	}
}

func (s *smyoo) GetTicket() (string, error) {
	//http请求获取ticket
	//请求参数
	req := GetTicketReq{
		Clientid:     clientID,
		Clientsecret: clientSecret,
		Phone:        phone,
		Password:     password,
		AutoLogin:    autoLogin,
		BaseReq:      s.BaseReq,
	}

	rep, err := PostJSON(url+API_SYNLOGINOPEN, req)
	if err != nil {
		fmt.Println(err)
	}
	var repData GetTicketRep
	json.Unmarshal(rep, &repData)
	// fmt.Println(repData)

	if repData.ResultCode != 0 {
		fmt.Println("获取ticket失败")
		return "", fmt.Errorf("获取ticket失败")
	}
	return repData.Data.Ticket, nil
}

type GetBpeSessionIdReq struct {
	Ticket string `json:"ticket"`

	BaseReq
}

type GetBpeSessionIdRep struct {
	BaseRes
	Data struct {
		BpeSessionId string `json:"bpeSessionId"`
	} `json:"data"`
}

func (s *smyoo) GetBpeSessionId(ticket string) (string, error) {
	//http请求获取BpeSessionId
	//请求参数
	req := GetBpeSessionIdReq{
		Ticket:  ticket,
		BaseReq: s.BaseReq,
	}
	rep, err := PostJSON(url+API_SYNLOGINTICKET, req)
	if err != nil {
		fmt.Println(err)
	}
	var repData GetBpeSessionIdRep
	json.Unmarshal(rep, &repData)
	// fmt.Println(repData)

	if repData.ResultCode != 0 {
		fmt.Println("获取BpeSessionId失败")
		return "", fmt.Errorf("获取BpeSessionId失败")
	}
	return repData.Data.BpeSessionId, nil

}

type SetChannelDataReq struct {
	McuId     string `json:"mcuid"`
	Datatype  int    `json:"datatype"`
	Datapoint string `json:"datapoint"`

	BaseReq
}

type Datapoint struct {
	Index   int `json:"index"`
	Status  int `json:"status"`
	Poweron int `json:"poweron"`
}

func (s *smyoo) SetChannelData(mcuId string, index int, status int) (code int, err error) {
	//http请求设置通道数据
	//请求参数
	Datapoint := Datapoint{
		Index:  index,
		Status: status,
	}
	bt, _ := json.Marshal(Datapoint)
	req := SetChannelDataReq{
		McuId:     mcuId,
		Datatype:  1,
		Datapoint: string(bt),
		BaseReq:   s.BaseReq,
	}
	ctx := context.Background()
	byteData, _ := json.Marshal(req)

	header := make(map[string]string)
	header["Content-Type"] = "application/json; charset=UTF-8"

	cookies := make(map[string]string)
	cookies["BpeSessionId"] = s.BpeSessionId
	rep, err := HTTPPostContext(ctx, url+API_SETCHANNELDATA, byteData, header, cookies)
	if err != nil {
		fmt.Println(err)
	}
	var repData BaseRes
	json.Unmarshal(rep, &repData)
	// fmt.Println(repData)

	if repData.ResultCode != 0 {

		fmt.Println("设置通道数据失败")
		return repData.ResultCode, fmt.Errorf("设置通道数据失败")
	}
	return 200, nil
}

type QuerydevicesDataRep struct {
	BaseRes
	Data struct {
		Mcuinfos string `json:"mcuinfos"`
	} `json:"data"`
}

type Mcuinfos struct {
	Mcuname string `json:"mcuname"`
	Note    string `json:"note"`
	Type    int    `json:"type"`

	Datatype   int    `json:"datatype"`
	Channelnum int    `json:"channelnum"`
	Datapoint  string `json:"datapoint"`
}

func (s *smyoo) Querydevices() (bool, error) {
	//http请求设置通道数据
	//请求参数

	ctx := context.Background()
	byteData, _ := json.Marshal(s.BaseReq)

	header := make(map[string]string)
	header["Content-Type"] = "application/json; charset=UTF-8"

	cookies := make(map[string]string)
	cookies["BpeSessionId"] = s.BpeSessionId
	rep, err := HTTPPostContext(ctx, url+API_QUERYDEVICES, byteData, header, cookies)
	if err != nil {
		fmt.Println(err)
	}
	var repData QuerydevicesDataRep
	json.Unmarshal(rep, &repData)
	// fmt.Println(repData)
	// fmt.Println(string(rep))

	var jsonData []Mcuinfos
	_ = json.Unmarshal([]byte(repData.Data.Mcuinfos), &jsonData)
	fmt.Println(jsonData)

	if repData.ResultCode != 0 {
		fmt.Println("查询通道数据失败")
		return false, fmt.Errorf("查询通道数据失败")
	}
	return true, nil
}

// HTTPPostContext post 请求
func HTTPPostContext(ctx context.Context, uri string, data []byte, header map[string]string, cookies map[string]string) ([]byte, error) {
	body := bytes.NewBuffer(data)
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, uri, body)
	if err != nil {
		return nil, err
	}

	for key, value := range header {
		request.Header.Set(key, value)
	}

	for key, value := range cookies {
		c := &http.Cookie{Name: key, Value: value}
		request.AddCookie(c)

	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http post error : uri=%v , statusCode=%v", uri, response.StatusCode)
	}
	return io.ReadAll(response.Body)
}

// PostJSON post json 数据请求
func PostJSON(uri string, obj interface{}) ([]byte, error) {
	jsonBuf := new(bytes.Buffer)
	enc := json.NewEncoder(jsonBuf)
	enc.SetEscapeHTML(false)
	err := enc.Encode(obj)
	if err != nil {
		return nil, err
	}
	response, err := http.Post(uri, "application/json;charset=utf-8", jsonBuf)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get error : uri=%v , statusCode=%v", uri, response.StatusCode)
	}
	return io.ReadAll(response.Body)
}
