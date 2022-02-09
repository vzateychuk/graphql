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
	err := meta.ImportJsonFileAndInitSchema("./meta/data.json")
	if err != nil {
		log.Fatalln("ImportJsonFileAndInitSchema", err)
	}

	http.HandleFunc("/meta", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("query"), meta.MetaSchema)
		json.NewEncoder(w).Encode(result)
	})

	fmt.Println("Now server is running on port 8080")
	fmt.Println("Test with Get: curl -g 'http://localhost:8080/meta?query={one(name:%22Dan%22){id,name,type}}'")
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
