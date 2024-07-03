package postgres

import (
	"composition_service/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func ConnectionDb(cnf *config.Config) (*sql.DB, error) {
	conDB := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable ", "localhost", 5433, "postgres", "composition", "1111")
	db, err := sql.Open("postgres", conDB)
	if err != nil {
		return nil, err
	}
	return db, nil
}
