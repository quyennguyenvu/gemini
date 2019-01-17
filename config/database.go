package config

import (
	"os"
	"sync"
)

// DBConnection ..
type DBConnection struct {
	DataSource string
	Database   string
}

var confConnection *DBConnection
var dbInit sync.Once

// GetDBConnection ..
func GetDBConnection() *DBConnection {
	dbInit.Do(func() {
		drive := os.Getenv("DB_CONNECTION")
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		database := os.Getenv("DB_DATABASE")
		userName := os.Getenv("DB_USERNAME")
		password := os.Getenv("DB_PASSWORD")

		// select datasource
		dataSource := drive + "://" +
			userName + ":" + password +
			"@" + dbHost + ":" + dbPort

		confConnection = &DBConnection{
			DataSource: dataSource,
			Database:   database,
		}
		println(dataSource)
	})
	return confConnection
}
