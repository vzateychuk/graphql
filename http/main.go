package main

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"http/meta"
	"log"
	"net/http"
)

func main() {
	err := meta.InitMetaStore("./meta/data.json")
	if err != nil {
		log.Fatalln("InitMetaStore", err)
	}

	http.HandleFunc("/meta", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("query"), schema)
		json.NewEncoder(w).Encode(result)
	})

	fmt.Println("Now server is running on port 8080")
	fmt.Println("Test with Get: curl -g 'http://localhost:8080/meta?query={Metadata(id:\"1\"){id, name, type}}'")
	http.ListenAndServe(":8080", nil)

}

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: meta.QueryType,
	},
)

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
