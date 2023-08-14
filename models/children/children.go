package childrenModel

import (
	"bebecare-go-api-1/beans/db_object"
	"bebecare-go-api-1/db"
	"bebecare-go-api-1/utils/log"
	"fmt"
)

func InsertChildren(insertRequest *db_object.InsertChildren) (int, error) {
	query := `
		INSERT INTO children
		SET 
			user_idx = :user_idx,
			name = :name,
			birthday = :birthday,
			gender = :gender,
			tall = :tall,
			weight = :weight,
			head_size = :head_size,
			image_url = :image_url`

	res, err := insertRequest.Trx.NamedExec(query, insertRequest)

	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	insertId, _ := res.LastInsertId()
	return int(insertId), nil
}

func InsertRelParentChildren(userIdx, childrenIdx int) error {
	query := `
		INSERT INTO rel_parent_children
		SET
		    user_idx = ?,
		    children_idx = ?`

	_, err := db.DB.Exec(query, userIdx, childrenIdx)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func GetUserChildrenList(userIdx int) ([]db_object.GetUserChildrenInfo, error) {
	var childrenInfo []db_object.GetUserChildrenInfo

	query := `
		SELECT cr.idx,
		    cr.name,
			DATE_FORMAT(cr.birthday, '%Y-%m-%d') as birthday,
			cr.gender,
			cr.tall,
			cr.weight,
			cr.head_size,
			cr.image_url
		FROM rel_parent_children AS rpc
		LEFT JOIN children AS cr ON rpc.children_idx = cr.idx
		WHERE rpc.user_idx = ?
		ORDER BY idx ASC`

	err := db.DB.Select(&childrenInfo, query, userIdx)

	if err != nil {
		log.ERROR(err.Error())
		return nil, err
	}
	return childrenInfo, nil
}

func GetChildrenCount(userIdx int) (int, error) {
	var count int

	query := `
		SELECT COUNT(*) AS COUNT
		FROM rel_parent_children
		WHERE user_idx = ?`

	err := db.DB.Get(&count, query, userIdx)

	if err != nil {
		log.ERROR(err.Error())
		return 0, err
	}
	return count, nil

}
