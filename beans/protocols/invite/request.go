package invite

type MakeInviteCodeRequest struct {
	ChildrenIdx int `json:"children_idx"`
}

func (request *MakeInviteCodeRequest) IsValidParameter() bool {
	if request.ChildrenIdx <= 0 {
		return false
	}
	return true
}

type JoinInviteCodeRequest struct {
	InviteCode int `json:"invite_code"`
}

func (request *JoinInviteCodeRequest) IsValidParameter() bool {
	if request.InviteCode <= 0 {
		return false
	}
	return true
}
