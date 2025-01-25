package graph

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/emoss08/trenova/internal/api/graph/generated"
	"github.com/emoss08/trenova/internal/api/graph/resolvers"
	"github.com/vektah/gqlparser/v2/ast"

	"go.uber.org/fx"
)

const defaultPort = "8080"

type ServerParams struct {
	fx.In

	LC fx.Lifecycle

	// Resolvers
	Resolver *resolvers.Resolver
}

type Server struct {
	srv *handler.Server
	lc  fx.Lifecycle
}

func NewServer(p ServerParams) *Server {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.New(generated.NewExecutableSchema(generated.Config{Resolvers: p.Resolver}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	server := &Server{srv: srv, lc: p.LC}

	p.LC.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return server.Start()
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})

	return server
}

func (s *Server) Start() error {
	go func() {
		http.Handle("/", playground.Handler("GraphQL playground", "/query"))
		http.Handle("/query", s.srv)

		log.Printf("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
		log.Fatal(http.ListenAndServe(":"+defaultPort, nil))
	}()

	return nil
}
