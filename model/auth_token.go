package model

import (
	"bebecare-go-api-1/db"
	"bebecare-go-api-1/utils/log"
	"github.com/jmoiron/sqlx"
	"time"
)

type AccessKey struct {
	Idx             int64     `db:"idx"`
	AccessKey       string    `db:"access_key"`
	AccessSecret    string    `db:"access_secret"`
	ApplicationName string    `db:"application_name"`
	IsValid         int       `db:"is_valid"`
	ExpiredAt       time.Time `db:"expired_at"`
	CreatedAt       time.Time `db:"created_at"`
}

type AuthToken struct {
	Trx             *sqlx.Tx
	Idx             int64     `db:"idx"`
	AccessKeyIdx    int64     `db:"access_key_idx"`
	Token           string    `db:"token"`
	Type            int       `db:"type"`
	IsValid         int       `db:"is_valid"`
	ApplicationName string    `db:"application_name"`
	ExpiredAt       time.Time `db:"expired_at"`
	CreatedAt       time.Time `db:"created_at"`
}

func SelectAccessKey(accessKey, accessSecret string) (*AccessKey, error) {
	k := new(AccessKey)
	err := db.DB.Get(k, `
		SELECT idx,
        	access_key,
        	access_secret,
        	application_name,
        	is_valid,
        	expired_at,
        	created_at
		FROM ACCESS_KEY_T
		WHERE access_key = ?
        		AND access_secret = ?
        		AND is_valid = 1
        		AND now() < expired_at`, accessKey, accessSecret)
	if err != nil {
		log.ERROR(err.Error())
		return nil, err
	}

	return k, nil
}

func SelectValidTokenFromAccessKey(accessKeyIdx int64) (*AuthToken, error) {
	t := new(AuthToken)
	err := db.DB.Get(t, `
		SELECT idx,           
		    access_key_idx,
		    token,         
        	IFNULL(type,0) as type,          
        	is_valid,      
        	expired_at,    
        	created_at        	
		FROM AUTH_TOKEN_T
		WHERE access_key_idx = ?
        		AND is_valid = 1
        		AND now() < expired_at`, accessKeyIdx)
	if err != nil {
		log.ERROR(err.Error())
		return nil, err
	}

	return t, nil
}

func SelectValidToken(authToken string) (*AuthToken, error) {
	t := new(AuthToken)
	err := db.DB.Get(t, `
		SELECT AT.idx,           
		    AT.access_key_idx,
		    AT.token,         
        	IFNULL(AT.type,0) as type,          
        	AT.is_valid,
		    AK.application_name,
        	AT.expired_at,    
        	AT.created_at        	
		FROM AUTH_TOKEN_T AT
		INNER JOIN ACCESS_KEY_T AK ON AT.access_key_idx = AK.idx
		WHERE AT.token = ?
        		AND AT.is_valid = 1
        		AND now() < AT.expired_at`, authToken)
	if err != nil {
		log.ERROR(err.Error())
		return nil, err
	}

	return t, nil
}

func (authToken *AuthToken) InsertAuthToken() error {
	r, err := authToken.Trx.NamedExec(`
		INSERT INTO AUTH_TOKEN_T
		SET 
			access_key_idx = :access_key_idx,
			token = :token,
			expired_at = :expired_at`, authToken)
	if err != nil {
		log.ERROR(err.Error())
		return err
	}

	idx, _ := r.LastInsertId()
	authToken.Idx = idx

	return nil
}

func (authToken *AuthToken) InvalidToken() error {
	_, err := authToken.Trx.NamedExec(`
		UPDATE AUTH_TOKEN_T
		SET 
		    is_valid = 0
		WHERE
			access_key_idx = :access_key_idx
			AND is_valid = 1`, authToken)
	if err != nil {
		log.ERROR(err.Error())
		return err
	}

	return nil
}
