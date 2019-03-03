package main

import (
	"fmt"

	"github.com/mcgraw-bb25/csvcake/v4/csvscanner"
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
func NewHouseHoldOccupantFromInterface(src []interface{}) []HouseHoldOccupant {
	hhos := make([]HouseHoldOccupant, 0)
	for _, hho := range src {
		newHHO, ok := hho.(HouseHoldOccupant)
		if ok {
			hhos = append(hhos, newHHO)
		}
	}
	return hhos
}

func main() {

	modelOccupant := HouseHoldOccupant{}
	rows, err := csvscanner.ScanCSV(filename, modelOccupant)
	if err != nil {
		fmt.Println(err)
	}
	records := NewHouseHoldOccupantFromInterface(rows)

	var sum int64
	for _, record := range records {
		sum += record.Age
	}
	fmt.Printf("Total age: %d\n", sum)
}
