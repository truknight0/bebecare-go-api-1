package inviteModel

import (
	"bebecare-go-api-1/beans/db_object"
	"bebecare-go-api-1/db"
	"bebecare-go-api-1/utils/log"
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

	_, err := db.DB.NamedExec(query, insertRequest)

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
		       us.name,
		       rpc.children_idx
		FROM invite_code AS ic
		LEFT JOIN user AS us ON ic.user_idx = us.idx
		LEFT JOIN rel_parent_children AS rpc ON rpc.user_idx = us.idx
		WHERE ic.invite_code = ?`

	err := db.DB.Get(inviteCodeInfo, query, inviteCode)

	if err != nil {
		log.ERROR(err.Error())
		return nil, err
	}
	return inviteCodeInfo, nil
}

func GetInviteCodeInfoWithUserIdx(userIdx int) (*db_object.GetInviteCodeInfo, error) {
	inviteCodeInfo := new(db_object.GetInviteCodeInfo)

	query := `
		SELECT ic.invite_code,
		       us.name,
		       rpc.children_idx
		FROM invite_code AS ic
		LEFT JOIN user AS us ON ic.user_idx = us.idx
		LEFT JOIN rel_parent_children AS rpc ON rpc.user_idx = us.idx
		WHERE us.idx = ?`

	err := db.DB.Get(inviteCodeInfo, query, userIdx)

	if err != nil {
		log.ERROR(err.Error())
		return nil, err
	}
	return inviteCodeInfo, nil
}
