package user

//type GetUserInfoRequest struct {
//	Idx int `json:"idx"`
//}
//
//func (request *GetUserInfoRequest) IsValidParameter() bool {
//	//log.DEBUG(fmt.Sprintf("%#v", request))
//	if request.Idx <= 0 {
//		return false
//	}
//	return true
//}

type CheckUserRequest struct {
	Name        string `json:"name" db:"name"`
	Phone       string `json:"phone" db:"phone"`
	Role        string `json:"role" db:"role"`
	IsPushAgree int    `json:"is_push_agree" db:"is_push_agree"`
}

func (request *CheckUserRequest) IsValidParameter() bool {
	//log.DEBUG(fmt.Sprintf("%#v", request))
	if request.Name == "" || request.Phone == "" || request.Role == "" || request.IsPushAgree == 0 {
		return false
	}
	return true
}

type ModifyUserRequest struct {
	Idx         int    `json:"idx" db:"idx"`
	Name        string `json:"name" db:"name"`
	Phone       string `json:"phone" db:"phone"`
	Role        string `json:"role" db:"role"`
	IsPushAgree int    `json:"is_push_agree" db:"is_push_agree"`
}

func (request *ModifyUserRequest) IsValidParameter() bool {
	//log.DEBUG(fmt.Sprintf("%#v", request))
	if request.Idx == 0 || request.Name == "" || request.Phone == "" || request.Role == "" {
		return false
	}
	return true
}

type DisconnectUserRequest struct {
	Idx int `json:"idx" db:"idx"`
}

func (request *DisconnectUserRequest) IsValidParameter() bool {
	//log.DEBUG(fmt.Sprintf("%#v", request))
	if request.Idx == 0 {
		return false
	}
	return true
}
