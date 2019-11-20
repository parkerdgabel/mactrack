package main

import (
	"os"
)

const (
	defaultDatabasePath = "~/.config/mactrack/mactrack.db"
)

func getDatabasePath() string {
	dbpath, ok := os.LookupEnv("MACTRACK_DATABASE")
	if !ok {
		dbpath = defaultDatabasePath
	}
	return dbpath
}

func main() {
}
