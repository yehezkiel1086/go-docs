package service

import (
	"context"
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"
	"sync"
)

var dataHeaders = make([]string, 0)

func DispatchWorkers(ctx context.Context, db *sql.DB, jobs <-chan []interface{}, wg *sync.WaitGroup, totalWorker int) {
	for workerIndex := 0; workerIndex <= totalWorker; workerIndex++ {
		go func(workerIndex int) {
			counter := 0
			for job := range jobs {
				// drop job silently if context is cancelled
				if ctx.Err() != nil {
					wg.Done()
					continue
				}
				doTheJob(ctx, workerIndex, counter, db, job)
				wg.Done()
				counter++
			}
		}(workerIndex)
	}
}

func ReadCsvFilePerLineThenSendToWorker(ctx context.Context, csvReader *csv.Reader, jobs chan<- []interface{}, wg *sync.WaitGroup) {
	for {
		// stop reading if context is cancelled
		if ctx.Err() != nil {
			log.Println("=> stopping csv read due to shutdown")
			break
		}

		row, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}

		if len(dataHeaders) == 0 {
			dataHeaders = row
			continue
		}

		rowOrdered := make([]interface{}, 0)
		for _, each := range row {
			rowOrdered = append(rowOrdered, each)
		}

		wg.Add(1)

		// use select so a pending send doesn't block shutdown
		select {
		case jobs <- rowOrdered:
		case <-ctx.Done():
			wg.Done()
			log.Println("=> stopping csv dispatch due to shutdown")
			break
		}
	}
	close(jobs)
}

func doTheJob(ctx context.Context, workerIndex, counter int, db *sql.DB, values []interface{}) {
	for {
		if ctx.Err() != nil {
			return
		}

		var outerError error
		func(outerError *error) {
			defer func() {
				if err := recover(); err != nil {
					*outerError = fmt.Errorf("%v", err)
				}
			}()

			conn, err := db.Conn(ctx)
			if err != nil {
				if ctx.Err() != nil {
					return // context cancelled, exit quietly
				}
				log.Fatal(err.Error())
			}
			defer conn.Close()

			query := fmt.Sprintf("INSERT INTO domain (%s) VALUES (%s)",
				strings.Join(dataHeaders, ","),
				strings.Join(generatePlaceholders(len(dataHeaders)), ","),
			)

			_, err = conn.ExecContext(ctx, query, values...)
			if err != nil {
				if ctx.Err() != nil {
					return // context cancelled mid-insert, exit quietly
				}
				log.Fatal(err.Error())
			}
		}(&outerError)
		if outerError == nil {
			break
		}
	}

	if counter%100 == 0 {
		log.Println("=> worker", workerIndex, "inserted", counter, "data")
	}
}

func generatePlaceholders(n int) []string {
	s := make([]string, 0)
	for i := 1; i <= n; i++ {
		s = append(s, fmt.Sprintf("$%d", i))
	}
	return s
}