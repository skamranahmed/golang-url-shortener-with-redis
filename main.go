package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// load .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error getting env variables:", err)
		log.Fatalf("Exiting the program")
	}
}
