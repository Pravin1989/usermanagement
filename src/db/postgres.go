package db

import (
	"database/sql"
	"fmt"
	"log"
)

type DbConfig struct {
	Host     string `db:"host"`
	Port     int    `db:"port"`
	DBName   string `db:"dbname"`
	UserName string `db:"user"`
	Password string `db:"pssword"`
}
type PostgreDB struct {
	db *sql.DB
}

var dBConnection PostgreDB

func GetPostgresDBConnection() PostgreDB {
	return dBConnection
}

func LoadConnection() error {
	config := DbConfig{Host: "localhost", Port: 5432, DBName: "user_details", UserName: "pravin123", Password: "user123"}
	psqlInfo := fmt.Sprintf("host=%s, port=%d, user=%s, password=%s, dbname=%s, sslmode=disable", config.Host, config.Port, config.UserName, config.Password, config.DBName)
	db, connectionError := sql.Open("postgres", psqlInfo)
	if connectionError != nil {
		log.Fatal("Failed to connect to DB")
		return connectionError
	}
	dBConnection.db = db
	return nil
}
