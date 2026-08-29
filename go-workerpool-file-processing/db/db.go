package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/yehezkiel1086/go-workerpool-file-processing/config"
)

func OpenDbConnection(cfg *config.Config) (*sql.DB, error) {
	log.Println("=> open db connection")

	db, err := sql.Open("postgres", cfg.DBConnString())
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.DBMaxConns)
	db.SetMaxIdleConns(cfg.DBMaxIdleConns)

	return db, nil
}