package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

func AutoMigrate(db *sql.DB) error {
	log.Println("=> running auto migration")

	query := `
		CREATE TABLE IF NOT EXISTS domain (
			GlobalRank      INT,
			TldRank         INT,
			Domain          VARCHAR(255),
			TLD             VARCHAR(255),
			RefSubNets      INT,
			RefIPs          INT,
			IDN_Domain      VARCHAR(255),
			IDN_TLD         VARCHAR(255),
			PrevGlobalRank  INT,
			PrevTldRank     INT,
			PrevRefSubNets  INT,
			PrevRefIPs      INT
		)
	`

	_, err := db.ExecContext(context.Background(), query)
	if err != nil {
		return fmt.Errorf("failed to migrate table: %w", err)
	}

	log.Println("=> migration complete")
	return nil
}