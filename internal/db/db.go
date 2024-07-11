package db

import (
	"fmt"
	"open-user/internal/app/config"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func NewDb(conf *config.Config) (db *gorm.DB, err error) {
	// DB Connection
	url := os.Getenv("BRIMERCHANT_DBURL")
	usr := os.Getenv("BRIMERCHANT_DBUSER")
	name := os.Getenv("DATABASE_URL")
	pw := os.Getenv("BRIMERCHANT_DBPASSWORD")

	conn := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", usr, pw, url, name)
	fmt.Println("CONN ", conn)

	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", usr, pw, url, name))
	if err != nil {
		panic(err)
	}

	db.SingularTable(true)
	db.LogMode(true)

	return db, nil
}
