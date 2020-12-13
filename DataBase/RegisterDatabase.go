package DataBase

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
	"zhome-server/Core"
	"zhome-server/LoadConfig"

	log "github.com/sirupsen/logrus"
)

var registerdb *RegisterDatabase

type DataBaseObj interface {
	Init(db *sql.DB) error
	Get(interface{})(interface{},error)
	Put(interface{})error
}

type RegisterDatabase struct {
	db *sql.DB
	databseMap map[reflect.Type]DataBaseObj
}
func (r *RegisterDatabase)RegisterDatabase(dbobj DataBaseObj,data interface{}){
	log.Debug(fmt.Sprint("Register Database :",reflect.TypeOf(dbobj)))
	r.databseMap[reflect.TypeOf(data)]=dbobj
	if r.db != nil{
		dbobj.Init(r.db)
	}

}
func(r *RegisterDatabase) initDatabases(db *sql.DB){
	r.db = db
	for _,dbobj := range r.databseMap{
		dbobj.Init(db)
	}
}
func (r *RegisterDatabase)Init() error {
	return nil
}

func (r *RegisterDatabase) Stop() error {
	if r.db != nil{
		err := r.db.Close()
		return err
	}
	return nil
}

func GetRegisterDatabase() *RegisterDatabase {
	if registerdb == nil{
		registerdb = new(RegisterDatabase)
	}
	return registerdb
}

func init(){
	Core.RegisterFeatures(GetRegisterDatabase())
	Core.RequireFeatures(
		reflect.TypeOf(
			new(LoadConfig.LoadConfig),
		),
			func(feature interface{}) error {
				c,ok := feature.(*LoadConfig.LoadConfig)
				if !ok{
					return errors.New("Register Databse need LoadConfig.")
				}
				config:=c.GetConfig()
				rd := GetRegisterDatabase()
				db, err := sql.Open("mysql", config.DataBase)
				if err != nil{
					return err
				}
				var version string
				err=db.QueryRow("SELECT VERSION()").Scan(&version)
				if err != nil{
					return err
				}
				log.Info(fmt.Sprint("Connected to:", version))
				rd.initDatabases(db)
				return nil
	})
}