package items

type InsertItemRequest struct {
	UserIdx     int         `json:"user_idx"`
	ChildrenIdx int         `json:"children_idx"`
	Type        string      `json:"type"`
	Etc1        interface{} `json:"etc1"`
	Etc2        interface{} `json:"etc2"`
	Etc3        interface{} `json:"etc3"`
	Etc4        interface{} `json:"etc4"`
	Etc5        interface{} `json:"etc5"`
	Etc6        interface{} `json:"etc6"`
	Etc7        interface{} `json:"etc7"`
	StartTime   string      `json:"start_time"`
	EndTime     string      `json:"end_time"`
}

func (request *InsertItemRequest) IsValidParameter() bool {
	if request.UserIdx <= 0 || request.ChildrenIdx <= 0 || request.StartTime == "" {
		return false
	}
	return true
}

type CompleteItemRequest struct {
	Idx int `json:"idx"`
}

func (request *CompleteItemRequest) IsValidParameter() bool {
	if request.Idx <= 0 {
		return false
	}
	return true
}

type ModifyItemRequest struct {
	Idx  int         `json:"idx"`
	Etc1 interface{} `json:"etc1"`
	Etc2 interface{} `json:"etc2"`
	Etc3 interface{} `json:"etc3"`
	Etc4 interface{} `json:"etc4"`
	Etc5 interface{} `json:"etc5"`
	Etc6 interface{} `json:"etc6"`
	Etc7 interface{} `json:"etc7"`
}

func (request *ModifyItemRequest) IsValidParameter() bool {
	if request.Idx <= 0 {
		return false
	}
	return true
}

type DeleteItemRequest struct {
	Idx int `json:"idx"`
}

func (request *DeleteItemRequest) IsValidParameter() bool {
	if request.Idx <= 0 {
		return false
	}
	return true
}

type GetItemListRequest struct {
	UserIdx     int    `json:"user_idx"`
	ChildrenIdx int    `json:"children_idx"`
	Type        string `json:"type"`
	SearchDate  string `json:"search_date"`
}

func (request *GetItemListRequest) IsValidParameter() bool {
	if request.UserIdx <= 0 || request.ChildrenIdx <= 0 {
		return false
	}
	return true
}
