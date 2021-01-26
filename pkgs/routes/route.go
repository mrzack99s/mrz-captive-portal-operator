package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mrzack99s/mrz-captive-portal-operator/pkgs/controllers"
)

var rootPath = "/zauth/v2beta"

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()

	group := r.Group(rootPath)
	{
		group.POST("logout", controllers.Logout)
		group.POST("allowNet", controllers.AllowNet)
		group.POST("clearAll", controllers.ClearAll)
	}

	return r
}

func preflight(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
}
