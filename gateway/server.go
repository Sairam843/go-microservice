package main

import (
	"errors"
	"gateway/graph/generated"
	"gateway/graph/resolvers"
	"log"
	"net/http"
	"os"

	userPB "gateway/graph/pb/user"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"github.com/vektah/gqlparser/v2/ast"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const defaultPort = "8080"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	userClient, err := setupGRPCClient(os.Getenv("USER_SERVICE_URL"), "USER")
	if err != nil {
		log.Fatalf("error setting up USER service GRPC connection: %v", err)
	}

	resolver := &resolvers.Resolver{
		UserClient: userClient.(userPB.UserClient),
	}

	srv := handler.New(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func setupGRPCClient(URL, serviceName string) (interface{}, error) {
	if URL == "" {
		return nil, errors.New("invalid URL")
	}

	switch serviceName {
	case "USER":
		conn, err := grpc.NewClient(URL, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return nil, err
		}
		log.Printf("gRPC connected to %s for service: %s", URL, serviceName)
		return userPB.NewUserClient(conn), nil

	default:
		return nil, errors.New("unknown service name")
	}

}
