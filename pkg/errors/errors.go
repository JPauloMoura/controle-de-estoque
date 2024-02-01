package errors

import "errors"

var (
	ErrorInvalidId                     = errors.New("invalid id")
	ErrorInvalidHttpMethod             = errors.New("invalid http method")
	ErrorInvalidProductFieldsJson      = errors.New("invalid product fields json")
	ErrorUnableToGetProducts           = errors.New("unable to get product")
	ErrorUnableToListProducts          = errors.New("unable to list products")
	ErrorUnableToInsertProduct         = errors.New("unable to insert product")
	ErrorUnableToScanProduct           = errors.New("unable to scan product")
	ErrorUnableToDeleteProduct         = errors.New("unable to delete product")
	ErrorUnableToUpdateProduct         = errors.New("unable to update product")
	ErrorQueryToInsertProductIsInvalid = errors.New("query to insert product is invalid")
	ErrorQueryToDeleteProductIsInvalid = errors.New("query to delete product is invalid")
	ErrorQueryToUpdateProductIsInvalid = errors.New("query to update product is invalid")
)
