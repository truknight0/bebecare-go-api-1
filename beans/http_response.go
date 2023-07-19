package beans

type BaseResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	//SessionId string      `json:"sessionId,omitempty"`
	Data interface{} `json:"data,omitempty"`
}
