package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/yehezkiel1086/go-workerpool-file-processing/config"
	"github.com/yehezkiel1086/go-workerpool-file-processing/db"
	"github.com/yehezkiel1086/go-workerpool-file-processing/service"
	"github.com/yehezkiel1086/go-workerpool-file-processing/util"
)

func main() {
	start := time.Now()

	cfg := config.LoadConfig()

	d, err := db.OpenDbConnection(cfg)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer d.Close()

	if err = db.AutoMigrate(d); err != nil {
		log.Fatal(err.Error())
	}

	csvReader, csvFile, err := util.OpenCsvFile(cfg)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer csvFile.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// listen for ctrl+c / SIGTERM
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-sigCh
		log.Printf("=> received signal: %s, shutting down gracefully...", sig)
		cancel()
	}()

	jobs := make(chan []interface{}, 0)
	wg := new(sync.WaitGroup)

	go service.DispatchWorkers(ctx, d, jobs, wg, cfg.TotalWorker)
	service.ReadCsvFilePerLineThenSendToWorker(ctx, csvReader, jobs, wg)

	wg.Wait()

	if ctx.Err() != nil {
		log.Println("=> import interrupted before completion")
	} else {
		log.Println("=> import completed successfully")
	}

	duration := time.Since(start)
	fmt.Println("done in", int(math.Ceil(duration.Seconds())), "seconds")
}