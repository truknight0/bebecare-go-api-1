package children

type InsertChildrenRequest struct {
	Name     string      `json:"name"`
	Birthday string      `json:"birthday"`
	Gender   string      `json:"gender"`
	Tall     string      `json:"tall"`
	Weight   string      `json:"weight"`
	HeadSize string      `json:"head_size"`
	ImageUrl interface{} `json:"image_url"`
}

func (request *InsertChildrenRequest) IsValidParameter() bool {
	//log.DEBUG(fmt.Sprintf("%#v", request))
	if request.Name == "" ||
		request.Birthday == "" ||
		request.Gender == "" {
		return false
	}
	return true
}

type GetChildrenInfoRequest struct {
	Idx int `json:"idx"`
}

func (request *GetChildrenInfoRequest) IsValidParameter() bool {
	//log.DEBUG(fmt.Sprintf("%#v", request))
	if request.Idx == 0 {
		return false
	}
	return true
}

type ModifyChildrenRequest struct {
	Idx      int         `json:"idx"`
	Name     string      `json:"name"`
	Birthday string      `json:"birthday"`
	Gender   string      `json:"gender"`
	Tall     string      `json:"tall"`
	Weight   string      `json:"weight"`
	HeadSize string      `json:"head_size"`
	ImageUrl interface{} `json:"image_url"`
}

func (request *ModifyChildrenRequest) IsValidParameter() bool {
	//log.DEBUG(fmt.Sprintf("%#v", request))
	if request.Idx == 0 ||
		request.Name == "" ||
		request.Birthday == "" ||
		request.Gender == "" {
		return false
	}
	return true
}

type DeleteChildrenRequest struct {
	Idx int `json:"idx"`
}

func (request *DeleteChildrenRequest) IsValidParameter() bool {
	//log.DEBUG(fmt.Sprintf("%#v", request))
	if request.Idx == 0 {
		return false
	}
	return true
}
