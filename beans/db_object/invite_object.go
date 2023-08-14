package db_object

import "github.com/jmoiron/sqlx"

type GetInviteCodeInfo struct {
	InviteCode int    `db:"invite_code"`
	UserName   string `db:"name"`
	UserIdx    int    `db:"user_idx"`
}

type RelInviteCodeAndUser struct {
	Trx        *sqlx.Tx
	InviteCode int    `db:"invite_code"`
	UserIdx    int    `db:"user_idx"`
	UserName   string `db:"user_name"`
	UserRole   string `db:"user_role"`
}
