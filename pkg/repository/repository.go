package repository

import (
	"fmt"

	"github.com/kiriwill/events-api/pkg/repository/mysql"
)

func New(connStr string, driver string) (*mysql.MysqlRepository, error) {
	switch driver {
	case "mysql", "mysqlx":
		return mysql.NewWithMigrate(connStr)
	}
	return nil, fmt.Errorf("could not find driver name: %s", driver)
}
