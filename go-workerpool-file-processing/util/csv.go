package util

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/yehezkiel1086/go-workerpool-file-processing/config"
)

func OpenCsvFile(cfg *config.Config) (*csv.Reader, *os.File, error) {
	log.Println("=> open csv file")

	f, err := os.Open(cfg.CSVFile)
	if err != nil {
		return nil, nil, err
	}

	reader := csv.NewReader(f)
	return reader, f, nil
}