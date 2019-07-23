package main

import (
	"fmt"
	"strconv"

	"github.com/mcgraw-bb25/csvcake/v3/filteredmapreader"
)

const filename string = `test.csv`

func main() {

	columnList := []string{"name", "age"}

	records, err := filteredmapreader.FilteredMapReader(filename, columnList...)
	if err != nil {
		fmt.Println(err)
	}

	var sum int64
	for _, record := range records {
		nextAge, err := strconv.ParseInt(record["age"], 10, 32)
		if err != nil {
			nextAge = 0
		}
		sum = sum + nextAge
	}
	fmt.Printf("Total age: %d\n", sum)

}
