package common

import (
	"net/http"
)

const (
	EmptyObject ErrorDataType = iota
	EmptyList
	EmptyString
	Zero

	ErrorBindingRequest    LogMessage = "error_binding_request"
	ErrorValidationRequest LogMessage = "error_validation_request"
	ErrorGeneral           LogMessage = "error_general"
)

func (e LogMessage) String() string {
	return string(e)
}

type (
	ErrorDataType int
	LogMessage    string

	Response struct {
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
)

func NewResponse(msg string, data interface{}) (res Response) {
	res.Message = msg
	res.Data = data
	return
}

func NewResponseOK(data interface{}) (res Response) {
	res.Message = http.StatusText(http.StatusOK)
	res.Data = data
	return
}

func NewResponseCreated(data interface{}) (res Response) {
	res.Message = http.StatusText(http.StatusCreated)
	res.Data = data
	return
}

func getDataType(dataType ErrorDataType) interface{} {
	switch dataType {
	case EmptyList:
		return []string{}
	case EmptyString:
		return ""
	case Zero:
		return 0
	default:
		return map[string]interface{}{}
	}
}

func NewBadRequestResponse(dataType ErrorDataType) (res Response) {
	res.Message = http.StatusText(http.StatusBadRequest)
	res.Data = getDataType(dataType)
	return
}

func NewNotFoundResponse(dataType ErrorDataType) (res Response) {
	res.Message = http.StatusText(http.StatusNotFound)
	res.Data = getDataType(dataType)
	return
}

func NewForbiddenResponse(dataType ErrorDataType) (res Response) {
	res.Message = http.StatusText(http.StatusForbidden)
	res.Data = getDataType(dataType)
	return
}

func NewUnauthorizedResponse(dataType ErrorDataType) (res Response) {
	res.Message = http.StatusText(http.StatusUnauthorized)
	res.Data = getDataType(dataType)
	return
}

func NewInternalServerErrorResponse(dataType ErrorDataType) (res Response) {
	res.Message = http.StatusText(http.StatusInternalServerError)
	res.Data = getDataType(dataType)
	return
}
