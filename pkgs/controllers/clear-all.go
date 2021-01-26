package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrzack99s/mrz-captive-portal-operator/pkgs/configs"
	"github.com/mrzack99s/mrz-captive-portal-operator/pkgs/instruction_sets"
	"github.com/mrzack99s/mrz-captive-portal-operator/pkgs/runtime"
	"github.com/mrzack99s/mrz-captive-portal-operator/pkgs/structs"
)

func ClearAll(c *gin.Context) {

	var input structs.ZAuthIPAddressAndShareKey
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.ShareKey == configs.SystemConfig.ZAuth.API.ShareKey {
		command := instruction_sets.GetInitCommand()
		runtime.Run(command)

		c.JSON(http.StatusOK, gin.H{
			"status": true,
		})

	} else {

		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"cause":  "Invalid! share key",
		})

	}
}
