package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/notif-service/internal/adapter/config"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/notif-service/internal/adapter/storage/postgres"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/notif-service/internal/adapter/storage/postgres/repository"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/notif-service/internal/adapter/storage/rabbitmq"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/notif-service/internal/core/domain"
)

func handleError(msg string, err error) {
	if err != nil {
		slog.Error(msg, "error", err)
		os.Exit(1)
	}
}

func main() {
	// load .env configs
	conf, err := config.New()
	handleError("failed to load .env configs", err)
	slog.Info(".env configs loaded successfully", "app", conf.App.Name, "env", conf.App.Env)

	ctx := context.Background()

	// init db
	db, err := postgres.New(ctx, conf.DB)
	handleError("failed to init db", err)
	slog.Info("db initialized successfully", "db", conf.DB.Name)

	// migrate dbs
	err = db.Migrate(&domain.Notification{})
	handleError("failed to migrate dbs", err)
	slog.Info("dbs migrated successfully")

	// init rabbitmq
	mq, err := rabbitmq.New(conf.Rabbitmq)
	handleError("failed to connect rabbitmq", err)

	// dependency injection
	notifRepo := repository.NewNotifRepository(db)

	// consume rabbitmq
	msgs, err := mq.Consume()
	handleError("failed to init rabbitmq consumer", err)

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			notifRepo.CreateNotification(ctx, &domain.Notification{
				UserID:  1, // TODO: Extract from message
				Message: string(d.Body),
				Type:    "order_notification",
			})
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
