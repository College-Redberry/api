package postgres

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	fmt.Print(os.Getenv("DB_STRING"))
	connection, err := sql.Open("postgres", os.Getenv("DB_STRING"))

	if err != nil {
		panic(err)
	}

	err = connection.Ping()
	if err != nil {
		panic(err)
	}
	return connection
}
