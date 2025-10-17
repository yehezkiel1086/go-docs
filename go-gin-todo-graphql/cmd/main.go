package main

import (
	controllers "go-gin-todo-graphql/controller"
	"go-gin-todo-graphql/graphql"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	graphql.InitSchema()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "GraphQL Todo API running on /graphql")
	})

	r.POST("/graphql", controllers.GraphQLHandler)

	r.Run(":3500")
}
