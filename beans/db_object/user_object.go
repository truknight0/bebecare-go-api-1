package db_object

import "github.com/jmoiron/sqlx"

type UserInfo struct {
	Name      string `db:"name" json:"name"`
	Phone     string `db:"phone" json:"phone"`
	Role      string `db:"role" json:"role"`
	CreatedAt string `db:"created_at" json:"created_at"`
}

type InsertUser struct {
	Trx   *sqlx.Tx
	Name  string `db:"name"`
	Phone string `db:"phone"`
	Role  string `db:"role"`
}
