package bebecare

import (
	"bebecare-go-api-1/beans/constants"
	"bebecare-go-api-1/beans/db_object"
	"bebecare-go-api-1/beans/global"
	"bebecare-go-api-1/beans/protocols"
	"bebecare-go-api-1/beans/protocols/user"
	"bebecare-go-api-1/db"
	"bebecare-go-api-1/models/children"
	"bebecare-go-api-1/models/invite"
	"bebecare-go-api-1/models/user"
	"bebecare-go-api-1/utils"
	"bebecare-go-api-1/utils/http_util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var isFirstUser bool

func GetUserInfo(c *gin.Context) {
	response := user.GetUserResponse{
		BaseResponse: protocols.BaseResponse{Code: constants.SUCCESS, Message: constants.GetResponseMsg(constants.SUCCESS)},
	}

	// 유저정보 가져오기
	userInfo, err := userModel.GetUserInfoWithToken(global.UserToken)
	if err != nil {
		response.Code = constants.ERR_DB_NODATA
		response.Message = constants.GetResponseMsg(constants.ERR_DB_NODATA)
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}
	// 아이정보 가져오기
	childrenInfo, _ := childrenModel.GetUserChildrenList(userInfo.Idx)
	// 초대코드 가져오기
	inviteCodeInfo, _ := inviteModel.GetInviteCodeInfoWithUserIdx(userInfo.Idx)
	// 엮인 가족 정보 가져오기
	parents, _ := inviteModel.GetUserListWithInviteCode(global.UserToken, inviteCodeInfo.InviteCode)

	isFirstUser = false
	userData := user.GetUserInfoData{
		IsFirstUser: isFirstUser,
		Token:       userInfo.Token,
		Idx:         userInfo.Idx,
		Name:        userInfo.Name,
		Phone:       userInfo.Phone,
		Role:        userInfo.Role,
		IsPushAgree: userInfo.IsPushAgree,
		UserType:    userInfo.UserType,
		CreatedAt:   userInfo.CreatedAt,
		InviteCode:  inviteCodeInfo.InviteCode,
		Children:    childrenInfo,
		Parents:     parents}

	response.Data = userData

	http_util.JsonResponse(c, http.StatusOK, response)
}

func CheckUser(c *gin.Context) {
	response := protocols.BaseResponse{Code: constants.SUCCESS, Message: constants.GetResponseMsg(constants.SUCCESS)}

	// 유저정보 가져오기
	userInfo, err := userModel.GetUserInfoWithToken(global.UserToken)
	if err != nil {
		response.Code = constants.ERR_DB_NODATA
		response.Message = constants.GetResponseMsg(constants.ERR_DB_NODATA)
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}
	// 아이정보 가져오기
	childrenInfo, _ := childrenModel.GetUserChildrenList(userInfo.Idx)

	isFirstUser = false
	response.Data = childrenInfo

	http_util.JsonResponse(c, http.StatusOK, response)
}

func LoginUser(c *gin.Context) {
	response := user.CheckUserResponse{
		BaseResponse: protocols.BaseResponse{Code: constants.SUCCESS, Message: constants.GetResponseMsg(constants.SUCCESS)},
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
			// 아이정보 가져오기
			childrenInfo, _ := childrenModel.GetUserChildrenList(userInfo.Idx)

			isFirstUser = false
			userData := user.SetUserData{
				IsFirstUser: isFirstUser,
				Token:       userToken,
				Children:    childrenInfo}

			response.Data = userData

			http_util.JsonResponse(c, http.StatusOK, response)
			return
		}

		// 토큰이 없으면 신규 회원가입
		insertData := new(db_object.InsertUser)
		insertData.Trx = trx
		insertData.Name = request.Name
		insertData.Phone = request.Phone
		insertData.Role = request.Role
		insertData.IsPushAgree = request.IsPushAgree
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
		isFirstUser = true
		userData := user.SetUserData{
			IsFirstUser: isFirstUser,
			Token:       token,
			Children:    nil}

		response.Data = userData

	} else {
		// 토큰이 있으면 토큰 리턴하고 종료
		tokenArr := strings.Split(authHeader, " ")
		checkToken := tokenArr[1]
		userInfo, err := userModel.GetUserInfoWithToken(checkToken)
		if err != nil {
			response.Code = constants.ERR_DB_NODATA
			response.Message = constants.GetResponseMsg(constants.ERR_DB_NODATA)
			http_util.JsonResponse(c, http.StatusOK, response)
			return
		}
		// 아이정보 가져오기
		childrenInfo, _ := childrenModel.GetUserChildrenList(userInfo.Idx)

		isFirstUser = false
		userData := user.SetUserData{
			IsFirstUser: isFirstUser,
			Token:       checkToken,
			Children:    childrenInfo}

		response.Data = userData
	}

	http_util.JsonResponse(c, http.StatusOK, response)
}

