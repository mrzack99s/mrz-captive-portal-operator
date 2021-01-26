package main

import (
	"github.com/gin-gonic/gin"
	conf "github.com/mrzack99s/mrz-captive-portal-operator/pkgs/configs"
	"github.com/mrzack99s/mrz-captive-portal-operator/pkgs/instruction_sets"
	"github.com/mrzack99s/mrz-captive-portal-operator/pkgs/routes"
	"github.com/mrzack99s/mrz-captive-portal-operator/pkgs/runtime"
)

// var err error

func main() {

	conf.ParseSystemConfig()

	mode := gin.DebugMode
	if conf.SystemConfig.ZAuth.API.Production {
		mode = gin.ReleaseMode
	}

	gin.SetMode(mode)
	r := routes.SetupRouter()

	runtime.Run(instruction_sets.GetInitCommand())
	//running
	r.Run(conf.SystemConfig.ZAuth.API.Port)
}
