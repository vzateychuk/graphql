package main

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"log"
	"net/http"
)

type Thing map[string]interface{}

func main() {

	mySchema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "RootQuery",
			Fields: graphql.Fields{
				"typed": &graphql.Field{
					Type: graphql.NewList(graphql.NewObject(graphql.ObjectConfig{
						Name: "Typed",
						Fields: graphql.Fields{
							"thing1": &graphql.Field{Type: graphql.String},
						},
					})),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return []Thing{
							{"thing1": "String Thing"},
							{"thing2": true},
							{"thing3": 123456},
						}, nil
					},
				},
				"untyped": &graphql.Field{
					Type: graphql.NewList(graphql.NewObject(graphql.ObjectConfig{
						Name: "Untyped",
						Fields: graphql.Fields{
							"name": &graphql.Field{Type: graphql.String},
						},
					})),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return []map[string]interface{}{
							{"name": "Swamp Thing"},
						}, nil
					},
				},
			},
		}),
	})

	if err != nil {
		log.Fatalln("Init mySchema", err)
	}

	http.HandleFunc("/meta", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("query"), mySchema)
		json.NewEncoder(w).Encode(result)
	})

	fmt.Println("Now server is running on port 8080")
	fmt.Println("Test with Get: curl -g 'http://localhost:8080/meta?query={Metadata(id:\"1\"){id, name, type}}'")
	http.ListenAndServe(":8080", nil)

}

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
