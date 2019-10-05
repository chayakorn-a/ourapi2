package main

import (
	"log"

	"github.com/chayakorn-a/ourapi2/router"
)

func main() {
	e := router.NewRouter()
	log.Fatal(e.Start(":1323"))
}
