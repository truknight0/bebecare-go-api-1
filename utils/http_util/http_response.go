package http_util

import (
	"bebecare-go-api-1/utils/log"
	"fmt"
	"github.com/gin-gonic/gin"
)

func JsonResponse(c *gin.Context, httpStatus int, response interface{}) {
	log.DEBUG(fmt.Sprintf("%#v", response))
	c.JSON(httpStatus, response)
}
