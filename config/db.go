package config

import (
	"fmt"

	"gorm.io/gorm/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"auto-rev/model"
)

var tables = []interface{}{
	&model.Packages{},
	&model.Subscriptions{},
	&model.SubscriptionPackages{},
	&model.Users{},
	&model.Roles{},
}

const (
	// host    = "localhost"
	// user    = "postgres"
	// pass    = "postgres"
	// dbname  = "bootest"
	// port    = 5432
	// sslmode = "disable"
)

func Conn() (*gorm.DB, error) {
	pg := fmt.Sprintf("root:sukmadjarna@tcp(127.0.0.1:3306)/auto-rev?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(mysql.Open(pg), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}

	fmt.Println("db flag")
	fmt.Println(db)

	return db, nil
}

func MigrateSchema(db *gorm.DB) {
	db.AutoMigrate(tables...)
}
