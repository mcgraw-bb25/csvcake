package iterscanner

import "strconv"

type prepareString struct{}

func (p prepareString) prepare(stringAsString string) (interface{}, error) {
	return stringAsString, nil
}

type prepareInt64 struct{}

func (p prepareInt64) prepare(i64AsString string) (interface{}, error) {
	var i64AsInt int64
	i64AsInt, err := strconv.ParseInt(i64AsString, 10, 64)
	if err != nil {
		return i64AsInt, err
	}
	return i64AsInt, nil
}

func getPreparer(preparerType string, dataAsString string) interface{} {
	var preparer preparable
	switch preparerType {
	case "string":
		preparer = prepareString{}
	case "int64":
		preparer = prepareInt64{}
	}
	value, err := preparer.prepare(dataAsString)
	if err != nil {
		return nil
	}
	return value
}
