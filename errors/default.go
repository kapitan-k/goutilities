package errors

import (
	"fmt"
)

var (
	ErrBusy          = DefaultErrorNew(uint32(DefaultErrorType_BUSY), "Busy")
	ErrTimeout       = DefaultErrorNew(uint32(DefaultErrorType_TIMEOUT), "Timeout")
	ErrInvalidMT     = DefaultErrorNew(uint32(DefaultErrorType_INVALID_MESSAGE_TYPE), "Invalid msg type")
	ErrNotFoundByKey = DefaultErrorNew(uint32(DefaultErrorType_NOT_FOUND_BY_KEY), "Not found by key")
	ErrNotConnected  = DefaultErrorNew(uint32(DefaultErrorType_NOT_CONNECTED), "Not connected")
	ErrFail          = DefaultErrorNew(uint32(DefaultErrorType_FAIL), "failed to perform action")
	ErrEncoding      = DefaultErrorNew(uint32(DefaultErrorType_ENCODING), "encoding error")
	ErrOther         = DefaultErrorNew(uint32(DefaultErrorType_OTHER), "other error")
)

func (self *DefaultError) Error() string {
	return fmt.Sprintf("Error %+v", *self)
}

func DefaultErrorNew(errorCode uint32, errorString string) (self *DefaultError) {
	return &DefaultError{Code: errorCode, Str: errorString}
}

func DefaultErrorNewSprintln(errorCode uint32, a ...interface{}) (self *DefaultError) {
	return &DefaultError{Code: errorCode, Str: fmt.Sprintln(a)}
}

func DefaultErrorNewSprintlnWithFields(errorCode uint32, fields map[string]string, a ...interface{}) (self *DefaultError) {
	return &DefaultError{Code: errorCode, Str: fmt.Sprintln(a), Fields: fields}
}

type ErrorGetter interface {
	Error() error
}

type ErrorsGetter interface {
	Errors() []error
}

func DefaultErrorsToErrors(berrs []*DefaultError) []error {
	errs := make([]error, len(berrs))
	for i, berr := range berrs {
		errs[i] = berr
	}
	return errs
}
