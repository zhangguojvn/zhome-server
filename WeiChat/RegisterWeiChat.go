package WeiChat

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"reflect"
	"zhome-server/Core"
	"zhome-server/LoadConfig"
)

var rwc *RegisterWeiChat

type RegisterWeiChat struct {
	config    LoadConfig.Config
	functions map[reflect.Type]WeiChatFunction
	requires  []struct {
		callback func(interface{}) error
		needType reflect.Type
	}
}

func (receiver *RegisterWeiChat) Init() error {
	for functionType, function := range receiver.functions {
		log.Debug(fmt.Sprint("Init WeiChat Function ", functionType))
		err := function.Init()
		if err != nil {
			return err
		}
	}
	for _, require := range receiver.requires {
		err := require.callback(receiver.functions[require.needType])
		return err
	}
	return nil
}
func (receiver *RegisterWeiChat) Stop() error {
	for functionType, function := range receiver.functions {
		log.Debug(fmt.Sprint("Init WeiChat Function ", functionType))
		err := function.Stop()
		if err != nil {
			return err
		}
	}
	return nil
}
func (receiver *RegisterWeiChat) RegisterWeiChatFunction(function WeiChatFunction) {
	log.Debug(fmt.Sprint("RegisterWeiChatFunction :", reflect.TypeOf(function)))
	if receiver.functions == nil {
		receiver.functions = map[reflect.Type]WeiChatFunction{}
	}
	receiver.functions[reflect.TypeOf(function)] = function
}
func (receiver *RegisterWeiChat) RequireWeiChatFunction(p reflect.Type, f func(interface{}) error) {
	receiver.requires = append(receiver.requires, struct {
		callback func(interface{}) error
		needType reflect.Type
	}{callback: f, needType: p})
}
func (receiver *RegisterWeiChat) GetConfig() LoadConfig.Config {
	return receiver.config
}

func GetRegisterWeiChat() *RegisterWeiChat {
	if rwc == nil {
		rwc = new(RegisterWeiChat)
	}
	return rwc
}
func init() {
	Core.RegisterFeatures(GetRegisterWeiChat())
	Core.RequireFeatures(
		reflect.TypeOf(
			new(
				LoadConfig.LoadConfig,
			),
		),
		func(feature interface{}) error {
			c, ok := feature.(*LoadConfig.LoadConfig)
			if !ok {
				return errors.New("Register WeiChat need *LoadConfig.")
			}
			config := c.GetConfig()
			GetRegisterWeiChat().config = config
			return nil
		},
	)
}
