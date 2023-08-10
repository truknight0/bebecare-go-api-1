package db_object

type GetInviteCodeInfo struct {
	InviteCode  int    `db:"invite_code"`
	UserName    string `db:"name"`
	ChildrenIdx int    `db:"children_idx"`
}

type RelInviteCodeAndUser struct {
	InviteCode int    `db:"invite_code"`
	UserIdx    int    `db:"user_idx"`
	UserName   string `db:"user_name"`
	UserRole   string `db:"user_role"`
}
