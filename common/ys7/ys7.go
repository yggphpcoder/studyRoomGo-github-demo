package ys7

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var AccessToken string
var ExpireTime int64

type Auth struct {
	AppKey string
	Secret string
}

// 萤石云监控配置
func NewAuth() *Auth {
	auth := &Auth{
		AppKey: "",
		Secret: "",
	}
	return auth
}

type AccessTokenResp struct {
	Code string
	Msg  string
	Data struct {
		AccessToken string
		ExpireTime  int64
	}
}

func (a *Auth) GetAccessToken() (string, error) {
	now := time.Now().UnixMilli()
	if now > ExpireTime {
		data := url.Values{}
		data.Set("appKey", a.AppKey)
		data.Set("appSecret", a.Secret)
		pryload := strings.NewReader(data.Encode())
		httpResp, err := http.Post("https://open.ys7.com/api/lapp/token/get", "application/x-www-form-urlencoded", pryload)
		if err != nil {
			return "", err
		}
		res := &AccessTokenResp{}
		defer httpResp.Body.Close()

		if httpResp.StatusCode != http.StatusOK {
			return "", errors.New("http.Status:" + httpResp.Status)
		}
		err = json.NewDecoder(httpResp.Body).Decode(res)
		if err != nil {
			return "", err
		}
		AccessToken = res.Data.AccessToken
		ExpireTime = res.Data.ExpireTime
	}
	return AccessToken, nil
}

type Device struct {
	DeviceSerial string
}

type DeviceListResp struct {
	Data []struct {
		DeviceSerial  string
		DeviceName    string
		DeviceType    string
		Status        int
		Defence       int
		DeviceVersion string
	}
}

func (a *Device) DeviceList() DeviceListResp {
	auth := NewAuth()
	data := url.Values{}
	token, err := auth.GetAccessToken()
	if err != nil {
		return DeviceListResp{}
	}
	data.Set("accessToken", token)
	pryload := strings.NewReader(data.Encode())
	httpResp, err := http.Post("https://open.ys7.com/api/lapp/device/list", "application/x-www-form-urlencoded", pryload)
	if err != nil {
		return DeviceListResp{}
	}
	res := &DeviceListResp{}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return DeviceListResp{}
	}
	err = json.NewDecoder(httpResp.Body).Decode(res)
	if err != nil {
		return DeviceListResp{}
	}
	return *res
}

type CameraListResp struct {
	Data []struct {
		DeviceSerial string
		ChannelNo    int
		ChannelName  string
		Status       int
		IsShared     string
		PicUrl       string
		IsEncrypt    int
		VideoLevel   int
		Permission   int
	}
}

func (a *Device) CameraList() CameraListResp {
	auth := NewAuth()
	data := url.Values{}
	token, err := auth.GetAccessToken()
	if err != nil {
		return CameraListResp{}
	}
	data.Set("accessToken", token)
	data.Set("deviceSerial", a.DeviceSerial)
	pryload := strings.NewReader(data.Encode())
	httpResp, err := http.Post("https://open.ys7.com/api/lapp/device/camera/list", "application/x-www-form-urlencoded", pryload)
	if err != nil {
		return CameraListResp{}
	}
	res := &CameraListResp{}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return CameraListResp{}
	}
	err = json.NewDecoder(httpResp.Body).Decode(res)
	if err != nil {
		return CameraListResp{}
	}
	return *res
}

type LiveAddressListResp struct {
	Code string
	Msg  string
	Data []struct {
		DeviceSerial string
		ChannelNo    int
		DeviceName   string
		LiveAddress  string
		HdAddress    string
		BeginTime    int
		EndTime      int
	}
}

func (a *Device) LiveAddressList() LiveAddressListResp {
	auth := NewAuth()
	data := url.Values{}
	token, err := auth.GetAccessToken()
	if err != nil {
		return LiveAddressListResp{}
	}
	data.Set("accessToken", token)
	pryload := strings.NewReader(data.Encode())
	httpResp, err := http.Post("https://open.ys7.com/api/lapp/live/video/list", "application/x-www-form-urlencoded", pryload)
	if err != nil {
		return LiveAddressListResp{}
	}
	res := &LiveAddressListResp{}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return LiveAddressListResp{}
	}
	err = json.NewDecoder(httpResp.Body).Decode(res)
	if err != nil {
		return LiveAddressListResp{}
	}
	return *res
}

type LiveAddressResp struct {
	Code string
	Msg  string
	Data struct {
		DeviceSerial string
		ChannelNo    int
		DeviceName   string
		LiveAddress  string
		HdAddress    string
		BeginTime    int
		EndTime      int
	}
}

func (a *Device) LiveAddress(deviceSerial string, channelNo string, expireTime string) LiveAddressResp {
	auth := NewAuth()
	data := url.Values{}
	token, err := auth.GetAccessToken()
	if err != nil {
		return LiveAddressResp{}
	}
	data.Set("accessToken", token)
	data.Set("deviceSerial", deviceSerial)
	data.Set("channelNo", channelNo)
	data.Set("expireTime", expireTime)
	pryload := strings.NewReader(data.Encode())
	httpResp, err := http.Post("https://open.ys7.com/api/lapp/live/address/limited", "application/x-www-form-urlencoded", pryload)
	if err != nil {
		return LiveAddressResp{}
	}
	res := &LiveAddressResp{}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return LiveAddressResp{}
	}
	err = json.NewDecoder(httpResp.Body).Decode(res)
	if err != nil {
		return LiveAddressResp{}
	}
	return *res
}
