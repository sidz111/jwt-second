package dbconfig

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() error {
	db_user := "root"
	db_pass := "root"
	db_host := "localhost"
	db_port := "3303"
	db_name := "jwt_db"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", db_user, db_pass, db_host, db_port, db_name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = db
	return nil
}
