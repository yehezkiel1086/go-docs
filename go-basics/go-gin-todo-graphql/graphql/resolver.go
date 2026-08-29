package graphql

import "github.com/graphql-go/graphql"

var Schema graphql.Schema

func InitSchema() {
	var err error
	Schema, err = graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    QueryType,
			Mutation: MutationType,
		},
	)
	if err != nil {
		panic(err)
	}
}

func ExecuteQuery(query string) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        Schema,
		RequestString: query,
	})
	return result
}
