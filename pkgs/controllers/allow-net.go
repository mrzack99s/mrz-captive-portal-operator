package controllers

import (
	"crypto/sha256"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrzack99s/mrz-captive-portal-operator/pkgs/configs"
	"github.com/mrzack99s/mrz-captive-portal-operator/pkgs/instruction_sets"
	"github.com/mrzack99s/mrz-captive-portal-operator/pkgs/runtime"
	"github.com/mrzack99s/mrz-captive-portal-operator/pkgs/structs"
)

func AllowNet(c *gin.Context) {

	var input structs.ZAuthAllowNet
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashInputShareKey := sha256.Sum256([]byte(input.ShareKey))
	hashServShareKey := sha256.Sum256([]byte(configs.SystemConfig.ZAuth.API.ShareKey))

	if hashInputShareKey == hashServShareKey {

		command := instruction_sets.GetAllowNet(input.IPAddress)
		runtime.Run(command)

		command = instruction_sets.GetAppendBandwidthControl(input.DlSpeed, input.UpSpeed, input.IPAddress)
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
