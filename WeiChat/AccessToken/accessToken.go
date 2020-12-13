package AccessToken

import (
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"time"
	"zhome-server/WeiChat"
	"zhome-server/WeiChat/WeiChatConfig"
)

type AccessToken struct {
	accessToken string
	grant_type  string
	url         string
	appid       string
	secret      string
}

func (a *AccessToken) Init() error {
	a.grant_type = "client_credential"
	a.url = "https://api.weixin.qq.com/cgi-bin/token"
	WeiChat.
		GetRegisterWeiChat().
		RequireWeiChatFunction(
			reflect.TypeOf(new(WeiChatConfig.WeiChatConfig)),
			func(i interface{}) error {
				id, ok := i.(*WeiChatConfig.WeiChatConfig)
				if !ok {
					return errors.New("Need *AppID.AppId")
				}
				a.appid = id.GetAppID()
				a.secret = id.GetSecret()
				err := a.StartTimer()
				if err != nil {
					log.Error(err.Error())
					return err
				}
				return nil
			})
	return nil
}

func (a *AccessToken) GetAccessToken() (int, error) {
	type weixinResult struct {
		access_token string
		expires_in   int
		errcode      int
		errmsg       string
	}
	baseURL, err := url.Parse(a.url)
	if err != nil {
		return 0, err
	}

	params := url.Values{}
	params.Add("grant_type", a.grant_type)
	params.Add("appid", a.appid)
	params.Add("secret", a.secret)
	baseURL.RawQuery = params.Encode()
	result, err := http.Get(baseURL.String())
	if err != nil {
		return 0, err
	}
	defer func() {
		if result != nil {
			err := result.Body.Close()
			if err != nil {
				log.Error(err.Error())
			}
		}
	}()
	if result.StatusCode != 200 {
		log.Error("Get weichat access token error")
		return 0, errors.New("Get weichat access token error")
	}
	body, err := ioutil.ReadAll(result.Body)
	var r weixinResult
	err = json.Unmarshal(body, &r)
	if err != nil {
		return 0, err
	}
	if r.errcode != 0 {
		log.Error(r.errmsg)
		return 0, errors.New(r.errmsg)
	}
	a.accessToken = r.access_token
	return r.expires_in, nil
}
func (a *AccessToken) Stop() error {
	return nil
}
func (a *AccessToken) StartTimer() error {
	go func() {
		for {
			t, err := a.GetAccessToken()
			if err != nil {
				time.Sleep(time.Second * 10)
			}
			time.Sleep(time.Second * time.Duration(t))
		}
	}()
	return nil
}

func init() {
	WeiChat.GetRegisterWeiChat().RegisterWeiChatFunction(new(AccessToken))
}
