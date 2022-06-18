package shows

import "github.com/graph-gophers/graphql-go"

type Struct struct {
	Id   string
	Name string
}

type Resolver struct {
	Show *Struct
}

func (r *Resolver) ID() graphql.ID {
	return graphql.ID(r.Show.Id)
}

func (r *Resolver) Name() string {
	return r.Show.Name
}
