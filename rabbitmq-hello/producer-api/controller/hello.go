package controller

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
)

type HelloController struct {
	ch *amqp.Channel
	q *amqp.Queue
}

func InitHelloController(ch *amqp.Channel, q *amqp.Queue) *HelloController {
	return &HelloController{
		ch: ch,
		q: q,
	}
}

type HelloReq struct {
	Message string `json:"message" binding:"required"`
}

func (h *HelloController) Hello(c *gin.Context) {
	// bind input
	var in HelloReq
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Errorf("failed to bind input: %v", err.Error()),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := h.ch.PublishWithContext(ctx,
		"",     // exchange
		h.q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
			ContentType: "text/plain",
			Body:        []byte(in.Message),
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Errorf("failed to publish message: %v", err.Error()),
		})
		return	
	}

	c.JSON(http.StatusOK, gin.H{
		"message": in.Message,
	})
}
