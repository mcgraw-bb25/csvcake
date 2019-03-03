package main

import (
	"fmt"

	"github.com/mcgraw-bb25/csvcake/v5/iterscanner"
)

const filename string = `test.csv`

// HouseHoldOccupant has a name and age
type HouseHoldOccupant struct {
	Name string `csvcake:"name"`
	Age  int64  `csvcake:"age"`
}

// Bake implements the Bakeable interface
func (h HouseHoldOccupant) Bake(src map[string]interface{}) interface{} {
	var newOccupant HouseHoldOccupant
	newOccupant.Name = src["name"].(string)
	newOccupant.Age = src["age"].(int64)
	return newOccupant
}

// NewHouseHoldOccupantFromInterface creates a slice of HouseHoldOccupants
// from a list of interfaces produced by calls to Bake()
func NewHouseHoldOccupantFromInterface(src interface{}) HouseHoldOccupant {
	newHHO, _ := src.(HouseHoldOccupant)
	return newHHO
}

func main() {

	modelOccupant := HouseHoldOccupant{}
	iCSVScanner, err := iterscanner.NewIterScanner(filename, modelOccupant)
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
		record := NewHouseHoldOccupantFromInterface(row)
		sum += record.Age
	}

	fmt.Printf("Total age: %d\n", sum)
}
