package TimeSeriesDatabase

import (
	"reflect"
	"zhome-server/Core"
	"zhome-server/LoadConfig"
)

var rtsdb *RegisterTSDB
type RegisterTSDB struct {

}

func (receiver *RegisterTSDB) Init() error  {
	return nil
}
func (receiver *RegisterTSDB) Stop() error  {
	return nil
}

func GetRegisterWeiChat() *RegisterTSDB{
	if rtsdb ==nil{
		rtsdb = new(RegisterTSDB)
	}
	return rtsdb
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
