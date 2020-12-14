package Code2Session

import (
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"zhome-server/WeiChat"
	"zhome-server/WeiChat/AccessToken"
	"zhome-server/WeiChat/WeiChatConfig"
)

type Code2Session struct {
	grant_type string
	url        string
	appid      string
	appSecret  *AccessToken.AccessToken
}

func (a *Code2Session) Init() error {
	a.grant_type = "authorization_code"
	a.url = "https://api.weixin.qq.com/sns/jscode2session"
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
				return nil
			})
	WeiChat.
		GetRegisterWeiChat().
		RequireWeiChatFunction(
			reflect.TypeOf(new(AccessToken.AccessToken)),
			func(i interface{}) error {
				id, ok := i.(*AccessToken.AccessToken)
				if !ok {
					return errors.New("Need *AccessToken.AccessToken")
				}
				a.appSecret = id
				return nil
			})
	return nil
}
func (a *Code2Session) GetSession(js_code string) (string, error) {
	type weiChatResult struct {
		openid      string
		session_key string
		unionid     string
		errcode     int
		errmsg      string
	}
	baseURL, err := url.Parse(a.url)
	if err != nil {
		return "", err
	}

	params := url.Values{}
	params.Add("grant_type", a.grant_type)
	params.Add("appid", a.appid)
	params.Add("secret", a.appSecret.GetAppSecret())
	params.Add("js_code", js_code)
	baseURL.RawQuery = params.Encode()
	result, err := http.Get(baseURL.String())
	if err != nil {
		return "", err
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
		return "", errors.New("Get weichat access token error")
	}
	body, err := ioutil.ReadAll(result.Body)
	var r weiChatResult
	err = json.Unmarshal(body, &r)
	if err != nil {
		return "", err
	}
	if r.errcode != 0 {
		log.Error(r.errmsg)
		return "", errors.New(r.errmsg)
	}
	return r.openid, nil
}
func (a *Code2Session) Stop() error {
	return nil
}
func init() {
	WeiChat.GetRegisterWeiChat().RegisterWeiChatFunction(new(Code2Session))
}
