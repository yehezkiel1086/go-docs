package handler

import (
	"context"
	"log"

	"github.com/yehezkiel1086/go-rabbitmq-email-notification/notif-service/internal/core/port"
)

type NotifHandler struct {
	svc port.NotifService
}

func NewNotifHandler(svc port.NotifService) *NotifHandler {
	return &NotifHandler{svc: svc}
}

func (h *NotifHandler) ReceiveNotif(ctx context.Context) {
	msgs, err := h.svc.ReceiveNotif(ctx)
	if err != nil {
		log.Printf("notif: failed to start consumer: %v", err)
		return
	}

	log.Printf("notif: waiting for messages. To exit press CTRL+C")

	for {
		select {
		case <-ctx.Done():
			// context cancellation for clean shutdown.
			log.Printf("notif: context cancelled, stopping consumer")
			return

		case msg, ok := <-msgs:
			if !ok {
				// channel closed — broker disconnected.
				log.Printf("notif: message channel closed, stopping consumer")
				return
			}

			if err := h.svc.SendConfirmationEmail(ctx, msg.Body); err != nil {
				// NACK with requeue=true so the message is retried.
				log.Printf("notif: failed to send email, requeuing message: %v", err)
				if err := msg.Nack(false, true); err != nil {
					log.Printf("notif: failed to nack message: %v", err)
				}
				continue
			}

			if err := msg.Ack(false); err != nil {
				log.Printf("notif: failed to ack message: %v", err)
			}
		}
	}
}
