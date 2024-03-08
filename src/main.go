package main

import (
	// "encoding/json"
	// "fmt"
	"fmt"
	"log"
	"os"
)

func main() {
	data ,err := os.ReadFile("../data.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data)) 
}