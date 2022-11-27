package business

import "errors"

var (
	//ErrInvalidSpec Error when data given is not valid on update or insert
	ErrInvalidSpec = errors.New("given spec is not valid")

	//ErrNotFound Error when data is not found
	ErrNotFound = errors.New("data was not found")
)
