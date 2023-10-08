package db_object

type InsertItem struct {
	UserIdx     int         `db:"user_idx"`
	ChildrenIdx int         `db:"children_idx"`
	Type        string      `db:"type"`
	Etc1        interface{} `db:"etc1"`
	Etc2        interface{} `db:"etc2"`
	Etc3        interface{} `db:"etc3"`
	Etc4        interface{} `db:"etc4"`
	Etc5        interface{} `db:"etc5"`
	Etc6        interface{} `db:"etc6"`
	Etc7        interface{} `db:"etc7"`
	StartTime   string      `db:"start_time"`
	EndTime     interface{} `db:"end_time"`
}

type ModifyItem struct {
	Idx       int         `db:"idx"`
	Etc1      interface{} `db:"etc1"`
	Etc2      interface{} `db:"etc2"`
	Etc3      interface{} `db:"etc3"`
	Etc4      interface{} `db:"etc4"`
	Etc5      interface{} `db:"etc5"`
	Etc6      interface{} `db:"etc6"`
	Etc7      interface{} `db:"etc7"`
	StartTime interface{} `db:"start_time"`
	EndTime   interface{} `db:"end_time"`
}

type GetItemList struct {
	Idx         int     `json:"idx" db:"idx"`
	UserIdx     int     `json:"user_idx" db:"user_idx"`
	ChildrenIdx int     `json:"children_idx" db:"children_idx"`
	Type        string  `json:"type" db:"type"`
	Name        *string `json:"name" db:"name"`
	Etc1        *string `json:"etc1" db:"etc1"`
	Etc2        *string `json:"etc2" db:"etc2"`
	Etc3        *string `json:"etc3" db:"etc3"`
	Etc4        *string `json:"etc4" db:"etc4"`
	Etc5        *string `json:"etc5" db:"etc5"`
	Etc6        *string `json:"etc6" db:"etc6"`
	Etc7        *string `json:"etc7" db:"etc7"`
	StartTime   *string `json:"start_time" db:"start_time"`
	EndTime     *string `json:"end_time" db:"end_time"`
}
