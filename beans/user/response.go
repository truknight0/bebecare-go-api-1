package user

import (
	"bebecare-go-api-1/beans"
)

type GetUserResponse struct {
	beans.BaseResponse
}

type CheckUserResponse struct {
	beans.BaseResponse
}

type SetUserToken struct {
	IsFirstUser bool   `json:"isFirstUser"`
	Token       string `json:"token"`
}
