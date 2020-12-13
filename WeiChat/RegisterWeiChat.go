package WeiChat

import (
	"reflect"
	"zhome-server/Core"
	"zhome-server/LoadConfig"
)

var rwc *RegisterWeiChat
type RegisterWeiChat struct {

}

func (receiver *RegisterWeiChat) Init() error  {
	return nil
}
func (receiver *RegisterWeiChat) Stop() error  {
	return nil
}

func GetRegisterWeiChat() *RegisterWeiChat{
	if rwc ==nil{
		rwc = new(RegisterWeiChat)
	}
	return rwc
}
func init(){
	Core.RegisterFeatures(GetRegisterWeiChat())
	Core.RequireFeatures(
		reflect.TypeOf(
			new(
				LoadConfig.LoadConfig,
				),
			),
			func(feature interface{}) error {
				return nil
			},
			)
}
