// Package iterscanner implements an iteration based CSV scanner.
package iterscanner

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"reflect"
)

const tagID string = `csvcake`

// Iterscanner is the centralized struct to manage iteratively
// parsing a CSV.  It holds the reference Bakeable interface
// and the original filename upon initialization via NewIterScanner.
// Upon initialization this struct holds the pointers to the
// os and csv package types File and Reader respectively.
type IterScanner struct {
	Filename     string
	ModelFactory Bakeable

	preparers map[string]string
	headers   map[int]string

	csvFile   *os.File
	csvReader *csv.Reader
}

// NewIterScanner is an initialization function to return a properly
// initialized IterScanner.  It delegates into three unexported functions
// initializeCSV, initializePreparers, and initializeHeader to
// properly initialize the IterScanner.
func NewIterScanner(filename string, modelFactory Bakeable) (IterScanner, error) {
	n := IterScanner{
		Filename:     filename,
		ModelFactory: modelFactory}
	n.preparers = make(map[string]string)
	n.headers = make(map[int]string)

	err := n.initializeCSV()
	if err != nil {
		return n, fmt.Errorf("Could not initialize CSV: %s\nError: %s", filename, err)
	}
	err = n.initializePreparers()
	if err != nil {
		return n, fmt.Errorf("Could not initialize CSV Preparers: %s\nError: %s", filename, err)
	}
	err = n.initializeHeader()
	if err != nil {
		return n, fmt.Errorf("Could not initialize CSV Headers: %s\nError: %s", filename, err)
	}

	return n, nil
}

// initializeCSV is an unexported method that opens the CSV file
// and initializes the csv Reader, embedding both on i, an IterScanner.
func (i *IterScanner) initializeCSV() error {
	csvFile, err := os.Open(i.Filename)
	if err != nil {
		return fmt.Errorf("Could not open file: %s\nError: %s", i.Filename, err)
	}
	i.csvFile = csvFile
	i.csvReader = csv.NewReader(i.csvFile)
	return nil
}

// initializePreparers uses reflect to learn from i.ModelFactory
// what fields and preparers are required to scan for this csv.
// This function drew inspiration from the following blog =>
// https://sosedoff.com/2016/07/16/golang-struct-tags.html
func (i *IterScanner) initializePreparers() error {
	bakeable := reflect.TypeOf(i.ModelFactory)

	idx := 0
	for idx < bakeable.NumField() {
		field := bakeable.Field(idx)
		csvColumn := field.Tag.Get(tagID)
		idx++
		i.preparers[csvColumn] = field.Type.Name()
	}

	return nil
}

// nextRow is an unexported method that calls the embedded
// csvReader.Read() function and proxies the result back to
// the caller.  It is split so that headers and Next() need
// not reimplement the functionality.
func (i *IterScanner) nextRow() ([]string, error) {
	errStrings := make([]string, 0)
	record, err := i.csvReader.Read()
	if err == io.EOF {
		return errStrings, err
	}
	if err != nil {
		return errStrings, err
	}
	return record, nil
}

// initializeHeader is an unexported method that reads the
// first row of a CSV file and initializes the header i.headers
// map.  It is called after i.preparers are initialized.
func (i *IterScanner) initializeHeader() error {
	row, err := i.nextRow()
	if err != nil {
		return err
	}

	for index, column := range row {
		_, ok := i.preparers[column]
		if ok {
			i.headers[index] = column
		}
	}

	return nil
}

// Close is the method to actually close the embedded *os.File
// that IterScanner is reading from.  This should be closed by
// using defer myIterScanner.Close() in the function that initializes
// the IterScanner.
func (i *IterScanner) Close() error {
	err := i.csvFile.Close()
	if err != nil {
		return fmt.Errorf("Could not close file: %s\nError: %s", i.Filename, err)
	}
	return nil
}

// Next calls the next Read() on the csv file and returns
// and interface that conforms to the Bake() factory.
// Some inspiration from this blog post => https://ewencp.org/blog/golang-iterators/
func (i *IterScanner) Next() (interface{}, error) {
	nextRecord := make(map[string]string)

	row, err := i.nextRow()
	if err != nil {
		return nil, err
	}

	for index, column := range row {
		columnName := i.headers[index]
		_, ok := i.preparers[columnName]
		if ok {
			nextRecord[columnName] = column
		}
	}

	newVal := make(map[string]interface{})
	for keyColumn, _ := range i.preparers {
		newVal[keyColumn] = getPreparer(i.preparers[keyColumn], nextRecord[keyColumn])
	}
	newBaked := i.ModelFactory.Bake(newVal)
	return newBaked, err
}
