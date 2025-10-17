package controllers

import (
	"go-gin-todo-graphql/graphql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GraphQLRequest struct {
	Query string `json:"query"`
}

func GraphQLHandler(c *gin.Context) {
	var req GraphQLRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := graphql.ExecuteQuery(req.Query)
	if len(result.Errors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"errors": result.Errors})
		return
	}

	c.JSON(http.StatusOK, result)
}
