package model

import (
	"bebecare-go-api-1/db"
	"bebecare-go-api-1/utils/log"
	"github.com/jmoiron/sqlx"
)

type AuthToken struct {
	Trx        *sqlx.Tx
	Idx        int    `db:"idx"`
	UserIdx    int    `db:"user_idx"`
	Token      string `db:"token"`
	ExpireDate string `db:"expire_date"`
	Name       string `db:"name"`
	Phone      string `db:"phone"`
	Role       string `db:"role"`
}

func SelectValidToken(authToken string) (*AuthToken, error) {
	t := new(AuthToken)
	err := db.DB.Get(t, `
		SELECT at.idx,
		       at.user_idx,
		       at.token,
		       at.expire_date,
		       us.name,
		       us.phone,
		       us.role
		FROM auth_token AS at
		LEFT JOIN user AS us ON at.user_idx = us.idx
		WHERE at.token = ?
        		AND now() < at.expire_date`, authToken)
	if err != nil {
		log.ERROR(err.Error())
		return nil, err
	}

	return t, nil
}
