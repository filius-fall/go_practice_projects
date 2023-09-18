package main

import (
	"net/http"
	"log"
	"encoding/json"
)

func respWithError (w http.ResponseWriter, resp int, msg string){
	if resp > 499 {
		log.Printf("Error response %v",resp)
	}

	type errorResponse struct{
		Int64String string `json:"err"`
	}
	log.Printf("Inside Response With Error %v",msg)
	responWithJson(w,resp,errorResponse{
		Int64String: msg,
	})
}


func responWithJson(w http.ResponseWriter, resp int, payload interface{}){
	dat,err := json.Marshal(payload)
	if err != nil{
		log.Printf("Failed to Marshall the response %v",payload)
		w.WriteHeader(500)
		return
	}
	log.Printf("/n Response with json %v",dat)
	w.Header().Add("Content-Type","application/json")
	w.WriteHeader(resp)
	w.Write(dat)


}