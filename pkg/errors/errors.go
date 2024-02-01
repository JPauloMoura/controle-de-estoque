package errors

import "errors"

var (
	ErrorInvalidId                = errors.New("invalid id")
	ErrorInvalidHttpMethod        = errors.New("invalid http method")
	ErrorInvalidProductFieldsJson = errors.New("invalid product fields json")
)
