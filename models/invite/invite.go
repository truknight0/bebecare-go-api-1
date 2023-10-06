package inviteModel

import (
	"bebecare-go-api-1/beans/db_object"
	"bebecare-go-api-1/db"
	"bebecare-go-api-1/utils/log"
	"fmt"
)

func DupInviteCode(inviteCode int) (bool, error) {
	var code int
	query := `
		SELECT invite_code
		FROM invite_code
		WHERE invite_code = ?`

	err := db.DB.Get(&code, query, inviteCode)

	// 중복된 데이터가 없으면 err가 됨
	if err != nil {
		return false, nil
	}

	if code == inviteCode {
		return true, nil
	}

	return false, nil
}

func InsertInviteCode(inviteCode, userIdx int) error {
	query := `
		INSERT INTO invite_code
		SET
		    invite_code = ?,
		    user_idx = ?`

	_, err := db.DB.Exec(query, inviteCode, userIdx)

	if err != nil {
		log.ERROR(err.Error())
		return err
	}

	return nil
}

func RelInviteCodeAndUser(insertRequest *db_object.RelInviteCodeAndUser) error {
	query := `
		INSERT INTO rel_invite_code_user
		SET 
			invite_code = :invite_code,
			user_idx = :user_idx,
			user_name = :user_name,
			user_role = :user_role`

	_, err := insertRequest.Trx.NamedExec(query, insertRequest)

	if err != nil {
		log.ERROR(err.Error())
		return err
	}
	return nil
}

func RelInviteUserAndChildren(userIdx int, childrenList []db_object.GetUserChildrenInfo) error {
	query := `
		INSERT INTO rel_parent_children(user_idx, children_idx) VALUES `

	var vals string
	for _, row := range childrenList {
		var multiInsertStr = "(%d, %d),"
		vals += fmt.Sprintf(multiInsertStr, userIdx, row.Idx)
	}
	query += vals
	query = query[0 : len(query)-1]

	_, err := db.DB.Exec(query)

	if err != nil {
		log.ERROR(err.Error())
		return err
	}
	return nil
}

func GetInviteCodeInfo(inviteCode int) (*db_object.GetInviteCodeInfo, error) {
	inviteCodeInfo := new(db_object.GetInviteCodeInfo)

	query := `
		SELECT ic.invite_code,
		       ic.user_idx,
		       us.name
		FROM invite_code AS ic
		LEFT JOIN user AS us ON ic.user_idx = us.idx
		WHERE ic.invite_code = ?`

	err := db.DB.Get(inviteCodeInfo, query, inviteCode)

	if err != nil {
		log.ERROR(err.Error())
		return nil, err
	}
	return inviteCodeInfo, nil
}

func CheckInviteCodeMaker(inviteCode, userIdx int) (int, error) {
	var checkUserIdx int

	query := `
		SELECT user_idx
		FROM invite_code
		WHERE invite_code = ?
			AND user_idx = ?`

	err := db.DB.Get(&checkUserIdx, query, inviteCode, userIdx)

	if err != nil {
		log.ERROR(err.Error())
		return 0, err
	}
	return checkUserIdx, nil
}

func GetInviteCodeInfoWithUserIdx(userIdx int) (*db_object.GetInviteCodeInfo, error) {
	inviteCodeInfo := new(db_object.GetInviteCodeInfo)

	query := `
		SELECT ric.invite_code,
		       ric.user_idx,
		       us.name
		FROM rel_invite_code_user AS ric
		LEFT JOIN user AS us ON ric.user_idx = us.idx
		WHERE us.idx = ?`

	err := db.DB.Get(inviteCodeInfo, query, userIdx)

	if err != nil {
		log.ERROR(err.Error())
		return nil, err
	}
	return inviteCodeInfo, nil
}

func GetUserListWithInviteCode(token string, inviteCode int) ([]db_object.GetUserListWithInviteCode, error) {
	var userList []db_object.GetUserListWithInviteCode

	query := `
		SELECT us.idx,
			us.name,
			us.phone,
			us.role,
			us.user_type,
			IF (at.token = ?, true, false) AS is_mine
		FROM rel_invite_code_user AS ric
		LEFT JOIN user AS us ON ric.user_idx = us.idx
		LEFT JOIN auth_token AS at ON at.user_idx = us.idx
		WHERE ric.invite_code = ?`

	err := db.DB.Select(&userList, query, token, inviteCode)

	if err != nil {
		log.ERROR(err.Error())
		return nil, err
	}
	return userList, nil
}
