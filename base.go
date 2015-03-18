package coinbase

import "fmt"

type Base struct {
}

type CbNotFoundError struct {
}

type CbNotImplementedError struct {
}

func (e CbNotImplementedError) Error() string {
	return fmt.Sprintf("coinbase-go error: not implemented")
}

func (e CbNotFoundError) Error() string {
	return fmt.Sprintf("coinbase-go error: not found")
}

type CbError struct {
	msg string // description of error
}

func (e CbError) Error() string {
	return fmt.Sprintf("coinbase-go error: %v", e.msg)
}

type CbHttpError struct {
	code int
}

func (e CbHttpError) Error() string {
	return fmt.Sprintf("coinbase-go error: Invalid HTTP response code: %d", e.code)
}
