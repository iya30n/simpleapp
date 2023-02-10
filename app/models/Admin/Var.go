package Admin

import (
	"database/sql"
	"simpleapp/app/database"
)

var DB *sql.DB = database.MakeConnection()
