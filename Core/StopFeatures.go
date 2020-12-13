package Core
import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

func StopFeatures()  error{
	for featureType,feature := range featureMap{
		log.Debug(fmt.Sprint("Stop ",featureType))
		err :=feature.Stop()
		if err!=nil{
			log.Error(fmt.Sprint("Error when init ",featureType,err.Error()))
			return err
		}
	}
	return nil
}
