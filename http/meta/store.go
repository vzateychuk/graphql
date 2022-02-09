package meta

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"io/ioutil"
)

type Thing map[string]interface{}

// Shared data variables to allow dynamic reloads
var MetaSchema graphql.Schema

//Helper function to import json from file to map
func ImportJsonFileAndInitSchema(fileName string) error {

	// Load JSON from file and unmarshall to Data collection
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("ioutil.ReadFile: %w", err)
	}
	var data []map[string]interface{}
	err = json.Unmarshal(content, &data)
	if err != nil {
		return fmt.Errorf("json.Unmarshal: %w", err)

	}

	// getting fields and values of the Data collection
	fields := make(graphql.Fields)
	args := make(graphql.FieldConfigArgument)

	for _, item := range data {
		for k := range item {
			if k != "meta" {
				fields[k] = &graphql.Field{
					Type: graphql.String,
				}
				args[k] = &graphql.ArgumentConfig{
					Type: graphql.String,
				}
			} else {
				for m := range item[k].(map[string]interface{}) {
					fmt.Println(m)
					fields[m] = &graphql.Field{
						Type: graphql.String,
					}
					args[m] = &graphql.ArgumentConfig{
						Type: graphql.String,
					}
				}
			}
		}
	}

	// Declare metaType object type with dynamic list of fields
	var metaType = graphql.NewObject(
		graphql.ObjectConfig{
			Name:   "Metadata",
			Fields: fields,
		},
	)

	// Query object type with fields "meta" has type [metaType]
	var metaQueryType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{

				// curl -g http://localhost:8080/meta?query={one(id:"1"){id, name, type}}
				"one": &graphql.Field{
					Type:        metaType,
					Description: "FindOne meta by criteria",
					Args:        args,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return filterMeta(data, p.Args), nil
					},
				},

				// curl -g 'http://localhost:8080/meta?query={list{id,name,type}}'
				"list": &graphql.Field{
					Type:        graphql.NewList(metaType),
					Description: "List of metas",
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return data, nil
					},
				},
			},
		})

	// MetaSchema Definition
	MetaSchema, _ = graphql.NewSchema(
		graphql.SchemaConfig{
			Query: metaQueryType,
		},
	)

	return nil
}

// filter metadata by args
func filterMeta(data []map[string]interface{},
	args map[string]interface{},
) map[string]interface{} {

	for _, item := range data {
		for idx, arg := range args {
			if item[idx] != arg {
				goto nextitem
			}
			return item
		}
	nextitem:
	}
	return nil
}
