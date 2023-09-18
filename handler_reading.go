package main

import (
	"fmt"
	"net/http"
)

func httpHandler(w http.ResponseWriter, r *http.Request){
	responWithJson(w,200,struct{}{})
	fmt.Printf("This is Respinse handler")
}