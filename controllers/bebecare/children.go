package bebecare

import (
	"bebecare-go-api-1/beans/constants"
	"bebecare-go-api-1/beans/db_object"
	"bebecare-go-api-1/beans/global"
	"bebecare-go-api-1/beans/protocols"
	"bebecare-go-api-1/beans/protocols/children"
	"bebecare-go-api-1/db"
	"bebecare-go-api-1/models/children"
	"bebecare-go-api-1/models/user"
	"bebecare-go-api-1/utils/http_util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InsertChildren(c *gin.Context) {
	response := protocols.BaseResponse{Code: constants.SUCCESS, Message: constants.GetResponseMsg(constants.SUCCESS)}

	// request 검증
	request := new(children.InsertChildrenRequest)
	err := c.BindJSON(&request)
	if err != nil || !request.IsValidParameter() {
		response.Code = constants.ERR_MSG_INVALID_PARAMETER
		response.Message = constants.GetResponseMsg(constants.ERR_MSG_INVALID_PARAMETER)
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}

	trx, _ := db.DB.Beginx()
	defer trx.Rollback()

	// 등록하는 부모 정보 가져오기
	parentData, err := userModel.GetUserInfoWithToken(global.UserToken)
	if err != nil {
		response.Code = constants.ERR_DB_NODATA
		response.Message = constants.GetResponseMsg(constants.ERR_DB_NODATA)
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}

	// 신규아기 등록
	insertData := new(db_object.InsertChildren)
	insertData.Trx = trx
	insertData.UserIdx = parentData.Idx
	insertData.Name = request.Name
	insertData.Birthday = request.Birthday
	insertData.Gender = request.Gender
	insertData.Tall = request.Tall
	insertData.Weight = request.Weight
	insertData.HeadSize = request.HeadSize
	insertData.ImageUrl = request.ImageUrl
	childrenIdx, err := childrenModel.InsertChildren(insertData)
	if err != nil {
		trx.Rollback()
		response.Code = constants.ERR_DB_INSERT_DATA
		response.Message = constants.GetResponseMsg(constants.ERR_DB_INSERT_DATA)
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}
	// 아기와 부모 연결
	err = childrenModel.InsertRelParentChildren(parentData.Idx, childrenIdx)
	if err != nil {
		trx.Rollback()
		response.Code = constants.ERR_DB_INSERT_DATA
		response.Message = constants.GetResponseMsg(constants.ERR_DB_INSERT_DATA)
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}

	// 아이 등록한 부모가 마스터가 아니면 마스터로 지정
	if parentData.UserType != constants.USER_TYPE_MASTER {
		err = userModel.SetMaster(parentData.Idx)
		if err != nil {
			trx.Rollback()
			response.Code = constants.ERR_DB_UPDATE_DATA
			response.Message = constants.GetResponseMsg(constants.ERR_DB_UPDATE_DATA)
			http_util.JsonResponse(c, http.StatusOK, response)
			return
		}
	}

	trx.Commit()
	http_util.JsonResponse(c, http.StatusOK, response)
}

func GetChildrenInfo(c *gin.Context) {
	response := protocols.BaseResponse{Code: constants.SUCCESS, Message: constants.GetResponseMsg(constants.SUCCESS)}

	// request 검증
	request := new(children.GetChildrenInfoRequest)
	err := c.BindJSON(&request)
	if err != nil || !request.IsValidParameter() {
		response.Code = constants.ERR_MSG_INVALID_PARAMETER
		response.Message = constants.GetResponseMsg(constants.ERR_MSG_INVALID_PARAMETER)
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}

	childrenInfo, err := childrenModel.GetChildrenInfo(request.Idx)
	if err != nil {
		response.Code = constants.ERR_DB_NODATA
		response.Message = constants.GetResponseMsg(constants.ERR_DB_NODATA)
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}

	response.Data = childrenInfo
	http_util.JsonResponse(c, http.StatusOK, response)
}

func ModifyChildren(c *gin.Context) {
	response := protocols.BaseResponse{Code: constants.SUCCESS, Message: constants.GetResponseMsg(constants.SUCCESS)}

	// request 검증
	request := new(children.ModifyChildrenRequest)
	err := c.BindJSON(&request)
	if err != nil || !request.IsValidParameter() {
		response.Code = constants.ERR_MSG_INVALID_PARAMETER
		response.Message = constants.GetResponseMsg(constants.ERR_MSG_INVALID_PARAMETER)
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}

	updateData := new(db_object.ModifyChildren)
	updateData.Idx = request.Idx
	updateData.Name = request.Name
	updateData.Birthday = request.Birthday
	updateData.Gender = request.Gender
	updateData.Tall = request.Tall
	updateData.Weight = request.Weight
	updateData.HeadSize = request.HeadSize
	updateData.ImageUrl = request.ImageUrl

	err = childrenModel.ModifyChildren(updateData)
	if err != nil {
		response.Code = constants.ERR_DB_UPDATE_DATA
		response.Message = constants.GetResponseMsg(constants.ERR_DB_UPDATE_DATA)
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}

	http_util.JsonResponse(c, http.StatusOK, response)
}

func DeleteChildren(c *gin.Context) {
	response := protocols.BaseResponse{Code: constants.SUCCESS, Message: constants.GetResponseMsg(constants.SUCCESS)}

	// request 검증
	request := new(children.DeleteChildrenRequest)
	err := c.BindJSON(&request)
	if err != nil || !request.IsValidParameter() {
		response.Code = constants.ERR_MSG_INVALID_PARAMETER
		response.Message = constants.GetResponseMsg(constants.ERR_MSG_INVALID_PARAMETER)
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}

	trx, _ := db.DB.Beginx()
	defer trx.Rollback()

	deleteData := new(db_object.DeleteChildren)
	deleteData.Trx = trx
	deleteData.Idx = request.Idx

	// 아기정보 삭제
	err = childrenModel.DeleteChildren(deleteData)
	if err != nil {
		trx.Rollback()
		response.Code = constants.ERR_DB_DELETE_DATA
		response.Message = constants.GetResponseMsg(constants.ERR_DB_DELETE_DATA)
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}

	// 아기와 부모 연결데이터 삭제
	err = childrenModel.DeleteRelParentChildren(deleteData)
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
