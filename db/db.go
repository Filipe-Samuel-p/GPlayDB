package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	connStr := "host=localhost port=5432 user=filipe_Gplay dbname=GplayDB password=1234 sslmode=disable"
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Erro ao se conectar ao banco: ", err)
	}
	// defer DB.Close()

	if err2 := DB.Ping(); err2 != nil {
		log.Fatal("Banco não responde:", err)
	}

	log.Println("Conexão com o banco realizada com sucesso!")

}

func Close() {
	if DB != nil {
		DB.Close()
	}
}
