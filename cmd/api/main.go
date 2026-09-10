package main

import (
	"log"
	"github.com/dakotaodev/cradle/internal/api"
)

func main() {

	router := api.NewRouter()
	if err:= router.Run(); err != nil {
		log.Fatalf("unable to start the router: %v", err)
	}

}