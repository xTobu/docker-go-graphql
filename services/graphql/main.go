package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

type Player struct {
	Name   string
	Number string
}

var playerType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Player",
	Fields: graphql.Fields{
		"Name": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				m := p.Source.(Player)
				return m.Name, nil
			},
		},
		"Number": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				m := p.Source.(Player)
				return m.Number, nil
			},
		},
	},
})

// Schema
var fields = graphql.Fields{
	"hello": &graphql.Field{
		Type: graphql.NewList(playerType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// client, _ := airtable.New("key9dbABSvpCHm5KX", "appU6lGjdhBDSaosz")

			// type task struct {
			// 	AirtableID string
			// 	Fields     struct {
			// 		Name   string
			// 		Number string
			// 	}
			// }

			// tasks := []task{}
			// if err := client.ListRecords("Player", &tasks); err != nil {
			// 	panic(err)
			// }

			// fmt.Println(tasks)
			// return tasks, nil

			// ... some more code ...
			// what I return here was placed to demonstrate the error.
			return []Player{
				Player{
					Name:   "new-post",
					Number: "new post",
				}, Player{
					Name:   "new-post",
					Number: "new post",
				},
			}, nil
			// return "hello", nil
		},
	},
}
var rootQuery = graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
var schemaConfig = graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
var schema, _ = graphql.NewSchema(schemaConfig)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

// Handler initializes the graphql middleware.
func Handler() gin.HandlerFunc {
	// Creates a GraphQL-go HTTP handler with the defined schema
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   false,
		GraphiQL: false,
	})

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {

	router := gin.Default()

	api := router.Group("/graphql")
	api.Use()
	{
		api.POST("/test", Handler())
	}

	fmt.Println("Now server is running on port 8889")
	fmt.Println("Test with POST      : curl -g --request POST 'http://localhost:8889/graphql?query={hello}'")

	router.Run(":8889")

}
