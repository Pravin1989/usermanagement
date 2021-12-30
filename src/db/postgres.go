package db

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DbConfig struct {
	Host     string `db:"host"`
	Port     int    `db:"port"`
	DBName   string `db:"dbname"`
	UserName string `db:"user"`
	Password string `db:"pssword"`
}
type PostgreDB struct {
	DB *sqlx.DB
}

var dBConnection PostgreDB

func GetPostgresDBConnection() PostgreDB {
	return dBConnection
}

func LoadConnection() error {
	config := DbConfig{Host: "database", Port: 5432, DBName: "usermanagement", UserName: "pravin123", Password: "user123"}
	psqlInfo := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		config.UserName,
		config.Password,
		config.Host,
		config.Port,
		config.DBName)
	db, connectionError := sqlx.Connect("postgres", psqlInfo)
	if connectionError != nil {
		log.Fatal("Failed to connect to DB", connectionError)
		return connectionError
	}
	dBConnection.DB = db
	return nil
}
