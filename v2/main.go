package main

import (
	"fmt"
	"strconv"

	"github.com/mcgraw-bb25/csvcake/v2/mapreader"
)

const filename string = `test.csv`

func main() {

	records, err := mapreader.MapReader(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		var sum int64
		var nextAge int64
		for _, record := range records {
			nextAge, err = strconv.ParseInt(record["age"], 10, 32)
			if err != nil {
				nextAge = 0
			}
			sum = sum + nextAge
		}
		fmt.Printf("Total age: %d\n", sum)
	}

}
