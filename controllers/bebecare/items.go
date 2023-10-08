package bebecare

import (
	"bebecare-go-api-1/beans/constants"
	"bebecare-go-api-1/beans/db_object"
	"bebecare-go-api-1/beans/protocols"
	"bebecare-go-api-1/beans/protocols/items"
	"bebecare-go-api-1/models/items"
	"bebecare-go-api-1/utils/http_util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InsertItem(c *gin.Context) {
	response := protocols.BaseResponse{Code: constants.SUCCESS, Message: constants.GetResponseMsg(constants.SUCCESS)}

	request := new(items.InsertItemRequest)
	err := c.BindJSON(&request)
	if err != nil || !request.IsValidParameter() {
		response.Code = constants.ERR_MSG_INVALID_PARAMETER
		response.Message = constants.GetResponseMsg(constants.ERR_MSG_INVALID_PARAMETER)
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}

	insertData := new(db_object.InsertItem)
	insertData.UserIdx = request.UserIdx
	insertData.ChildrenIdx = request.ChildrenIdx
	insertData.Type = request.Type
	insertData.Etc1 = request.Etc1
	insertData.Etc2 = request.Etc2
	insertData.Etc3 = request.Etc3
	insertData.Etc4 = request.Etc4
	insertData.Etc5 = request.Etc5
	insertData.Etc6 = request.Etc6
	insertData.Etc7 = request.Etc7
	insertData.StartTime = request.StartTime
	if request.EndTime == "" {
		insertData.EndTime = nil
	} else {
		insertData.EndTime = request.EndTime
	}

	err = itemsModel.InsertItem(insertData)
	if err != nil {
		response.Code = constants.ERR_DB_INSERT_DATA
		response.Message = constants.GetResponseMsg(constants.ERR_DB_INSERT_DATA)
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}

	http_util.JsonResponse(c, http.StatusOK, response)
}

func CompleteItem(c *gin.Context) {
	response := protocols.BaseResponse{Code: constants.SUCCESS, Message: constants.GetResponseMsg(constants.SUCCESS)}

	request := new(items.CompleteItemRequest)
	err := c.BindJSON(&request)
	if err != nil || !request.IsValidParameter() {
		response.Code = constants.ERR_MSG_INVALID_PARAMETER
		response.Message = constants.GetResponseMsg(constants.ERR_MSG_INVALID_PARAMETER)
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}

	err = itemsModel.CompleteItem(request.Idx)
	if err != nil {
		response.Code = constants.ERR_DB_UPDATE_DATA
		response.Message = constants.GetResponseMsg(constants.ERR_DB_UPDATE_DATA)
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}

	http_util.JsonResponse(c, http.StatusOK, response)
}

func ModifyItem(c *gin.Context) {
	response := protocols.BaseResponse{Code: constants.SUCCESS, Message: constants.GetResponseMsg(constants.SUCCESS)}

	request := new(items.ModifyItemRequest)
	err := c.BindJSON(&request)
	if err != nil || !request.IsValidParameter() {
		response.Code = constants.ERR_MSG_INVALID_PARAMETER
		response.Message = constants.GetResponseMsg(constants.ERR_MSG_INVALID_PARAMETER)
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}

	modifyData := new(db_object.ModifyItem)
	modifyData.Idx = request.Idx
	modifyData.Etc1 = request.Etc1
	modifyData.Etc2 = request.Etc2
	modifyData.Etc3 = request.Etc3
	modifyData.Etc4 = request.Etc4
	modifyData.Etc5 = request.Etc5
	modifyData.Etc6 = request.Etc6
	modifyData.Etc7 = request.Etc7
	// request end_time 값이 없으면 기존 데이터의 end_time 세팅
	itemInfo, err := itemsModel.GetItemInfo(request.Idx)
	if err != nil {
		response.Code = constants.ERR_DB_NODATA
		response.Message = constants.GetResponseMsg(constants.ERR_DB_NODATA)
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}
	if request.StartTime == nil || request.StartTime == "" {
		modifyData.StartTime = itemInfo.StartTime
	} else {
		modifyData.StartTime = request.StartTime
	}
	if request.EndTime == nil || request.EndTime == "" {
		modifyData.EndTime = itemInfo.EndTime
	} else {
		modifyData.EndTime = request.EndTime
	}

	err = itemsModel.ModifyItem(modifyData)
	if err != nil {
		response.Code = constants.ERR_DB_UPDATE_DATA
		response.Message = constants.GetResponseMsg(constants.ERR_DB_UPDATE_DATA)
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}

	http_util.JsonResponse(c, http.StatusOK, response)
}

func GetItemList(c *gin.Context) {
	response := protocols.BaseResponse{Code: constants.SUCCESS, Message: constants.GetResponseMsg(constants.SUCCESS)}

	request := new(items.GetItemListRequest)
	err := c.BindJSON(&request)
	if err != nil || !request.IsValidParameter() {
		response.Code = constants.ERR_MSG_INVALID_PARAMETER
		response.Message = constants.GetResponseMsg(constants.ERR_MSG_INVALID_PARAMETER)
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}

	itemList, err := itemsModel.GetItemList(request.ChildrenIdx, request.Type, request.SearchDate)
	if err != nil {
		response.Code = constants.ERR_DB_NODATA
		response.Message = constants.GetResponseMsg(constants.ERR_DB_NODATA)
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}

	response.Data = itemList
	http_util.JsonResponse(c, http.StatusOK, response)
}

func DeleteItem(c *gin.Context) {
	response := protocols.BaseResponse{Code: constants.SUCCESS, Message: constants.GetResponseMsg(constants.SUCCESS)}

	request := new(items.DeleteItemRequest)
	err := c.BindJSON(&request)
	if err != nil || !request.IsValidParameter() {
		response.Code = constants.ERR_MSG_INVALID_PARAMETER
		response.Message = constants.GetResponseMsg(constants.ERR_MSG_INVALID_PARAMETER)
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}

	err = itemsModel.DeleteItem(request.Idx)
	if err != nil {
		response.Code = constants.ERR_DB_UPDATE_DATA
		response.Message = constants.GetResponseMsg(constants.ERR_DB_UPDATE_DATA)
		http_util.JsonResponse(c, http.StatusOK, response)
		return
	}

	http_util.JsonResponse(c, http.StatusOK, response)
}
