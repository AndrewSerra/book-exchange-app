package utils

import (
	"fmt"
	"reflect"
)

type MySQLErrorCode = uint16

const (
	MYSQL_DUPLICATE_ERROR MySQLErrorCode = 1062
)

// NOT FOUND
type DataNotFoundError struct {
	Data interface{}
	Err  error
}

func (e *DataNotFoundError) Error() string {
	return fmt.Sprintf("Content cannot be found: %v\n", e.Data)
}

// UNKNOWN
type UnknownError struct {
	Err error
}

func (e *UnknownError) Error() string {
	return fmt.Sprintln("Unknown server error")
}

// DATA EXISTS - Conflict
type DataExistsError struct {
	Data interface{}
	Err  error
}

func (e *DataExistsError) Error() string {
	return fmt.Sprintf("Data already exists: %v", e.Data)
}

// UNRECOGNIZED TYPE
type UnrecognizedTypeError struct {
	Data interface{}
	Err  error
}

func (e *UnrecognizedTypeError) Error() string {
	return fmt.Sprintf("Unrecognized type: %v", reflect.TypeOf(e.Data))
}

// MALFORMED DATA
type MalformedDataError struct {
	Data interface{}
	Err  error
}

func (e *MalformedDataError) Error() string {
	return fmt.Sprintf("Malformed data type: %s", e.Data)
}

// QUERY PROCESSING
type QueryProcessingError struct {
	Data interface{}
	Err  error
}

func (e *QueryProcessingError) Error() string {
	return fmt.Sprintf("Query processing failure: %s", e.Data)
}
