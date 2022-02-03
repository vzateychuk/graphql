package meta

import "github.com/graphql-go/graphql"

/*
   Create Query object type with fields "meta" has type [metaType] by using GraphQLObjectTypeConfig:
       - Name: name of object type ("Metadata")
       - Fields: a map of fields by using GraphQLFields
   Setup type of field use GraphQLFieldConfig to define:
       - Type: type of field
       - Args: arguments to query with current field
       - Resolve: function to query data using params from [Args] and return value with current type
*/
var QueryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"Metadata": &graphql.Field{
				Type: metadataType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, isOK := p.Args["id"].(string)
					if isOK {
						return store[idQuery], nil
					}
					return nil, nil
				},
			},
		},
	})
