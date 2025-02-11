package core

import (
	"fmt"
	"log"
	"os"
	"database/sql"
	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)

type Conn_MySQL struct {
	DB *sql.DB
	Err string
}

func GetDBPool() *Conn_MySQL {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
	}
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbSchema := os.Getenv("DB_NAME")

dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPass, dbHost, dbSchema)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error al abrir la base de datos: %v", err)
	}
	return &Conn_MySQL{DB: db, Err: ""}
}
