package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	urlConexao := "my_user:my_password@/my_database?charset=utf8&parseTime=True&loc=Local"
	fmt.Println("Conectando com o banco...")
	db, err := sql.Open("mysql", urlConexao)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Conexão realizada!")

	rows, err := db.Query("SELECT * FROM usuarios")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	
	fmt.Println(rows)
}
