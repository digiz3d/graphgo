package features

import "github.com/digiz3d/graphgo/features/shows"

type Query struct{}

func (q *Query) Show(args struct{ Id string }) *shows.Resolver {
	return &shows.Resolver{Show: &shows.Struct{Id: args.Id, Name: "pas ok"}}
}
