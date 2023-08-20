package db_object

import "github.com/jmoiron/sqlx"

type UserInfo struct {
	Idx         int    `db:"idx" json:"idx"`
	Name        string `db:"name" json:"name"`
	Phone       string `db:"phone" json:"phone"`
	Role        string `db:"role" json:"role"`
	IsPushAgree int    `db:"is_push_agree"`
	CreatedAt   string `db:"created_at" json:"created_at"`
}

type UserInfoWithToken struct {
	Idx         int    `db:"idx" json:"idx"`
	Name        string `db:"name" json:"name"`
	Phone       string `db:"phone" json:"phone"`
	Role        string `db:"role" json:"role"`
	IsPushAgree int    `db:"is_push_agree" json:"is_push_agree"`
	UserType    string `db:"user_type" json:"user_type"`
	Token       string `db:"token" json:"token"`
	CreatedAt   string `db:"created_at" json:"created_at"`
}

type InsertUser struct {
	Trx         *sqlx.Tx
	Name        string `db:"name"`
	Phone       string `db:"phone"`
	Role        string `db:"role"`
	IsPushAgree int    `db:"is_push_agree"`
}

type ModifyUser struct {
	Trx         *sqlx.Tx
	Idx         int    `db:"idx"`
	Name        string `db:"name"`
	Phone       string `db:"phone"`
	Role        string `db:"role"`
	IsPushAgree int    `db:"is_push_agree"`
}

type DisconnectUser struct {
	Trx *sqlx.Tx
	Idx int `db:"idx"`
}
