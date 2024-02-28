package app

import "github.com/BarnabyCharles/zg5week1/mysql"

func Init(str ...string) error {
	var err error
	for _, val := range str {
		switch val {
		case "mysql":
			err = mysql.InitMysql()
		}
	}
	return err
}
