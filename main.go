package main

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
	db, err := connection()
	defer db.Close()
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to PlanetScale!")

	users, err := readUsers(db)
	if err != nil {
		panic(err)
	}
	for _, u := range users {
		fmt.Println(u.ID, ", ", u.Email, ", ", u.FistName, ", ", u.LastName)
	}
}

func connection() (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", os.Getenv("DSN"))
	return db, err
}

type users struct {
	ID       int    `db:"id"`
	Email    string `db:"email"`
	FistName string `db:"first_name"`
	LastName string `db:"last_name"`
}

func readUsers(readDB *sqlx.DB) ([]users, error) {
	users := []users{}
	err := readDB.Select(&users, `select * from users`)
	if err != nil {
		return nil, err
	}
	return users, nil
}
