package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var db *sql.DB

func GetConnection() *sql.DB {
	var err error

	if db == nil {
		log.Info().
			Str("Connection:", config.GetString("DB.Connection")).
			Int("Idle", config.GetInt("DB.MaxIdle")).
			Int("Max", config.GetInt("DB.MaxOpen")).
			Int("Lifetime (sec)", config.GetInt("DB.Lifetime")).
			Msg("MySQL")

		db, err = sql.Open("mysql", config.GetString("DB.Connection"))
		db.SetMaxIdleConns(config.GetInt("DB.MaxIdle"))
		db.SetMaxOpenConns(config.GetInt("DB.MaxOpen"))
		db.SetConnMaxLifetime(time.Second * config.GetDuration("DB.Lifetime"))

		if err != nil {
			log.Fatal().Err(err)
		}
	}

	return db
}

func initDB()  {
	config.SetDefault("DB.Connection", "root:@tcp(127.0.0.1:3306)/db?charset=utf8mb4")
	config.SetDefault("DB.MaxIdle", 1)
	config.SetDefault("DB.MaxOpen", 1)
	config.SetDefault("DB.Lifetime", 10)
	GetConnection()
}