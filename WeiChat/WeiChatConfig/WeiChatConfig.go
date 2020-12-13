package WeiChatConfig

import "zhome-server/WeiChat"

type WeiChatConfig struct {
	appid  string
	secret string
}

func (a *WeiChatConfig) Init() error {
	a.appid = WeiChat.GetRegisterWeiChat().GetConfig().AppID
	a.secret = WeiChat.GetRegisterWeiChat().GetConfig().Secret
	return nil
}

func (a *WeiChatConfig) Stop() error {
	return nil
}

func (a *WeiChatConfig) GetAppID() string {
	return a.appid
}
func (a *WeiChatConfig) GetSecret() string {
	return a.secret
}
func init() {
	WeiChat.GetRegisterWeiChat().RegisterWeiChatFunction(new(WeiChatConfig))
}
