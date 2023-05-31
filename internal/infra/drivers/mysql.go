package drivers

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type MySqlDriver struct {
	DB *sql.DB
}

func NewMySqlDriver(username string, password string, ipAddres string, port string, dbName string) *MySqlDriver {

	db, err := sql.Open("mysql", username+":"+password+"@tcp("+ipAddres+":"+port+")/"+dbName)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	return &MySqlDriver{
		DB: db,
	}
}

// func (d *MySqlDriver) Close() {
// 	d.DB.Close()
// }
