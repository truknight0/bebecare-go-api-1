package userModel

import (
	"bebecare-go-api-1/beans/db_object"
	"bebecare-go-api-1/db"
	"bebecare-go-api-1/utils/log"
	"fmt"
)

func GetUserInfo(name, phone, role string) (*db_object.UserInfo, error) {
	userInfo := new(db_object.UserInfo)

	err := db.DB.Get(userInfo, `
		SELECT	idx,
			name,
		    phone,
		    role,
		    DATE_FORMAT(created_at, '%Y-%m-%d %H:%i:%s') as created_at
		FROM
			user
		WHERE
			name = ?
			AND phone = ?
			AND role = ?`, name, phone, role)
	if err != nil {
		log.ERROR(err.Error())
		return nil, err
	}
	return userInfo, nil
}

func GetUserInfoWithToken(token string) (*db_object.UserInfoWithToken, error) {
	userInfo := new(db_object.UserInfoWithToken)

	query := `
		SELECT us.idx,
		       us.name,
		       us.phone,
		       us.role,
		       us.created_at,
		       at.token
		FROM auth_token AS at
		LEFT JOIN user AS us ON at.user_idx = us.idx
		WHERE at.token = ?
			AND DATE_FORMAT(at.expire_date, '%Y-%m-%d %H:%i:%s') >= NOW()`

	err := db.DB.Get(userInfo, query, token)

	if err != nil {
		log.ERROR(err.Error())
		return nil, err
	}
	return userInfo, nil
}

func GetUserTokenWithToken(token string) (string, error) {
	var getToken string
	query := `
		SELECT token 
		FROM auth_token
		WHERE token = ?
			AND DATE_FORMAT(expire_date, '%Y-%m-%d %H:%i:%s') >= NOW()`

	err := db.DB.Get(&getToken, query, token)

	if err != nil {
		log.ERROR(err.Error())
		return "", err
	}
	return getToken, nil
}

func GetUserToken(name, phone, role string) (string, error) {
	var getToken string
	query := `
		SELECT at.token 
		FROM auth_token AS at
		LEFT JOIN user AS mb ON at.user_idx = mb.idx
		WHERE mb.name = ?
			AND mb.phone = ?
			AND mb.role = ?
			AND DATE_FORMAT(at.expire_date, '%Y-%m-%d %H:%i:%s') >= NOW()`

	err := db.DB.Get(&getToken, query, name, phone, role)

	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return getToken, nil

}

func InsertUser(insertUser *db_object.InsertUser) (int, error) {
	query := `
		INSERT INTO user
		SET 
			name = :name,
			phone = :phone,
			role = :role
		ON DUPLICATE KEY UPDATE
			name = :name,
			phone = :phone,
			role = :role`

	r, err := db.DB.NamedExec(query, insertUser)

	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	insertId, _ := r.LastInsertId()
	return int(insertId), nil
}

func InsertToken(userIdx int, token string) error {
	query := `
		INSERT INTO auth_token
		SET
		    token = ?,
			user_idx = ?`
	_, err := db.DB.Exec(query, token, userIdx)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
