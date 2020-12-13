package Route

import "github.com/gin-gonic/gin"

type TestRoute struct {

}

func (re *TestRoute)Init(g *gin.RouterGroup)  {
	g.GET("/test",re.Test)
}

func (re *TestRoute) Test(c *gin.Context){
	 c.JSON(200,nil)
}
