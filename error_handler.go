package main

import(
	"net/http"
	"fmt"
)


func errorHandler(w http.ResponseWriter, r *http.Request){
	fmt.Printf("Inside Error Handler")
	respWithError(w,500,"Something went very wrong")
}