func ModifyUser(c *gin.Context) {
	response := protocols.BaseResponse{Code: constants.SUCCESS, Message: constants.GetResponseMsg(constants.SUCCESS)}

	request := new(user.ModifyUserRequest)
	err := c.BindJSON(&request)
	if err != nil || !request.IsValidParameter() {
		response.Code = constants.ERR_MSG_INVALID_PARAMETER
		response.Message = constants.GetResponseMsg(constants.ERR_MSG_INVALID_PARAMETER)
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}

	modifyData := new(db_object.ModifyUser)
	modifyData.Name = request.Name
	modifyData.Phone = request.Phone
	modifyData.Role = request.Role
	modifyData.IsPushAgree = request.IsPushAgree
	// 회원정보 수정
	err = userModel.ModifyUser(modifyData)
	if err != nil {
		response.Code = constants.ERR_DB_UPDATE_DATA
		response.Message = constants.GetResponseMsg(constants.ERR_DB_UPDATE_DATA)
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}

	http_util.JsonResponse(c, http.StatusOK, response)
}

func DisconnectUser(c *gin.Context) {
	response := protocols.BaseResponse{Code: constants.SUCCESS, Message: constants.GetResponseMsg(constants.SUCCESS)}

	request := new(user.DisconnectUserRequest)
	err := c.BindJSON(&request)
	if err != nil || !request.IsValidParameter() {
		response.Code = constants.ERR_MSG_INVALID_PARAMETER
		response.Message = constants.GetResponseMsg(constants.ERR_MSG_INVALID_PARAMETER)
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}

	trx, _ := db.DB.Beginx()
	defer trx.Rollback()

	// 연결끊기 권한이 있는 유저인지 확인
	userInfo, err := userModel.GetUserInfoWithToken(global.UserToken)
	if err != nil {
		response.Code = constants.ERR_DB_NODATA
		response.Message = constants.GetResponseMsg(constants.ERR_DB_NODATA)
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}
	if userInfo.UserType != constants.USER_TYPE_MASTER {
		response.Code = constants.ERR_PROCESS_USER_LEVEL
		response.Message = constants.GetResponseMsg(constants.ERR_PROCESS_USER_LEVEL)
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}

	queryData := new(db_object.DisconnectUser)
	queryData.Trx = trx
	queryData.Idx = request.Idx
	err = userModel.DisconnectChildren(queryData)
	if err != nil {
		trx.Rollback()
		response.Code = constants.ERR_DB_DELETE_DATA
		response.Message = constants.GetResponseMsg(constants.ERR_DB_DELETE_DATA)
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}
	err = userModel.DisconnectInviteCode(queryData)
	if err != nil {
		trx.Rollback()
		response.Code = constants.ERR_DB_DELETE_DATA
		response.Message = constants.GetResponseMsg(constants.ERR_DB_DELETE_DATA)
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}

	trx.Commit()
	http_util.JsonResponse(c, http.StatusOK, response)
}

func DeleteUser(c *gin.Context) {
	response := protocols.BaseResponse{Code: constants.SUCCESS, Message: constants.GetResponseMsg(constants.SUCCESS)}

	trx, _ := db.DB.Beginx()
	defer trx.Rollback()

	// 연결끊기 권한이 있는 유저인지 확인
	userInfo, err := userModel.GetUserInfoWithToken(global.UserToken)
	if err != nil {
		response.Code = constants.ERR_DB_NODATA
		response.Message = constants.GetResponseMsg(constants.ERR_DB_NODATA)
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}

	queryData := new(db_object.DisconnectUser)
	queryData.Trx = trx
	queryData.Idx = userInfo.Idx
	err = userModel.DisconnectChildren(queryData)
	if err != nil {
		trx.Rollback()
		response.Code = constants.ERR_DB_DELETE_DATA
		response.Message = constants.GetResponseMsg(constants.ERR_DB_DELETE_DATA)
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}
	err = userModel.DisconnectInviteCode(queryData)
	if err != nil {
		trx.Rollback()
		response.Code = constants.ERR_DB_DELETE_DATA
		response.Message = constants.GetResponseMsg(constants.ERR_DB_DELETE_DATA)
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}
	err = userModel.DeleteAuthToken(queryData)
	if err != nil {
		trx.Rollback()
		response.Code = constants.ERR_DB_DELETE_DATA
		response.Message = constants.GetResponseMsg(constants.ERR_DB_DELETE_DATA)
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}
	err = userModel.DeleteUser(queryData)
	if err != nil {
		trx.Rollback()
		response.Code = constants.ERR_DB_DELETE_DATA
		response.Message = constants.GetResponseMsg(constants.ERR_DB_DELETE_DATA)
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}

	trx.Commit()
	http_util.JsonResponse(c, http.StatusOK, response)
}
