package Admin

import (
	"database/sql"
	"simpleapp/core/database"
)

var DB *sql.DB = database.MakeConnection()
