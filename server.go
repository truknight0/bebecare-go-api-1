package main

import (
	"bebecare-go-api-1/controllers/bebecare"
	"bebecare-go-api-1/service/middleware"
	"bebecare-go-api-1/utils/config"
	"bebecare-go-api-1/utils/log"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.New()
	g.Use(gin.Logger())
	g.Use(cors.New(
		cors.Config{
			AllowOrigins: []string{"http://localhost:8000", "http://3.39.235.12:8000"},
			AllowHeaders: []string{"Content-Type", "Authorization"},
			AllowMethods: []string{"POST", "GET", "PUT", "PATCH", "DELETE", "OPTION"},
		}))

	// ROUTE Setting
	bebecareNoAuthAPI := g.Group("/api/bebecare/noauth/")
	bebecareNoAuthAPI.Use(middleware.AccessPass())
	{
		bebecareNoAuthAPI.POST("/login", bebecare.LoginUser)
	}

	bebecareMainAPI := g.Group("/api/bebecare/service")
	bebecareMainAPI.Use(middleware.AccessPass())
	bebecareMainAPI.Use(middleware.CreateTokenAuthorizer())
	{
		bebecareMainAPI.POST("/user/check", bebecare.CheckUser)
		bebecareMainAPI.POST("/user/info", bebecare.GetUserInfo)
		bebecareMainAPI.POST("/user/modify", bebecare.ModifyUser)
		bebecareMainAPI.POST("/user/delete", bebecare.DeleteUser)

		bebecareMainAPI.POST("/parents/disconnect", bebecare.DisconnectUser)

		bebecareMainAPI.POST("/children/info", bebecare.GetChildrenInfo)
		bebecareMainAPI.POST("/children/insert", bebecare.InsertChildren)
		bebecareMainAPI.POST("/children/modify", bebecare.ModifyChildren)
		bebecareMainAPI.POST("/children/delete", bebecare.DeleteChildren)

		bebecareMainAPI.POST("/items/list", bebecare.GetItemList)
		bebecareMainAPI.POST("/items/insert", bebecare.InsertItem)
		bebecareMainAPI.POST("/items/complete", bebecare.CompleteItem)
		bebecareMainAPI.POST("/items/modify", bebecare.ModifyItem)
		bebecareMainAPI.POST("/items/delete", bebecare.DeleteItem)

		bebecareMainAPI.POST("/invite/make", bebecare.MakeInviteCode)
		bebecareMainAPI.POST("/invite/join", bebecare.JoinInviteCode)
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
