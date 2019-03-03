package mapreader

import (
	"encoding/csv"
	"io"
	"os"
)

// CSVHeader is a map that allows a lookup of an index
// to show what the column name is.
type CSVHeader map[int]string

// NewCSVHeader creates a new CSVHeader.
func NewCSVHeader() CSVHeader {
	newCSVHeader := make(map[int]string)
	return newCSVHeader
}

// CSVRecord is represents each line in a csv file.
type CSVRecord map[string]string

// NewCSVRecord creates a new CSVRecord
func NewCSVRecord() CSVRecord {
	newCSVRecord := make(map[string]string)
	return newCSVRecord
}

// MapReader returns an array of CSVRecords given a filename
func MapReader(filename string) ([]CSVRecord, error) {
	csvFile, _ := os.Open(filename)
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	csvHeader := NewCSVHeader()
	csvRecords := make([]CSVRecord, 0)

	headerRow := false
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return csvRecords, err
		}
		if !headerRow {
			for index, column := range row {
				csvHeader[index] = column
			}
			headerRow = true
		} else {
			csvRecord := NewCSVRecord()
			for index, column := range row {
				columnName := csvHeader[index]
				csvRecord[columnName] = column
			}
			csvRecords = append(csvRecords, csvRecord)
		}
	}
	return csvRecords, nil
}
