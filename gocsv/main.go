package main

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

const filename string = `test2.csv`

// HouseHoldOccupant has a name and age
type HouseHoldOccupant struct {
	Name string `csv:"name"`
	Age  int64  `csv:"age"`
}

func main() {

	hhoFile, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer hhoFile.Close()

	records := []*HouseHoldOccupant{}

	if err := gocsv.UnmarshalFile(hhoFile, &records); err != nil {
		panic(err)
	}

	var sum int64
	for _, record := range records {
		// fmt.Printf("%+v\n", record)
		sum += record.Age
	}

	fmt.Printf("Total age: %d\n", sum)

}
