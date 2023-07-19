package user

type GetUserRequest struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Role  string `json:"role"`
}

func (request *GetUserRequest) IsValidParameter() bool {
	//log.DEBUG(fmt.Sprintf("%#v", request))
	if request.Name == "" || request.Phone == "" || request.Role == "" {
		return false
	}
	return true
}

type CheckUserRequest struct {
	Name  string `json:"name" db:"name"`
	Phone string `json:"phone" db:"phone"`
	Role  string `json:"role" db:"role"`
}

func (request *CheckUserRequest) IsValidParameter() bool {
	//log.DEBUG(fmt.Sprintf("%#v", request))
	if request.Name == "" || request.Phone == "" || request.Role == "" {
		return false
	}
	return true
}
