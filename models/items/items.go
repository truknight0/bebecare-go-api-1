package itemsModel

import (
	"bebecare-go-api-1/beans/db_object"
	"bebecare-go-api-1/db"
	"bebecare-go-api-1/utils/log"
	"fmt"
)

func InsertItem(insertRequest *db_object.InsertItem) error {
	query := `
		INSERT INTO items
		SET 
			user_idx = :user_idx,
			children_idx = :children_idx,
			type = :type,
			etc1 = :etc1,
			etc2 = :etc2,
			etc3 = :etc3,
			etc4 = :etc4,
			etc5 = :etc5,
			etc6 = :etc6,
			etc7 = :etc7,
			start_time = :start_time,
			end_time = :end_time`

	_, err := db.DB.NamedExec(query, insertRequest)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func ModifyItem(modifyRequest *db_object.ModifyItem) error {
	query := `
		UPDATE items
		SET 
			etc1 = :etc1,
			etc2 = :etc2,
			etc3 = :etc3,
			etc4 = :etc4,
			etc5 = :etc5,
			etc6 = :etc6,
			etc7 = :etc7,
			start_time = :start_time,
			end_time = :end_time
		WHERE
		    idx = :idx`

	_, err := db.DB.NamedExec(query, modifyRequest)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func CompleteItem(idx int) error {
	query := `
		UPDATE items
		SET
		    end_time = NOW()
		WHERE
		    idx = ?`

	_, err := db.DB.Exec(query, idx)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func GetItemInfo(idx int) (*db_object.GetItemList, error) {
	itemInfo := new(db_object.GetItemList)

	query := `
		SELECT idx,
			user_idx,
			children_idx,
			type,
			CASE 
			    WHEN (type = 'A') THEN '모유'
			    WHEN (type = 'B') THEN '분유'
			    WHEN (type = 'C') THEN '이유식'
			    WHEN (type = 'D') THEN '소변'
			    WHEN (type = 'E') THEN '대변'
			    WHEN (type = 'F') THEN '낮잠'
			    WHEN (type = 'G') THEN '밤잠'
			    ELSE NULL
			END AS name,
			etc1,
			etc2,
			etc3,
			etc4,
			etc5,
			etc6,
			etc7,
			DATE_FORMAT(start_time, '%Y-%m-%d %H:%i:%s') AS start_time,
			DATE_FORMAT(end_time, '%Y-%m-%d %H:%i:%s') AS end_time
		FROM  items
		WHERE
		    idx = ?`

	err := db.DB.Get(itemInfo, query, idx)

	if err != nil {
		log.ERROR(err.Error())
		return nil, err
	}
	return itemInfo, nil
}

func GetItemList(childrenIdx int, itemType string, searchDate string) ([]db_object.GetItemList, error) {
	var itemList []db_object.GetItemList

	var typeQuery string
	var searchDateQuery = ` AND DATE_FORMAT(start_time, '%Y-%m-%d') = DATE_FORMAT(NOW(), '%Y-%m-%d')`

	if itemType != "" {
		typeQuery = ` AND type = '` + itemType + `'`
	}
	if searchDate != "" {
		searchDateQuery = ` AND DATE_FORMAT(start_time, '%Y-%m-%d') = '` + searchDate + `'`
	}

	query := `
		SELECT idx,
			user_idx,
			children_idx,
			type,
			CASE 
			    WHEN (type = 'A') THEN '모유'
			    WHEN (type = 'B') THEN '분유'
			    WHEN (type = 'C') THEN '이유식'
			    WHEN (type = 'D') THEN '소변'
			    WHEN (type = 'E') THEN '대변'
			    WHEN (type = 'F') THEN '낮잠'
			    WHEN (type = 'G') THEN '밤잠'
			    ELSE NULL
			END AS name,
			etc1,
			etc2,
			etc3,
			etc4,
			etc5,
			etc6,
			etc7,
			DATE_FORMAT(start_time, '%Y-%m-%d %H:%i:%s') AS start_time,
			DATE_FORMAT(end_time, '%Y-%m-%d %H:%i:%s') AS end_time
		FROM  items
		WHERE
		    children_idx = ?` + typeQuery + searchDateQuery + `
		ORDER BY start_time DESC, idx DESC`

	err := db.DB.Select(&itemList, query, childrenIdx)

	if err != nil {
		log.ERROR(err.Error())
		return nil, err
	}
	return itemList, nil
}

func DeleteItem(idx int) error {
	query := `
		DELETE FROM items
		WHERE
		    idx = ?`

	_, err := db.DB.Exec(query, idx)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
