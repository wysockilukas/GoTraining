package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-oci8"
)

func main() {
	dsn := `ora/oracle@//localhost:1521/orcl`

	_, err := sql.Open("oci8", dsn)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("successfully connected to oracle server using dsn:", dsn)
}
