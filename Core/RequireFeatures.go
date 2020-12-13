package Core

import "reflect"

type Require struct {
	callback func(feature interface{})error
	needType reflect.Type
}

var requireMap []Require

func RequireFeatures(p reflect.Type,f func(feature interface{})error)  {
	requireMap=append(requireMap,Require{
		callback: f,
		needType: p,
	})
}
