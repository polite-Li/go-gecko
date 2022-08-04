package main

import (
	"fmt"
	"log"

	gecko "github.com/polite-Li/go-gecko/v3"
)

func main() {
	cg := gecko.NewClient(nil)
	ping, err := cg.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ping.GeckoSays)
}
