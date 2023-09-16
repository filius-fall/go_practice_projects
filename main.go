package main

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
)


func main(){

	godotenv.Load(".env")

	port := os.Getenv("PORT")
	fmt.Printf("Hello world %v",port)
}
