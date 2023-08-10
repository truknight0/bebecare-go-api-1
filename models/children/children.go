package childrenModel

import (
	"bebecare-go-api-1/beans/db_object"
	"bebecare-go-api-1/db"
	"fmt"
)

func InsertChildren(insertRequest *db_object.InsertChildren) (int, error) {
	query := `
		INSERT INTO children
		SET 
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
