package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"zhome-server/Core"
	_ "zhome-server/DataBase"
	_ "zhome-server/LoadConfig"
	_ "zhome-server/Route"
)


func main()  {
	//set logrus
	//log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: false,
	})

	defer func() {
		err := Core.StopFeatures()
		if err != nil{
			log.Error(fmt.Sprint("Error Stop ",err.Error()))
		}
	}()
	log.Info("Init Features")
	err:=Core.InitFeatures()
	if err != nil{
		log.Error(fmt.Sprint(err.Error()))
		os.Exit(1)
	}
	err=Core.InitRequire()
	if err != nil{
		log.Error(fmt.Sprint(err.Error()))
		os.Exit(1)
	}
	osSignals := make(chan os.Signal,1)
	signal.Notify(osSignals,os.Interrupt,syscall.SIGTERM)
	<-osSignals
}

