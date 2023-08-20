package middleware

import (
	"bebecare-go-api-1/beans"
	"bebecare-go-api-1/beans/constants"
	"bebecare-go-api-1/beans/global"
	"bebecare-go-api-1/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const TokenType = "Bearer"

func CreateTokenAuthorizer() gin.HandlerFunc {
	return func(c *gin.Context) {
		response := new(beans.BaseResponse)
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Code = constants.ERR_LOGIN_UNAUTHORIZED_TOKEN
			response.Message = constants.GetResponseMsg(constants.ERR_LOGIN_UNAUTHORIZED_TOKEN)
			c.AbortWithStatusJSON(http.StatusOK, response)
			return
		}

		authKeys := strings.Split(authHeader, " ")
		if strings.Compare(authKeys[0], TokenType) != 0 {
			response.Code = constants.ERR_LOGIN_UNAUTHORIZED_TOKEN
			response.Message = constants.GetResponseMsg(constants.ERR_LOGIN_UNAUTHORIZED_TOKEN)
			c.AbortWithStatusJSON(http.StatusOK, response)
			return
		}
		if authKeys[1] == "" {
			response.Code = constants.ERR_LOGIN_UNAUTHORIZED_TOKEN
			response.Message = constants.GetResponseMsg(constants.ERR_LOGIN_UNAUTHORIZED_TOKEN)
			c.AbortWithStatusJSON(http.StatusOK, response)
			return
		}

		authToken := authKeys[1]
		_, err := models.SelectValidToken(authToken)
		if err != nil {
			response.Code = constants.ERR_LOGIN_UNAUTHORIZED_TOKEN
			response.Message = constants.GetResponseMsg(constants.ERR_LOGIN_UNAUTHORIZED_TOKEN)
			c.AbortWithStatusJSON(http.StatusOK, response)
			return
		}

		// 검증이 완료되면 글로벌 변수값 할당
		global.UserToken = authToken

		c.Next()
	}
}
