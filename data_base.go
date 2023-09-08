package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func database() {
	// Строка подключения к базе данных PostgreSQL//****** данные БД (user и dbname)
	connStr := "user=******* dbname=****** sslmode=disable"

	// Открываем соединение с базой данных PostgreSQL
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Проверяем, что соединение работает
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Подключение к PostgreSQL успешно!")

}

