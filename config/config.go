package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type (
	Config struct {
		AppMainPort      string
		AppMaxRowPerPage string
		DBDriver         string
		// mysql
		DBMySQLHost     string
		DBMySQLPort     string
		DBMySQLUser     string
		DBMySQLPassword string
		DBMySQLName     string
	}
)

//DatabaseConnection Database connection
type DatabaseConnection struct {
	MySQLDB *sql.DB
}

func GetAPPConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	return &Config{
		AppMainPort:      os.Getenv("CAKE_STORE_APP_MAIN_PORT"),
		AppMaxRowPerPage: os.Getenv("CAKE_STORE_APP_MAX_ROW_PER_PAGE"),
		DBDriver:         os.Getenv("CAKE_STORE_DB_DRIVER"),
		DBMySQLHost:      os.Getenv("CAKE_STORE_DB_MYSQL_HOST"),
		DBMySQLPort:      os.Getenv("CAKE_STORE_DB_MYSQL_PORT"),
		DBMySQLUser:      os.Getenv("CAKE_STORE_DB_MYSQL_USER"),
		DBMySQLPassword:  os.Getenv("CAKE_STORE_DB_MYSQL_PASS"),
		DBMySQLName:      os.Getenv("CAKE_STORE_DB_MYSQL_NAME"),
	}
}

func (conf *Config) GetDatabaseConnection() *DatabaseConnection {
	var db DatabaseConnection
	if conf.DBDriver == "mysql" {
		uri := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true&&multiStatements=true",
			conf.DBMySQLUser,
			conf.DBMySQLPassword,
			conf.DBMySQLHost,
			conf.DBMySQLPort,
			conf.DBMySQLName,
		)
		dbMySQL, err := sql.Open("mysql", uri)
		if err != nil {
			log.Fatal(err)
		}

		dbMySQL.SetConnMaxLifetime(time.Minute * 3)
		dbMySQL.SetMaxOpenConns(10)
		dbMySQL.SetMaxIdleConns(10)

		db.MySQLDB = dbMySQL

		return &db
	}

	log.Fatal("unsupported driver")

	return nil
}
