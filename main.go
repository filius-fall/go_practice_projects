package main

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"net/http"
	"log"
)


func main(){

	godotenv.Load(".env")

	port := os.Getenv("PORT")
	fmt.Printf("Hello world %v",port)


	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins : []string{"http://*","https://*"},
	}))

	v1router := chi.NewRouter()

	router.Mount("/v1",v1router)

	v1router.Get("/test",httpHandler)
	v1router.Get("/err",errorHandler)
	server := &http.Server{
		Handler: router,
		Addr : ":" + port,
	}

	log.Printf("Server starting at port %v",port)

	err := server.ListenAndServe()
	if err != nil{
		log.Fatal(err)
	}

}
