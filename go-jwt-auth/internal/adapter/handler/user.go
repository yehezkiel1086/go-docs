package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yehezkiel1086/go-jwt-auth/internal/core/domain"
	"github.com/yehezkiel1086/go-jwt-auth/internal/core/port"
)

type UserHandler struct {
	svc port.UserService
}

func NewUserHandler(svc port.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

type createUserReq struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type updateUserReq struct {
	Email    string `json:"email" binding:"omitempty,email"`
	Password string `json:"password" binding:"omitempty,min=8"`
}

type userRes struct {
	ID    uint   `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Role domain.Role `json:"role"`
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req createUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := h.svc.CreateUser(c.Request.Context(), user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.svc.GetUsers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := make([]userRes, len(users))
	for i, u := range users {
		res[i] = toUserRes(&u)
	}

	c.JSON(http.StatusOK, res)
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}

	user, err := h.svc.GetUserByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, toUserRes(user))
}

func (h *UserHandler) GetUserByEmail(c *gin.Context) {
	email := c.Param("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email parameter is required"})
		return
	}

	user, err := h.svc.GetUserByEmail(c.Request.Context(), email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, toUserRes(user))
}

func (h *UserHandler) UpdateUserByID(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}

	var req updateUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input := &domain.User{
		Email:    req.Email,
		Password: req.Password,
	}

	if err := h.svc.UpdateUserByID(c.Request.Context(), id, input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user updated successfully"})
}

func (h *UserHandler) DeleteUserByID(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}

	if err := h.svc.DeleteUserByID(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}

func toUserRes(user *domain.User) userRes {
	return userRes{
		ID:    user.ID,
		Name: user.Name,
		Email: user.Email,
		Role: user.Role,
	}
}

func parseID(c *gin.Context) (uint, bool) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return 0, false
	}
	return uint(id), true
}

