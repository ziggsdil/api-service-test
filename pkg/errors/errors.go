package errors

import "net/http"

// interface
type Error interface {
	Error() string
	Code() int
}

// base errors
type NotFoundError struct{}

func (e NotFoundError) Error() string {
	return "NOT_FOUND"
}

func (e NotFoundError) Code() int {
	return http.StatusNotFound
}

type ConflictError struct{}

func (e ConflictError) Error() string {
	return "CONFLICT"
}

func (e ConflictError) Code() int {
	return http.StatusConflict
}

type InternalError struct{}

func (e InternalError) Error() string {
	return "INTERNAL_SERVER_ERROR"
}

func (e InternalError) Code() int {
	return http.StatusInternalServerError
}

type BadRequest struct{}

func (e BadRequest) Error() string {
	return "BAD_REQUEST"
}

func (e BadRequest) Code() int {
	return http.StatusBadRequest
}

// custom errors

type NameCannotChangeError struct{}

func (e NameCannotChangeError) Error() string {
	return "NAME_CANNOT_CHANGE"
}

func (e NameCannotChangeError) Code() int {
	return http.StatusBadRequest
}
