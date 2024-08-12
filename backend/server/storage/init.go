package storage

import (
	"fmt"
	"momentum/utilities"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresDB struct {
	sqlx.DB
}

func DBConnectionString(settings utilities.DatabaseSettings) string {
	sslMode := "disable"
	if settings.RequireSSL {
		sslMode = "require"
	}

	return fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=%s dbname=%s", 
		settings.Host, settings.Port, settings.Username, settings.Password, sslMode, settings.Name)
}

func NewPostgresDB(settings utilities.DatabaseSettings) *PostgresDB {
	fmt.Println(DBConnectionString(settings))
	db := &PostgresDB{*sqlx.MustConnect("postgres", DBConnectionString(settings))}
	fmt.Printf("Listening to the database on host: %s, post: %s\n", settings.Host, settings.Port);
	return db
}
