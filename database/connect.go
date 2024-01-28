package database

import (
	"fmt"
	"nasi-golang/config"
	"nasi-golang/model"
	"time"

	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	// DBConn is a pointer to gorm.DB
	DBConn *gorm.DB
)

func Connect() {
	var (
		user     = config.Config("DB_USER")
		password = config.Config("DB_PASSWORD")
		host     = config.Config("DB_HOST")
		dbname   = config.Config("DB_NAME")
	)
	dsn := fmt.Sprintf("clickhouse://%s/%s?username=%s&password=%s&read_timeout=10s", host, dbname, user, password)

	db, err := gorm.Open(clickhouse.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.Product{})

	if sqlDB, err := db.DB(); err == nil {
		// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
		sqlDB.SetMaxIdleConns(10)

		// SetMaxOpenConns sets the maximum number of open connections to the database.
		sqlDB.SetMaxOpenConns(100)

		// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
		sqlDB.SetConnMaxLifetime(time.Hour)
	}

	DBConn = db
}
