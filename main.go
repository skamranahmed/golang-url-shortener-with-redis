package main

import (
	"flag"
	"log"

	"github.com/golang/glog"
	"github.com/joho/godotenv"
	"github.com/skamranahmed/golang-url-shortner-with-redis/models"
)

func main() {
	// set flag parse before glog to avoid clumpsy glog statement
	flag.Parse()
	flag.Lookup("alsologtostderr").Value.Set("true")

	// load .env file
	err := godotenv.Load()
	if err != nil {
		glog.Error("Error getting env variables: ", err.Error())
		log.Fatal("Exiting the program")
	}

	// connect to db
	glog.Info("Initializing the db...")
	db, err := models.InitRedisDB()
	if err != nil {
		glog.Error("Cannot initialize the db: ", err.Error())
		log.Fatal("Exiting the program")
	}
	glog.Info("Successfully established connection with the db!")

	// Defer this so that if our application exits, we close the db.
	defer db.Close()

}
