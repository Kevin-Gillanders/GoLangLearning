package main

import (
	"fmt"
	"net/http"
	// "github.com/gorilla/mux"
)


func BookHandler(res http.ResponseWriter, req *http.Request){
	fmt.Println("Entered Book end point")
	fmt.Printf("The type of request is %v\n", req.Method)

	// res.WriteHeader(200)
	res.Header().Set("Content-Type", "application/json")

	data := []byte(`{"content" : "Book world"}`)
	res.Write(data)

}



func main (){
	fmt.Println("Test")

	mux := http.NewServeMux()


	mux.HandleFunc("/v1/Book", BookHandler)


	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	server.ListenAndServe()
}