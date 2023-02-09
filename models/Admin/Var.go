package Admin

import (
	"database/sql"
	"simpleapp/database"
)

var DB *sql.DB = database.MakeConnection()
