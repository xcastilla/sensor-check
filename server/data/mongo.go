package main

import (
	"fmt"
	"os"

	"gopkg.in/mgo.v2"
)

func main() {
	fmt.Println("pong")
	url := os.Getenv("MONGO_URL")
	fmt.Printf("[%v]\n", url)
	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}

	fmt.Println("ping", session.Ping())

}
