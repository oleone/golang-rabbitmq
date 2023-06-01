package drivers

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type MySqlDriver struct {
	DB *sql.DB
}

func NewMySqlDriver(username string, password string, ipAddress string, port string, dbName string) *MySqlDriver {

	db, err := sql.Open("mysql", username+":"+password+"@tcp("+ipAddress+":"+port+")/"+dbName)

	if err != nil {
		panic(err)
	}

	fmt.Print("MySqlDriver connected\n")

	return &MySqlDriver{
		DB: db,
	}
}

func (d *MySqlDriver) Close() {
	d.DB.Close()
	fmt.Println("MySqlDriver closed")
}
