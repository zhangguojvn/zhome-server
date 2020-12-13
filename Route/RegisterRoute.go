package Route

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"reflect"
	"zhome-server/Core"
	"zhome-server/LoadConfig"
)

var registerRoute *RegisterRoute
type RouteObj interface {
	Init(g *gin.RouterGroup)
}
type RegisterRoute struct {
	routeArray []RouteObj
	route *gin.RouterGroup
}
func (re *RegisterRoute) RegisterRoutes(r RouteObj){
	log.Debug(fmt.Sprint("Register Database :",reflect.TypeOf(r)))
	re.routeArray = append(re.routeArray,r)
	if re.route != nil{
		r.Init(re.route)
	}
}

func(re *RegisterRoute) initRoutes(g *gin.RouterGroup)  {
	re.route=g
	for _,routeobj := range re.routeArray{
		routeobj.Init(g)
	}
}
func (re *RegisterRoute)Init()error{
	return nil
}
func (re *RegisterRoute)Stop() error{
	return nil
}

func GetRegisterRoute() *RegisterRoute {
	if registerRoute == nil{
		registerRoute = new(RegisterRoute)
	}
	return registerRoute
}

func init(){
	Core.RegisterFeatures(GetRegisterRoute())
	Core.RequireFeatures(
		reflect.TypeOf(
			new(LoadConfig.LoadConfig),
			),
			func(feature interface{}) error {
				c,ok := feature.(*LoadConfig.LoadConfig)
				if !ok{
					return errors.New("Register Route need LoadConfig.")
				}
				config:=c.GetConfig()
				rr := GetRegisterRoute()
				router := gin.New()
				router.Use(gin.Logger())
				router.Use(gin.Recovery())
				//set base path
				r := router.Group(config.RootPath)
				rr.initRoutes(r)
				go router.Run(fmt.Sprint(config.Listen,":",config.Port))
				return nil
	})
}