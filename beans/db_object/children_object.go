package db_object

import "github.com/jmoiron/sqlx"

type InsertChildren struct {
	Trx      *sqlx.Tx
	UserIdx  int         `db:"user_idx"`
	Name     string      `db:"name"`
	Birthday string      `db:"birthday"`
	Gender   string      `db:"gender"`
	Tall     string      `db:"tall"`
	Weight   string      `db:"weight"`
	HeadSize string      `db:"head_size"`
	ImageUrl interface{} `db:"image_url"`
}

type ModifyChildren struct {
	Idx      int         `db:"idx"`
	UserIdx  int         `db:"user_idx"`
	Name     string      `db:"name"`
	Birthday string      `db:"birthday"`
	Gender   string      `db:"gender"`
	Tall     string      `db:"tall"`
	Weight   string      `db:"weight"`
	HeadSize string      `db:"head_size"`
	ImageUrl interface{} `db:"image_url"`
}

type GetUserChildrenInfo struct {
	Idx      int         `json:"idx" db:"idx"`
	Name     string      `json:"name" db:"name"`
	Birthday string      `json:"birthday" db:"birthday"`
	Gender   string      `json:"gender" db:"gender"`
	Tall     string      `json:"tall" db:"tall"`
	Weight   string      `json:"weight" db:"weight"`
	HeadSize string      `json:"head_size" db:"head_size"`
	ImageUrl interface{} `json:"image_url" db:"image_url"`
}

type DeleteChildren struct {
	Trx *sqlx.Tx
	Idx int `db:"idx"`
}
