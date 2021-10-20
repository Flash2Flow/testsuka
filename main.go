package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

)


func main() {
	server()
}

func server()  {
	//design pages
	http.HandleFunc("/", home)

	//tech pages
	http.HandleFunc("/reg", reg)
	http.HandleFunc("/auth", auth)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))


	//server methods
	port := os.Getenv("PORT")
	log.Println("start server with port: " +port)
	PortConnect := fmt.Sprintf(":%s", port)
	http.ListenAndServe(PortConnect, nil)
}