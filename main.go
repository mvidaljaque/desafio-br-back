package main

import (
	. "desafio-br-back/routes"
	//"log"
	"net/http"
)

func main() {
	router := NewRouter()

	// api := router.Group("/api"){
	// 	api.router := NewRouter()
	// }

	http.ListenAndServe("localhost:5000", router)
	//http.ListenAndServe("0.0.0.0:5000", router)

	//router.run("0.0.0.0:5000", router)
	//log.Fatal(server)
}
