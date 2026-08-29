package controller

import (
	"go-rabbitmq-work-queues/messaging/rabbitmq"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	mq rabbitmq.Rabbitmq
}

func InitController(mq rabbitmq.Rabbitmq) *Controller {
	return &Controller{
		mq: mq,
	}
}

type MessageReq struct {
	Message string `json:"message" binding:"required"`
}

func (ctl *Controller) PublishMessage(c *gin.Context) {
	// bind user input
	var req MessageReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "message is required",
		})
		return
	}

	// declare queue
	q, err := ctl.mq.DeclareQueue("message")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// publish message
	if err := ctl.mq.PublishMessage(q, req.Message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, "message sent successfully")
}
