package main

import (
	"fmt"
	"./models"
	"./data"
	"time"
)

func main() {
	db, err := data.GetDBConnection()
	fmt.Println(db, err)
	fromDate := time.Date(2014, time.November, 4, 0, 0, 0, 0, time.UTC)
	results, err := models.GetReadings(db, fromDate)
	for _, d := range results {
		fmt.Println(d)
	}

}
