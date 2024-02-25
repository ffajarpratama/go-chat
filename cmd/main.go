package main

import (
	"log"

	"github.com/ffajarpratama/go-chat/cmd/app"
)

func main() {
	// r := chi.NewRouter()
	// handler := handler.NewHandler(r)

	// fmt.Println("server listening on port :3000")
	// err := http.ListenAndServe(":3000", handler)
	// if err != nil {
	// 	log.Fatal("server error: ", err)
	// }

	err := app.StartServer()
	if err != nil {
		log.Fatal("start server error: ", err)
	}
}
