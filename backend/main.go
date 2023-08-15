package main

import (
	"fmt"
	"moflo-be/config"
	"moflo-be/routers"
	"net/http"
)

const portNumber int = 5001

func main() {
	router := routers.Router

	config.ConnectDatabase()

	go func() {
		if err := http.ListenAndServe(fmt.Sprintf(":%d", portNumber), router); err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Println("Server is running on port 8080")

	select {}
}
