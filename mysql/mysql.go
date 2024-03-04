package mysql

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func init() {
	fmt.Println("mysql package initialized")
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
}

func Connect() {
	db, err := sql.Open("mysql", "root:Be2xmmsq@tcp(localhost:3306)/demo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var version string
	err = db.QueryRow("SELECT VERSION()").Scan(&version)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to MySQL version:", version)
}
