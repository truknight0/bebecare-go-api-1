package bebecare

import (
	"bebecare-go-api-1/beans"
	"bebecare-go-api-1/beans/constants"
	"bebecare-go-api-1/beans/db_object"
	"bebecare-go-api-1/beans/global"
	"bebecare-go-api-1/beans/invite"
	"bebecare-go-api-1/db"
	childrenModel "bebecare-go-api-1/models/children"
	"bebecare-go-api-1/models/invite"
	"bebecare-go-api-1/models/user"
	"bebecare-go-api-1/utils"
	"bebecare-go-api-1/utils/http_util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// MakeInviteCode 초대코드 생성
func MakeInviteCode(c *gin.Context) {
	response := beans.BaseResponse{Code: constants.SUCCESS, Message: constants.GetResponseMsg(constants.SUCCESS)}

	request := new(invite.MakeInviteCodeRequest)
	err := c.BindJSON(&request)
	if err != nil || !request.IsValidParameter() {
		response.Code = constants.ERR_MSG_INVALID_PARAMETER
		response.Message = constants.GetResponseMsg(constants.ERR_MSG_INVALID_PARAMETER)
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}

	// 난수 생성
	var inviteCode = utils.MakeInviteCodeUnique()

	if inviteCode == 0 {
		response.Code = constants.ERR_FAIL
		response.Message = constants.GetResponseMsg(constants.ERR_FAIL)
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}

	// 생성된 난수 등록
	var token = global.UserToken
	userInfo, err := userModel.GetUserInfoWithToken(token) // 난수 생성자 정보 가져오기
	if err != nil {
		response.Code = constants.ERR_DB_NODATA
		response.Message = constants.GetResponseMsg(constants.ERR_DB_NODATA)
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}
	err = inviteModel.InsertInviteCode(inviteCode, userInfo.Idx) // 난수 저장
	if err != nil {
		response.Code = constants.ERR_DB_INSERT_DATA
		response.Message = constants.GetResponseMsg(constants.ERR_DB_INSERT_DATA)
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}
	insertData := new(db_object.RelInviteCodeAndUser)
	insertData.InviteCode = inviteCode
	insertData.UserIdx = userInfo.Idx
	insertData.UserName = userInfo.Name
	insertData.UserRole = userInfo.Role
	err = inviteModel.RelInviteCodeAndUser(insertData) // 난수와 유저 연결
	if err != nil {
		response.Code = constants.ERR_DB_INSERT_DATA
		response.Message = constants.GetResponseMsg(constants.ERR_DB_INSERT_DATA)
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}

	jsonInviteCode := invite.MakeInviteCodeResponse{InviteCode: inviteCode}
	response.Data = jsonInviteCode

	http_util.JsonResponse(c, http.StatusOK, response)
}

// JoinInviteCode 초대코드로 등록하기
func JoinInviteCode(c *gin.Context) {
	response := beans.BaseResponse{Code: constants.SUCCESS, Message: constants.GetResponseMsg(constants.SUCCESS)}

	request := new(invite.JoinInviteCodeRequest)
	err := c.BindJSON(&request)
	if err != nil || !request.IsValidParameter() {
		response.Code = constants.ERR_MSG_INVALID_PARAMETER
		response.Message = constants.GetResponseMsg(constants.ERR_MSG_INVALID_PARAMETER)
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}

	// 난수검증
	inviteInfo, err := inviteModel.GetInviteCodeInfo(request.InviteCode)
	if err != nil {
		response.Code = constants.ERR_INVITE_CODE
		response.Message = constants.GetResponseMsg(constants.ERR_INVITE_CODE)
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}

	trx, _ := db.DB.Beginx()
	defer trx.Rollback()

	// 초대데이터 등록
	var token = global.UserToken
	userInfo, err := userModel.GetUserInfoWithToken(token) // 초대받은 유저 정보 가져오기
	insertData := new(db_object.RelInviteCodeAndUser)
	insertData.Trx = trx
	insertData.InviteCode = inviteInfo.InviteCode
	insertData.UserIdx = userInfo.Idx
	insertData.UserName = userInfo.Name
	insertData.UserRole = userInfo.Role
	err = inviteModel.RelInviteCodeAndUser(insertData) // 초대코드와 유저 연결
	if err != nil {
		response.Code = constants.ERR_DB_INSERT_DATA
		response.Message = constants.GetResponseMsg(constants.ERR_DB_INSERT_DATA)
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}

	// 초대코드를 생성한 유저와 연결된 아이 정보 가져오기
	childrenList, _ := childrenModel.GetUserChildrenList(inviteInfo.UserIdx)
	// 아기와 초대받은 유저 연결
	err = inviteModel.RelInviteUserAndChildren(userInfo.Idx, childrenList)
	if err != nil {
		response.Code = constants.ERR_DB_INSERT_DATA
		response.Message = constants.GetResponseMsg(constants.ERR_DB_INSERT_DATA)
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}

	// 초대받은 유저 타입 변경
	err = userModel.SetVisitor(userInfo.Idx)
	if err != nil {
		response.Code = constants.ERR_DB_UPDATE_DATA
		response.Message = constants.GetResponseMsg(constants.ERR_DB_UPDATE_DATA)
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}

	trx.Commit()
	http_util.JsonResponse(c, http.StatusOK, response)
}
