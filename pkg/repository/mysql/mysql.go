package mysql

import (
	"github.com/kiriwill/events-api/pkg/service"
	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

func NewConn(connStr string) (*MysqlRepository, error) {
	mysqlDb := mysql.Open(connStr)

	db, err := gorm.Open(mysqlDb, &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return &MysqlRepository{Db: db, ConnectionString: connStr}, nil
}

func NewWithMigrate(connStr string) (*MysqlRepository, error) {
	repo, err := NewConn(connStr)
	if err != nil {
		return nil, err
	}

	err = repo.Db.AutoMigrate(&service.Event{}, &service.User{})
	if err != nil {
		return nil, err
	}

	return repo, nil
}

type MysqlRepository struct {
	Db               *gorm.DB
	ConnectionString string
}
