package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMysql() error {
	var err error
	DB, err = gorm.Open(mysql.Open("root:yuling@tcp(127.0.0.1)/zg5?parseTime=true"))
	if err != nil {
		return err
	}
	return nil
}
