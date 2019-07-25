package main

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/mcgraw-bb25/csvcake/v3/filteredmapreader"
)

const filename string = `test.csv`

// HouseHoldOccupant has a name and age
type HouseHoldOccupant struct {
	Name string `csvcake:"name"`
	Age  int64  `csvcake:"age"`
}

func (h HouseHoldOccupant) SayHello() {
	fmt.Printf("Hello!  My name is %s!  See you soon!\n", h.Name)
}

func main() {

	// records, err := mapreader.MapReader(filename)
	// if err != nil {
	// fmt.Println(err)
	// }
	// fmt.Printf("Records: %+v\n", records)

	var modelOccupant HouseHoldOccupant

	// var myInt int64
	// var myString string

	// myInt = 50
	// myString = "Matt"

	aValue := reflect.New(reflect.TypeOf(modelOccupant))
	// fmt.Printf("aValue: %+v\n", aValue)
	aConcrete := reflect.Indirect(aValue)
	// fmt.Printf("aConcrete: %+v\n", aConcrete)
	aType := reflect.TypeOf(modelOccupant)
	// fmt.Printf("aType: %+v\n", aType)
	// aName := reflect.ValueOf(modelOccupant.Name)
	// aInt := reflect.ValueOf(modelOccupant.Age)
	// aValue.Name = myString
	// aValue.Age = myInt
	// fmt.Printf("%+v", aValue)

	// fmt.Printf("aConcrete NumField: %d\n", aConcrete.NumField())
	// fmt.Printf("aType NumField: %d\n", aType.NumField())

	columnList := make([]string, 0)
	idx := 0
	for idx < aType.NumField() {
		field := aType.Field(idx)
		csvColumn := field.Tag.Get("csvcake")
		columnList = append(columnList, csvColumn)
		idx++
	}

	filteredRecords, err := filteredmapreader.FilteredMapReader(filename, columnList...)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Printf("FilteredRecords: %+v\n", filteredRecords)

	typedRecords := make([]interface{}, 0)

	for _, filteredRecord := range filteredRecords {
		// fmt.Printf("filteredRecord: %+v\n", filteredRecord)

		idx = 0
		for idx < aType.NumField() {
			field := aType.Field(idx)
			csvColumn := field.Tag.Get("csvcake")

			// fmt.Printf("A1-%+v\n", field)
			// fmt.Printf("A2-%+v\n", csvColumn)
			// fmt.Printf("A3-%+v\n", field.Index[0])
			// fmt.Printf("A4-%+v\n", field.Type)
			// fmt.Printf("B-%+v\n", aValue)
			// fmt.Printf("C-%+v\n", reflect.ValueOf(aValue))
			// fmt.Printf("D-%+v\n", reflect.Indirect(aValue))

			thisField := aConcrete.Field(idx)
			// fmt.Printf("E-%+v\n", thisField)
			fieldType := fmt.Sprintf("%s", field.Type)
			switch fieldType {
			case "string":
				thisField.SetString(filteredRecord[csvColumn])
			case "int64":
				var i64AsInt int64
				i64AsInt, _ = strconv.ParseInt(filteredRecord[csvColumn], 10, 64)
				thisField.SetInt(i64AsInt)
			}
			idx++
			// fmt.Printf("F-%+v\n", aConcrete)
			// fmt.Printf("G-%s\n", reflect.TypeOf(aConcrete))
			// fmt.Printf("H-%+v\n", aConcrete.Interface())
			// fmt.Printf("I-%+v\n", aConcrete.Interface().(HouseHoldOccupant))

		}
		typedRecords = append(typedRecords, aConcrete.Interface())
	}
	// fmt.Printf("J-%+v\n", typedRecords)
	for _, typedRec := range typedRecords {
		hho, ok := typedRec.(HouseHoldOccupant)
		if !ok {
			panic("Help!")
		}
		// fmt.Printf("HHO: %v\n", hho)
		fmt.Printf("HHO+: %+v\n", hho)
		hho.SayHello()
	}
}
