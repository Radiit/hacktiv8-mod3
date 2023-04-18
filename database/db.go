package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sesi4/model"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "pswd"
	dbname   = "hacktiv8"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	config := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = gorm.Open(postgres.Open(config))
	if err != nil {
		panic(err)
	}

	sqlDb, err := db.DB()
	if err != nil {
		panic(err)
	}

	err = sqlDb.Ping()
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.ItemBook{})
}

func GetDB() *gorm.DB {

	return db
}
