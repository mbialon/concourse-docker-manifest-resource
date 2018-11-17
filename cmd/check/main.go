package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {
	err := json.NewEncoder(os.Stdout).Encode([]interface{}{})
	if err != nil {
		log.Fatal(err)
	}
}
