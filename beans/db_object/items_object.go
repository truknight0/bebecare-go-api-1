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
	Idx  int         `db:"idx"`
	Etc1 interface{} `db:"etc1"`
	Etc2 interface{} `db:"etc2"`
	Etc3 interface{} `db:"etc3"`
	Etc4 interface{} `db:"etc4"`
	Etc5 interface{} `db:"etc5"`
	Etc6 interface{} `db:"etc6"`
	Etc7 interface{} `db:"etc7"`
}

type GetItemList struct {
	Idx         int         `db:"idx"`
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
	EndTime     string      `db:"end_time"`
}
