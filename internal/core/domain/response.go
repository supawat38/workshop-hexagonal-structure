package domain

import (
	"net/http"
)

var (
	Success             = Status{Code: http.StatusOK, Message: []string{"Success"}}
	BadRequest          = Status{Code: http.StatusBadRequest, Message: []string{"Sorry, Not responding because of incorrect syntax"}}
	Unauthorized        = Status{Code: http.StatusUnauthorized, Message: []string{"Sorry, We are not able to process your request. Please try again"}}
	Forbidden           = Status{Code: http.StatusForbidden, Message: []string{"Sorry, Permission denied"}}
	InternalServerError = Status{Code: http.StatusInternalServerError, Message: []string{"Internal Server Error"}}
	ConFlict            = Status{Code: http.StatusBadRequest, Message: []string{"Sorry, Data is conflict"}}
	FieldsPermission    = Status{Code: http.StatusBadRequest, Message: []string{"Sorry, Fields are not able to update"}}
	NotFound            = Status{Code: http.StatusNotFound, Message: []string{"Sorry, Record Not Found"}}
)

// Status struct
type Status struct {
	Code    int      `json:"code,omitempty"`
	Message []string `json:"message,omitempty"`
}
