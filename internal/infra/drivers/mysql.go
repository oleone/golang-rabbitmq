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

	fmt.Printf("Connecting to MySQL in address %s:%s \n", ipAddress, port)

	db, err := sql.Open("mysql", username+":"+password+"@tcp("+ipAddress+":"+port+")/"+dbName)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Connected to MySQL in %s:%s with success!\n", ipAddress, port)

	return &MySqlDriver{
		DB: db,
	}
}

func (d *MySqlDriver) Close() {
	d.DB.Close()
}
