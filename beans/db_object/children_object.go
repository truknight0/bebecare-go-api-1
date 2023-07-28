package db_object

import "github.com/jmoiron/sqlx"

type InsertChildren struct {
	Trx      *sqlx.Tx
	Name     string      `db:"name"`
	Birthday string      `db:"birthday"`
	Gender   string      `db:"gender"`
	ImageUrl interface{} `db:"image_url"`
}
