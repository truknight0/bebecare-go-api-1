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

type GetUserInfoData struct {
	IsFirstUser bool        `json:"is_first_user"`
	Token       string      `json:"token"`
	Idx         int         `json:"idx"`
	Name        string      `json:"name"`
	Phone       string      `json:"phone"`
	Role        string      `json:"role"`
	IsPushAgree int         `json:"is_push_agree"`
	CreatedAt   string      `json:"created_at"`
	InviteCode  int         `json:"invite_code"`
	Children    interface{} `json:"children"`
}

type SetUserToken struct {
	IsFirstUser bool   `json:"isFirstUser"`
	Token       string `json:"token"`
}

type SetUserData struct {
	IsFirstUser bool        `json:"is_first_user"`
	Token       string      `json:"token"`
	Children    interface{} `json:"children"`
}
