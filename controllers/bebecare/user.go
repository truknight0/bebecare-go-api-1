package bebecare

import (
	"bebecare-go-api-1/beans"
	"bebecare-go-api-1/beans/constants"
	"bebecare-go-api-1/beans/db_object"
	"bebecare-go-api-1/beans/user"
	"bebecare-go-api-1/db"
	"bebecare-go-api-1/model/user"
	"bebecare-go-api-1/utils"
	"bebecare-go-api-1/utils/http_util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type Token struct {
	Token string `json:"token"`
}

func GetUser(c *gin.Context) {
	response := user.GetUserResponse{
		BaseResponse: beans.BaseResponse{Code: constants.SUCCESS, Message: constants.GetResponseMsg(constants.SUCCESS)},
	}

	request := new(user.GetUserRequest)
	err := c.BindJSON(&request)
	if err != nil || !request.IsValidParameter() {
		response.Code = constants.ERR_MSG_INVALID_PARAMETER
		response.Message = constants.GetResponseMsg(constants.ERR_MSG_INVALID_PARAMETER)
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}

	// 유저정보 가져오기
	userInfo, err := userModel.GetUserInfo(request.Name, request.Phone, request.Role)
	if err != nil {
		response.Code = constants.ERR_DB_NODATA
		response.Message = constants.GetResponseMsg(constants.ERR_DB_NODATA)
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}
	response.Data = userInfo

	http_util.JsonResponse(c, http.StatusOK, response)
}

func CheckUser(c *gin.Context) {
	response := user.CheckUserResponse{
		BaseResponse: beans.BaseResponse{Code: constants.SUCCESS, Message: constants.GetResponseMsg(constants.SUCCESS)},
	}

	trx, _ := db.DB.Beginx()
	defer trx.Rollback()

	// 해더에 토큰이 있는지 확인
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		request := new(user.CheckUserRequest)
		err := c.BindJSON(&request)
		if err != nil || !request.IsValidParameter() {
			response.Code = constants.ERR_MSG_INVALID_PARAMETER
			response.Message = constants.GetResponseMsg(constants.ERR_MSG_INVALID_PARAMETER)
			http_util.JsonResponse(c, http.StatusOK, response)
			return
		}

		// 기존 등록된 회원인지 확인
		userInfo, _ := userModel.GetUserInfo(request.Name, request.Phone, request.Role)
		if userInfo != nil {
			var userToken string
			userToken, err = userModel.GetUserToken(request.Name, request.Phone, request.Role)
			if err != nil {
				response.Code = constants.ERR_LOGIN_UNAUTHORIZED_TOKEN
				response.Message = constants.GetResponseMsg(constants.ERR_LOGIN_UNAUTHORIZED_TOKEN)
				http_util.JsonResponse(c, http.StatusOK, response)
				return
			}
			jsonToken := Token{userToken}
			response.Data = jsonToken

			http_util.JsonResponse(c, http.StatusOK, response)
			return
		}

		// 토큰이 없으면 신규 회원가입
		insertData := new(db_object.InsertUser)
		insertData.Name = request.Name
		insertData.Phone = request.Phone
		insertData.Role = request.Role
		// 회원정보 저장
		var userIdx int // 마지막에 입력된 회원 일련번호
		userIdx, err = userModel.InsertUser(insertData)
		if err != nil {
			trx.Rollback()
			response.Code = constants.ERR_DB_INSERT_DATA
			response.Message = constants.GetResponseMsg(constants.ERR_DB_INSERT_DATA)
			http_util.JsonResponse(c, http.StatusOK, response)
			return
		}
		// 토큰 신규등록
		token := utils.Base64Encoding(utils.GenerateUUID())
		err = userModel.InsertToken(userIdx, token)
		if err != nil {
			trx.Rollback()
			response.Code = constants.ERR_DB_INSERT_DATA
			response.Message = constants.GetResponseMsg(constants.ERR_DB_INSERT_DATA)
			http_util.JsonResponse(c, http.StatusOK, response)
			return
		}
		trx.Commit()

		// 토큰 리턴
		jsonToken := Token{token}
		response.Data = jsonToken

	} else {
		// 토큰이 있으면 토큰 리턴하고 종료
		tokenArr := strings.Split(authHeader, " ")
		checkToken := tokenArr[1]
		var userToken string
		userToken, err := userModel.GetUserInfoWithToken(checkToken)
		if err != nil {
			response.Code = constants.ERR_DB_NODATA
			response.Message = constants.GetResponseMsg(constants.ERR_DB_NODATA)
			http_util.JsonResponse(c, http.StatusOK, response)
			return
		}

		// 토큰 리턴
		jsonToken := Token{userToken}
		response.Data = jsonToken
	}

	http_util.JsonResponse(c, http.StatusOK, response)
}
