package middleware

import (
	"bebecare-go-api-1/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IpAuthorized(c *gin.Context) {
	fmt.Println("call by " + c.ClientIP())
	//if c.ClientIP() != "localhost"  && c.ClientIP() != "::1"{
	//	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"return": "error"})
	//	return
	//}
	c.Next()
}
func SessionAuthorized(c *gin.Context) {

	var apiKeys = [...]string{
		// 이브릿지 API KEY HASH
		"BvpYrqEBsmvOaorE8JOvNXAj1z1lOsfgHluBUmPoPVA=",
		// 인스타워시 API KEY HASH
		"OGY0MTU5ZWZmNjNmZTFkOGYxZWYzMGM2MDY1YmRhNWM=",
	}
	fmt.Println("SessionAuthorized")
	authHeader := c.GetHeader("Authorization")

	count := 0
	for i := 0; i < len(apiKeys); i++ {
		key := apiKeys[i]
		if authHeader == key {
			count++
		}
	}

	if count <= 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "message": "Authorization ERROR"})
		fmt.Println("SessionAuthorized FAILED!!!")
		utils.PrintConsoleLog(SessionAuthorized, "END", "Authorization ERROR")
		return
	}
	c.Next()
}
