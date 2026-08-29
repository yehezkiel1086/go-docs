package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/order-service/internal/core/domain"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/order-service/internal/core/port"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/order-service/internal/core/util"
)

type OrderHandler struct {
	svc port.OrderService
}

func NewOrderHandler(svc port.OrderService) *OrderHandler {
	return &OrderHandler{
		svc: svc,
	}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var req domain.CreateOrderReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get userID from context (set by AuthMiddleware)
	user, exists := c.Get(string(ContextKeyUser))
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	claims, ok := user.(*util.JWTClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	order, err := h.svc.CreateOrder(c.Request.Context(), uint(claims.UserID), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, order)
}

func (h *OrderHandler) GetOrderByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order id"})
		return
	}

	order, err := h.svc.GetOrderByID(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) GetMyOrders(c *gin.Context) {
	// Get userID from context (set by AuthMiddleware)
	user, exists := c.Get(string(ContextKeyUser))
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	claims, ok := user.(*util.JWTClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	orders, err := h.svc.GetOrdersByUserID(c.Request.Context(), uint(claims.UserID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func (h *OrderHandler) UpdateOrder(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order id"})
		return
	}

	var req domain.UpdateOrderReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := h.svc.UpdateOrder(c.Request.Context(), uint(id), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) DeleteOrder(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order id"})
		return
	}

	if err := h.svc.DeleteOrder(c.Request.Context(), uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "order deleted successfully"})
}
