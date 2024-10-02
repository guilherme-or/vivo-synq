package main

import (
	"fmt"
	"time"
)

func main() {
	startDate := int64(956079801000000)
	endDate := int64(1059186314000000)
	loc, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		panic(err)
	}

	startDateTime := time.Unix(startDate/1000000, (startDate%1000000)*1000000).In(loc)
	endDateTime := time.Unix(endDate/1000000, (endDate%1000000)*1000000).In(loc)

	fmt.Println(startDateTime.Format(time.RFC3339)) // Output: 2018-01-16 23:07:10 -02:00
	fmt.Println(endDateTime.Format(time.RFC3339))   // Output: 2006-07-10 19:31:21 -03:00
}
