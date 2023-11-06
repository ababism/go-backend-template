package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"hw_3/internal/config"
)

func NewPostgresDB(cfg *config.Config) (*sqlx.DB, error) {
	fmt.Println("ST")
	db, err := sqlx.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			cfg.Postgres.Host,
			cfg.Postgres.Port,
			cfg.Postgres.Username,
			cfg.Postgres.DBName,
			cfg.Postgres.Password,
			cfg.Postgres.SSLMode))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("HI")
	return db, err
}
