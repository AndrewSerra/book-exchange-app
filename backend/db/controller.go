package db

import "database/sql"

type DBController struct {
	Database *sql.DB
}
