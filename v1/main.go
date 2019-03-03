package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

const filename string = `test.csv`

func main() {

	csvFile, _ := os.Open(filename)
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("Error reading CSV file %s: %s", filename, err)
	}

	var sum int64
	var nextAge int64
	for _, record := range records {
		nextAge, err = strconv.ParseInt(record[2], 10, 32)
		if err != nil {
			nextAge = 0
		}
		sum = sum + nextAge
	}

	fmt.Printf("Total age: %d\n", sum)

}
