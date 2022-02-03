package meta

import (
	"github.com/graphql-go/graphql"
)

/*
   Create metadataType object type with fields "id", "name", "type" by using GraphQLObjectTypeConfig:
       - Name: name of object type ("Metadata")
       - Fields: a map of fields by using GraphQLFields
   Setup type of field use GraphQLFieldConfig
*/
var metadataType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Metadata",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"type": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
