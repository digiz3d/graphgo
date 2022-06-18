package main

import (
	"log"
	"net/http"

	features "github.com/digiz3d/graphgo/features"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	gqlmerge "github.com/mattdamon108/gqlmerge/lib"
)

func main() {
	s := gqlmerge.Merge(" ", "./features/shows/schema.graphql")
	schemaString := *s
	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	schema := graphql.MustParseSchema(schemaString, &features.Query{}, opts...)
	http.Handle("/", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":3000", nil))
}
