package dbcli

import (
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var mysqlCli *gorm.DB

func InitMySQL(user, pass, db, host string, port int) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, db)

	conn, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return errors.WithMessage(err, "Failed to connect to MySQL")
	}

	mysqlCli = conn

	return nil
}

func DB() *gorm.DB {
	return mysqlCli
}
