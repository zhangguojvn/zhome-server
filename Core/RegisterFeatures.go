package Core

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"reflect"
)

func RegisterFeatures(f Feature)  {
	log.Debug(fmt.Sprint("Add ",reflect.TypeOf(f), " to Feature"))
	featureMap[reflect.TypeOf(f)]=f
}
