package controllers

import (
	"bebecare-go-api-1/beans"
	"bebecare-go-api-1/beans/constants"
	"bebecare-go-api-1/db"
	"bebecare-go-api-1/model"
	"bebecare-go-api-1/utils"
	"bebecare-go-api-1/utils/http_util"
	"bebecare-go-api-1/utils/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func GetAuthToken(c *gin.Context) {
	//N95RxRTZHWUsaD6HEdz0 : VGhiWGZRNnBZU1EzbjI2N2wxVlE=
	response := beans.GetAuthTokenResponse{
		BaseResponse: beans.BaseResponse{Code: constants.SUCCESS, Message: constants.GetResponseMsg(constants.SUCCESS)},
	}
	request := new(beans.GetAuthTokenRequest)
	err := c.BindJSON(request)
	if err != nil {
		log.ERROR(err.Error())
		response.Code = constants.ERR_MSG_INVALID_PARAMETER
		response.Message = constants.GetResponseMsg(constants.ERR_MSG_INVALID_PARAMETER)
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}
	if !request.IsValidParameter() {
		log.ERROR("Invalid Parameter : " + request.ToString())
		response.Code = constants.ERR_MSG_INVALID_PARAMETER
		response.Message = constants.GetResponseMsg(constants.ERR_MSG_INVALID_PARAMETER)
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}

	// AccessKey, AccessSecret 유효한지 확인
	accessKey, err := model.SelectAccessKey(request.AccessKey, request.AccessSecret)
	if err != nil {
		response.Code = constants.ERR_LOGIN_UNAUTHORIZED_TOKEN
		response.Message = constants.GetResponseMsg(constants.ERR_LOGIN_UNAUTHORIZED_TOKEN)
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}

	// 유효한 authToken 이 있으면 그냥 그 token 을 줌
	authToken, err := model.SelectValidTokenFromAccessKey(accessKey.Idx)
	if authToken != nil {
		response.AuthToken = authToken.Token
		response.ExpiredAt = authToken.ExpiredAt.Unix()
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}

	// authToken 새로 발급
	tx, _ := db.DB.Beginx()
	defer tx.Rollback()

	authToken = new(model.AuthToken)
	authToken.Trx = tx
	authToken.AccessKeyIdx = accessKey.Idx
	authToken.Token = utils.Base64Encoding(utils.GenerateUUID())
	authToken.ExpiredAt = time.Now().AddDate(0, 6, 0)

	// 기존 token invalid
	err = authToken.InvalidToken()
	if err != nil {
		response.Code = constants.ERR_LOGIN_UNAUTHORIZED_TOKEN
		response.Message = constants.GetResponseMsg(constants.ERR_LOGIN_UNAUTHORIZED_TOKEN)
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}
	// 새 token 저장
	err = authToken.InsertAuthToken()
	if err != nil {
		response.Code = constants.ERR_LOGIN_UNAUTHORIZED_TOKEN
		response.Message = constants.GetResponseMsg(constants.ERR_LOGIN_UNAUTHORIZED_TOKEN)
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}

	tx.Commit()

	response.AuthToken = authToken.Token
	response.ExpiredAt = authToken.ExpiredAt.Unix()

	http_util.JsonResponse(c, http.StatusOK, response)
}

func CreateAccessKey(c *gin.Context) {
	accessKey := utils.GenerateRandomString(20)
	accessSecret := utils.GenerateRandomString(20)
	accessSecret = utils.Base64Encoding(accessSecret)
	authToken := utils.Base64Encoding(utils.GenerateUUID())

	c.String(http.StatusOK, accessKey+" : "+accessSecret+" : "+authToken)
}
