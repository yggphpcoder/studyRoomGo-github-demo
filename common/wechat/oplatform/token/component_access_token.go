package token

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type AccessTokenServer struct {
	ComponentAccessToken string `json:"component_access_token"`
	ExpiresIn            int64  `json:"expires_in"` // 当前时间 + 过期时间
}

type responseData struct {
	Data struct {
		Token string
	}
}

func (d *AccessTokenServer) Token() (token string, err error) {
	response, err := http.Get("http://10.0.4.9:8081/inner/component-access-token")
	if err != nil {
		log.Println("ioutil read error")
		return "", nil
	}

	defer func(response *http.Response) {
		if response != nil && response.Body != nil {
			err := response.Body.Close()
			if err != nil {
				return
			}
		}

	}(response)
	body, err2 := ioutil.ReadAll(response.Body)
	if err2 != nil {
		log.Println("ioutil read error")
		return "", nil
	}
	var j responseData
	err = json.Unmarshal(body, &j)
	if err != nil {
		return "", err
	}
	return j.Data.Token, nil
}

func (d *AccessTokenServer) GetToken() (token string, expiresIn int64, err error) {
	if token, err = d.Token(); err == nil {
		expiresIn = d.ExpiresIn
	}

	return
}

func (d *AccessTokenServer) RestoreToken(token string, expiresIn int64) (err error) {
	d.ExpiresIn = expiresIn
	d.ComponentAccessToken = token
	return
}
