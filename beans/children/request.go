package children

type InsertChildrenRequest struct {
	Name     string      `json:"name"`
	Birthday string      `json:"birthday"`
	Gender   string      `json:"gender"`
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
