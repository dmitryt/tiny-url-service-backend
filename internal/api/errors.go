package api

import "errors"

var (
	ErrFileCreate        = errors.New("error during creating the file")
	ErrFileRead          = errors.New("error during reading the file")
	ErrFileWrite         = errors.New("error during writing the file")
	ErrContentUnmarshal  = errors.New("error during unmarshaling the content")
	ErrContentMarshal    = errors.New("error during marshaling the content")
	ErrParsingBody       = errors.New("error during parsing the request body")
	ErrValidationAddItem = errors.New("error, incoming data is empty")
)
