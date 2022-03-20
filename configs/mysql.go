package configs

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type database struct {
	sqlDB *sql.DB
}

var db *database

func InitMySql() error {
	db = new(database)
	connectionString := viper.GetString("MySqlDatabase.ConnectionString")
	sqlDb, err := sql.Open("mysql", connectionString)
	if err != nil {
		return errors.Wrap(err, "Error while connecting mysql database")
	}
	err = sqlDb.Ping()
	if err != nil {
		return errors.Wrap(err, "unable to ping database")
	}
	sqlDb.SetConnMaxIdleTime(3 * time.Minute)
	sqlDb.SetMaxOpenConns(10)
	sqlDb.SetMaxIdleConns(10)
	db.sqlDB = sqlDb
	return nil
}

func GetMySqlDB() *sql.DB {
	return db.sqlDB
}
