// Package iterstructscanner implements an iteration based CSV scanner.
package iterstructscanner

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
)

const tagID string = `csvcake`

// IterStructScanner needs a doc
type IterStructScanner struct {
	Filename     string
	ModelFactory interface{}

	headers        map[int]string
	reverseHeaders map[string]int

	csvFile   *os.File
	csvReader *csv.Reader

	aConcrete reflect.Value
	aType     reflect.Type
}

// NewIterScanner needs a doc string
func NewIterStructScanner(filename string, modelFactory interface{}) (IterStructScanner, error) {
	n := IterStructScanner{
		Filename:     filename,
		ModelFactory: modelFactory}
	n.headers = make(map[int]string)
	n.reverseHeaders = make(map[string]int)

	aValue := reflect.New(reflect.TypeOf(modelFactory))
	n.aConcrete = reflect.Indirect(aValue)
	n.aType = reflect.TypeOf(modelFactory)

	err := n.initializeCSV()
	if err != nil {
		return n, fmt.Errorf("Could not initialize CSV: %s\nError: %s", filename, err)
	}
	err = n.initializeHeader()
	if err != nil {
		return n, fmt.Errorf("Could not initialize CSV Headers: %s\nError: %s", filename, err)
	}

	return n, nil
}

// initializeCSV is an unexported method that opens the CSV file
// and initializes the csv Reader, embedding both on i, an IterScanner.
func (i *IterStructScanner) initializeCSV() error {
	csvFile, err := os.Open(i.Filename)
	if err != nil {
		return fmt.Errorf("Could not open file: %s\nError: %s", i.Filename, err)
	}
	i.csvFile = csvFile
	i.csvReader = csv.NewReader(i.csvFile)
	return nil
}

// initializeHeader is an unexported method that reads the
// first row of a CSV file and initializes the header i.headers
// map.  It is called after i.preparers are initialized.
func (i *IterStructScanner) initializeHeader() error {
	row, err := i.nextRow()
	if err != nil {
		return err
	}

	for index, column := range row {
		i.headers[index] = column
		i.reverseHeaders[column] = index
	}

	return nil
}

// nextRow is an unexported method that calls the embedded
// csvReader.Read() function and proxies the result back to
// the caller.  It is split so that headers and Next() need
// not reimplement the functionality.
func (i *IterStructScanner) nextRow() ([]string, error) {
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

// Close is the method to actually close the embedded *os.File
// that IterScanner is reading from.  This should be closed by
// using defer myIterScanner.Close() in the function that initializes
// the IterScanner.
func (i *IterStructScanner) Close() error {
	err := i.csvFile.Close()
	if err != nil {
		return fmt.Errorf("Could not close file: %s\nError: %s", i.Filename, err)
	}
	return nil
}

// Next calls the next Read() on the csv file and returns
// and interface.
// Some inspiration from this blog post => https://ewencp.org/blog/golang-iterators/
func (i *IterStructScanner) Next() (interface{}, error) {
	nextRecord := make(map[string]string)

	row, err := i.nextRow()
	if err != nil {
		return nil, err
	}

	for columnName, idx := range i.reverseHeaders {
		nextRecord[columnName] = row[idx]
	}

	idx := 0
	for idx < i.aType.NumField() {
		field := i.aType.Field(idx)
		csvColumn := field.Tag.Get("csvcake")

		thisField := i.aConcrete.Field(idx)
		fieldType := fmt.Sprintf("%s", field.Type)
		switch fieldType {
		case "string":
			thisField.SetString(nextRecord[csvColumn])
		case "int64":
			var i64AsInt int64
			i64AsInt, _ = strconv.ParseInt(nextRecord[csvColumn], 10, 64)
			thisField.SetInt(i64AsInt)
		}
		idx++

	}
	return i.aConcrete.Interface(), nil
}
