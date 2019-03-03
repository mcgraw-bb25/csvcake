package csvscanner

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/mcgraw-bb25/csvcake/v3/filteredmapreader"
)

const tagID string = `csvcake`

type Bakeable interface {
	Bake(map[string]interface{}) interface{}
}

type Preparable interface {
	Prepare(string) (interface{}, error)
}

type PrepareString struct{}

func (p PrepareString) Prepare(stringAsString string) (interface{}, error) {
	return stringAsString, nil
}

type PrepareInt64 struct{}

func (p PrepareInt64) Prepare(i64AsString string) (interface{}, error) {
	var i64AsInt int64
	i64AsInt, err := strconv.ParseInt(i64AsString, 10, 64)
	if err != nil {
		return i64AsInt, err
	}
	return i64AsInt, nil
}

// getPreparer is an internal method that switches to pass
// the appropriate Preparer back to the caller.
func getPreparer(preparerType string, dataAsString string) interface{} {
	var preparer Preparable
	switch preparerType {
	case "string":
		preparer = PrepareString{}
	case "int64":
		preparer = PrepareInt64{}
	}
	value, err := preparer.Prepare(dataAsString)
	if err != nil {
		return nil
	}
	return value
}

// ScanCSV uses the tags from the Bakeable interface and
// determines what columns to parse and return to the caller.
// The caller can then use the appropriate interface factory
// to get back their expected struct.
// This function drew inspiration from the following blog =>
// https://sosedoff.com/2016/07/16/golang-struct-tags.html
func ScanCSV(filename string, modelFactory Bakeable) ([]interface{}, error) {
	bakeable := reflect.TypeOf(modelFactory)

	returnVals := make([]interface{}, 0)
	columnList := make([]string, 0)
	preparers := make(map[string]string)

	idx := 0
	for idx < bakeable.NumField() {
		field := bakeable.Field(idx)
		csvColumn := field.Tag.Get(tagID)
		columnList = append(columnList, csvColumn)
		preparers[csvColumn] = field.Type.Name()
		idx++
	}

	records, err := filteredmapreader.FilteredMapReader(filename, columnList...)
	if err != nil {
		fmt.Println(err)
	}

	for _, record := range records {
		newVal := make(map[string]interface{})
		for _, column := range columnList {
			newVal[column] = getPreparer(preparers[column], record[column])
		}
		newBaked := modelFactory.Bake(newVal)
		returnVals = append(returnVals, newBaked)
	}

	return returnVals, nil
}
