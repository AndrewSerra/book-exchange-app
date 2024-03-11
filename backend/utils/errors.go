package utils

import "fmt"

type DataNotFoundError struct {
	Data interface{}
	Err  error
}

func (e *DataNotFoundError) Error() string {
	return fmt.Sprintf("Content cannot be found: %v\n", e.Data)
}

type UnknownError struct {
	Err error
}

func (e *UnknownError) Error() string {
	return fmt.Sprintln("Unknown server error")
}
