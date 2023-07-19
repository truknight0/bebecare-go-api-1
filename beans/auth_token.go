package beans

import (
	"bebecare-go-api-1/utils/log"
	"fmt"
)

type GetAuthTokenRequest struct {
	AccessKey    string `json:"accessKey"`
	AccessSecret string `json:"accessSecret"`
}

type GetAuthTokenResponse struct {
	BaseResponse
	AuthToken string `json:"authToken,omitempty"`
	ExpiredAt int64  `json:"expiredAt,omitempty"`
}

func (request *GetAuthTokenRequest) IsValidParameter() bool {
	log.DEBUG(fmt.Sprintf("%#v", request))
	if request.AccessKey == "" || request.AccessSecret == "" {
		return false
	}
	return true
}

func (request *GetAuthTokenRequest) ToString() string {
	return fmt.Sprintf("accessKey : %s, accessSecret : %s", request.AccessKey, request.AccessSecret)
}
