package main

import (
	"fmt"

	"github.com/mcgraw-bb25/csvcake/v6/iterstructscanner"
)

const filename string = `test.csv`

// HouseHoldOccupant has a name and age
type HouseHoldOccupant struct {
	Name string `csvcake:"name"`
	Age  int64  `csvcake:"age"`
}

func main() {

	modelOccupant := HouseHoldOccupant{}
	iCSVScanner, err := iterstructscanner.NewIterStructScanner(filename, modelOccupant)
	if err != nil {
		panic(err)
	}
	defer iCSVScanner.Close()

	var sum int64
	for {
		row, err := iCSVScanner.Next()
		if err != nil {
			break
		}
		record := row.(HouseHoldOccupant)
		sum += record.Age
	}

	fmt.Printf("Total age: %d\n", sum)
}
