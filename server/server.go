package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/rs/cors"
	"trackingApp/graph/resolvers"
	"trackingApp/middleware"
	"trackingApp/services"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"trackingApp/database"
	"trackingApp/graph/generated"
)

const defaultPort = "8080"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	Database := database.Connect()
	usersRepo := database.UsersRepo{DB: Database}
	itemsRepo := database.ItemsRepo{DB: Database}

	s := services.NewService(usersRepo, itemsRepo)
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	router := chi.NewRouter()
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080", "http://localhost:3000"},
		AllowedHeaders:   []string{"Authorization", "Content-Type", "Token"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{Services: s}}))

	wrapped := middleware.AuthMiddleware(srv)

	//wrapped := middleware.CorsMiddleware(srv)
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", middleware.DataloaderMiddleware(Database, wrapped))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
