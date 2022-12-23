package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/eapache/go-resiliency/retrier"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBCon() (*gorm.DB, error) {
	var err error

	dsn := fmt.Sprintf("host=%v user=%v	password=%v dbname=%v port=%v sslmode=disable  TimeZone=Asia/Shanghai", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	r := retrier.New(retrier.ExponentialBackoff(10, 100*time.Millisecond), nil)
	var db *gorm.DB
	err = r.Run(func() error {
		dbc, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		db = dbc
		return err
	})

	return db, err
}

func InitDB() error {

	return Migrate()
}

func GetDBCon() (*gorm.DB, *sql.DB, error) {

	db, err := DBCon()
	if err != nil {
		return nil, nil, err
	}

	dbc, err := db.DB()
	if err != nil {
		return nil, nil, err
	}

	return db, dbc, nil
}
