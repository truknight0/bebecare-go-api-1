package db

import (
	"bebecare-go-api-1/utils/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"time"
)

var DB *sqlx.DB

func init() {
	initDb()
}

func initDb() {
	//dbDriver := config.Cfg.DBDriver
	dbHost := config.GetString("db.host")
	dbPort := config.GetIntDefault("db.port", 3306)
	dbUser := config.GetString("db.user")
	dbPass := config.GetString("db.pass")
	dbName := config.GetString("db.name")

	dbConn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", dbUser, dbPass, dbHost, dbPort, dbName)
	var err error
	DB, err = sqlx.Open("mysql", dbConn)
	if err != nil {
		//log.ERROR(err.Error())
		panic(err.Error())
	}

	ch1 := make(chan bool)
	go func(done chan bool) {
		for {
			time.Sleep(time.Minute * 5)
			done <- true
		}
	}(ch1)

	go func(done chan bool) {
		for {
			select {
			case <-done:
				err := DB.Ping()
				if err != nil {
					fmt.Println(err.Error())
				}
			}
		}
	}(ch1)
}

func Closed() {
	DB.Close()
}
