package db

import (
	"fmt"

	"agrha/open-user/internal/app/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func NewDb() (db *gorm.DB, err error) {

	// DB Connection
	url := config.BRIMerchantDbURL
	usr := config.BRIMerchantDbUsername
	pw := config.BRIMerchantDbPassword
	name := config.BRIMerchantDbName

	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", usr, pw, url, name))
	if err != nil {
		panic(err)
	}

	db.SingularTable(true)
	db.LogMode(true)

	return db, nil
}
