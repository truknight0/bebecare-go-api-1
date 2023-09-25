package user

import (
	"bebecare-go-api-1/beans/protocols"
)

type GetUserResponse struct {
	protocols.BaseResponse
}

type CheckUserResponse struct {
	protocols.BaseResponse
}

type GetUserInfoData struct {
	IsFirstUser bool        `json:"is_first_user"`
	Token       string      `json:"token"`
	Idx         int         `json:"idx"`
	Name        string      `json:"name"`
	Phone       string      `json:"phone"`
	Role        string      `json:"role"`
	IsPushAgree int         `json:"is_push_agree"`
	UserType    string      `json:"user_type"`
	CreatedAt   string      `json:"created_at"`
	InviteCode  interface{} `json:"invite_code"`
	Children    interface{} `json:"children"`
	Parents     interface{} `json:"parents"`
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
