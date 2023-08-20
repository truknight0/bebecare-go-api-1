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

type GetUserListWithInviteCode struct {
	Idx      int    `json:"idx" db:"idx"`
	Name     string `json:"name" db:"name"`
	Phone    string `json:"phone" db:"phone"`
	Role     string `json:"role" db:"role"`
	UserType string `json:"user_type" db:"user_type"`
	IsMine   bool   `json:"is_mine" db:"is_mine"`
}
