package graphql

import (
	models "go-gin-todo-graphql/model"

	"github.com/graphql-go/graphql"
)

var TodoType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Todo",
		Fields: graphql.Fields{
			"id":    &graphql.Field{Type: graphql.Int},
			"title": &graphql.Field{Type: graphql.String},
			"done":  &graphql.Field{Type: graphql.Boolean},
		},
	},
)

var QueryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"todos": &graphql.Field{
				Type: graphql.NewList(TodoType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return models.GetAllTodos(), nil
				},
			},
			"todo": &graphql.Field{
				Type: TodoType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{Type: graphql.Int},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(int)
					if ok {
						return models.GetTodoByID(id), nil
					}
					return nil, nil
				},
			},
		},
	},
)

var MutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createTodo": &graphql.Field{
				Type: TodoType,
				Args: graphql.FieldConfigArgument{
					"title": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					title := p.Args["title"].(string)
					return models.CreateTodo(title), nil
				},
			},
			"toggleTodo": &graphql.Field{
				Type: TodoType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := p.Args["id"].(int)
					return models.ToggleTodoDone(id), nil
				},
			},
		},
	},
)
