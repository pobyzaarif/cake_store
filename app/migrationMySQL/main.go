package main

import (
	"database/sql"
	"errors"
	"flag"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/pobyzaarif/cake_store/config"
	goLogger "github.com/pobyzaarif/go-logger/logger"
)

var logger = goLogger.NewLog("MIGRATION")

func main() {
	migrationPath := "app/migrationMySQL/migrations"
	cmdFlag := flag.String("cmd", "", "command for migration")
	flag.Parse()
	cmd := *cmdFlag

	conf := config.GetAPPConfig()
	dbCon := conf.GetDatabaseConnection()

	driver, _ := mysql.WithInstance(dbCon.MySQLDB, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file://"+migrationPath,
		conf.DBMySQLName,
		driver,
	)

	logger.SetTrackerID("migration")
	if cmd == "up" {
		if err := m.Steps(1); err != nil {
			if !errors.Is(err, migrate.ErrNoChange) {
				panic(err)
			}
		}

		logger.InfoWithData("migrate up", printVersion(dbCon.MySQLDB))
		return
	}

	if cmd == "down" {
		if err := m.Steps(-1); err != nil {
			if !errors.Is(err, migrate.ErrNoChange) {
				panic(err)
			}
		}

		logger.InfoWithData("migrate down", printVersion(dbCon.MySQLDB))
		return
	}

	if err := m.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			panic(err)
		}
	}

	logger.InfoWithData("migrate up to latest", printVersion(dbCon.MySQLDB))
}

func printVersion(db *sql.DB) map[string]interface{} {
	var version, dirty int
	db.QueryRow("SELECT * FROM schema_migrations LIMIT 1").Scan(&version, &dirty)

	return map[string]interface{}{"version": version, "dirty": dirty}
}
