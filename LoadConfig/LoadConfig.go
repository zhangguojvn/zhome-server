package LoadConfig

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"zhome-server/Core"
)

type LoadConfig struct {
	config Config
}

func(l *LoadConfig) loadConfig(path string) (Config, error) {
	fileObj, err := os.Open(path)
	defer func() {
		if fileObj != nil{
			err = fileObj.Close()
			if err != nil {
				log.Error(err.Error())
			}
		}
	}()
	var config Config
	if err != nil {
		return config, err
	}
	byteConfig, err := ioutil.ReadAll(fileObj)
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(byteConfig, &config)
	return config, err
}

func(l *LoadConfig) GetConfig() Config{
	return l.config
}

func(l *LoadConfig) Init() error{
	var err error
	l.config, err = l.loadConfig("config.json")
	if err!=nil{
		log.Error(fmt.Sprint("can't read config"+err.Error()))
	}
	return err
}

func (l *LoadConfig) Stop() error  {
	return nil
}

func init(){
	Core.RegisterFeatures(new(LoadConfig))
}
