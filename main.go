package main

import (
	"log"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	gqlmerge "github.com/mattdamon108/gqlmerge/lib"
)

type showStruct struct {
	id   string
	name string
}

type showResolver struct {
	show *showStruct
}

func (r *showResolver) ID() graphql.ID {
	return graphql.ID(r.show.id)
}

func (r *showResolver) Name() string {
	return r.show.name
}

type query struct{}

func (q *query) Show(args struct{ Id string }) *showResolver {
	return &showResolver{&showStruct{id: args.Id, name: "pas ok"}}
}

func main() {
	s := gqlmerge.Merge(" ", "./features/shows/schema.graphql")
	schemaString := *s
	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	schema := graphql.MustParseSchema(schemaString, &query{}, opts...)
	http.Handle("/", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":3000", nil))
}
