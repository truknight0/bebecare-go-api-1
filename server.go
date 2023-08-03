package main

import (
	"bebecare-go-api-1/controllers/bebecare"
	"bebecare-go-api-1/service/middleware"
	"bebecare-go-api-1/utils/config"
	"bebecare-go-api-1/utils/log"
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.New()
	g.Use(gin.Logger())

	// ROUTE Setting
	bebecareNoAuthAPI := g.Group("/api/bebecare/noauth/")
	{
		bebecareNoAuthAPI.POST("/login", bebecare.CheckUser)
	}

	bebecareMainAPI := g.Group("/api/bebecare/service")
	bebecareMainAPI.Use(middleware.CreateTokenAuthorizer())
	{
		bebecareMainAPI.POST("/get_user", bebecare.GetUser)
		bebecareMainAPI.POST("/children/insert", bebecare.InsertChildren)
	}

	port := config.GetStringDefault("server.port", "8080")
	sslEnable := config.GetBoolDefault("ssl.enable", false)

	var err error
	if sslEnable {
		sslCert := config.GetString("ssl.cert")
		sslKey := config.GetString("ssl.key")
		err = g.RunTLS(":"+port, sslCert, sslKey)
	} else {
		err = g.Run(":" + port)
	}
	if err != nil {
		log.ERROR(err.Error())
	}
}
