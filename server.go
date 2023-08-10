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

	//mux := http.NewServeMux()
	//mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	w.Header().Set("Access-Control-Allow-Origin", "*")
	//})
	//
	//// cors.Default() setup the middleware with default options being
	//// all origins accepted with simple methods (GET, POST). See
	//// documentation below for more options.
	//handler := cors.Default().Handler(mux)
	//http.ListenAndServe(":8080", handler)

	// ROUTE Setting
	bebecareNoAuthAPI := g.Group("/api/bebecare/noauth/")
	bebecareNoAuthAPI.Use(middleware.AccessPass())
	{
		bebecareNoAuthAPI.POST("/login", bebecare.CheckUser)
	}

	bebecareMainAPI := g.Group("/api/bebecare/service")
	bebecareMainAPI.Use(middleware.AccessPass())
	bebecareMainAPI.Use(middleware.CreateTokenAuthorizer())
	{
		bebecareMainAPI.POST("/get_user", bebecare.GetUser)
		bebecareMainAPI.POST("/children/insert", bebecare.InsertChildren)
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
