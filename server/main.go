package main

import (
	"fmt"
	"net/http"
	"log"
	"time"

	"./data"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fromDate := time.Date(2019, time.November, 4, 0, 0, 0, 0, time.UTC)
	results, err := data.GetReadings(fromDate)
	if err != nil {
		log.Fatal("Could not fetch readings")
	}
	for _, data := range results {
		fmt.Fprintf(w, "%s %f\n", data.Timestamp, data.Temperature)
	}
}

func main() {
	err := data.InitDBConnection()
	if err != nil {
		log.Fatal("Couldn't initialize DB connection")
		return
	}
	
	http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
