package Core

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

func InitFeatures()  error{
	for featureType,feature := range featureMap{
		log.Debug(fmt.Sprint("Init ",featureType))
		err :=feature.Init()
		if err!=nil{
			log.Error(fmt.Sprint("Error when init ",featureType,err.Error()))
			return err
		}
	}
	return nil
}

