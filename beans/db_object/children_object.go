package db_object

import "github.com/jmoiron/sqlx"

type InsertChildren struct {
	Trx      *sqlx.Tx
	Name     string      `db:"name"`
	Birthday string      `db:"birthday"`
	Gender   string      `db:"gender"`
	Tall     string      `db:"tall"`
	Weight   string      `db:"weight"`
	HeadSize string      `db:"head_size"`
	ImageUrl interface{} `db:"image_url"`
}

type GetUserChildrenInfo struct {
	Idx      int    `db:"idx"`
	Name     string `db:"name"`
	Birthday string `db:"birthday"`
	Gender   string `db:"gender"`
	Tall     string `db:"tall"`
	Weight   string `db:"weight"`
	HeadSize string `db:"head_size"`
	ImageUrl string `db:"image_url"`
}
