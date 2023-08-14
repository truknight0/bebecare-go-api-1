package bebecare

import (
	"bebecare-go-api-1/beans"
	"bebecare-go-api-1/beans/children"
	"bebecare-go-api-1/beans/constants"
	"bebecare-go-api-1/beans/db_object"
	"bebecare-go-api-1/beans/global"
	"bebecare-go-api-1/db"
	"bebecare-go-api-1/models/children"
	"bebecare-go-api-1/models/user"
	"bebecare-go-api-1/utils/http_util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InsertChildren(c *gin.Context) {
	response := beans.BaseResponse{Code: constants.SUCCESS, Message: constants.GetResponseMsg(constants.SUCCESS)}

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
		response.Code = constants.ERR_DB_INSERT_DATA
		response.Message = constants.GetResponseMsg(constants.ERR_DB_INSERT_DATA)
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}
	// 아기와 부모 연결
	err = childrenModel.InsertRelParentChildren(parentData.Idx, childrenIdx)
	if err != nil {
		response.Code = constants.ERR_DB_INSERT_DATA
		response.Message = constants.GetResponseMsg(constants.ERR_DB_INSERT_DATA)
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}

	// 아이 등록한 부모가 마스터가 아니면 마스터로 지정
	if parentData.UserType != constants.USER_TYPE_MASTER {
		err = userModel.SetMaster(parentData.Idx)
		if err != nil {
			response.Code = constants.ERR_DB_UPDATE_DATA
			response.Message = constants.GetResponseMsg(constants.ERR_DB_UPDATE_DATA)
			http_util.JsonResponse(c, http.StatusOK, response)
			return
		}
	}

	trx.Commit()
	http_util.JsonResponse(c, http.StatusOK, response)
}

func UpdateChildren(c *gin.Context) {
	response := beans.BaseResponse{Code: constants.SUCCESS, Message: constants.GetResponseMsg(constants.SUCCESS)}

	http_util.JsonResponse(c, http.StatusOK, response)
}
