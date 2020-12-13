package Core

import "reflect"

type Feature interface {
	Init() error
	Stop() error
}

var featureMap map[reflect.Type]Feature = map[reflect.Type]Feature{}
