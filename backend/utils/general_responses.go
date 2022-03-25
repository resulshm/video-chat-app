package utils

import (
	"errors"
)

type GeneralResponse struct {
	Success    bool        `json:"success"`
	StatusCode int         `json:"status"`
	ErrMsg     string      `json:"err_msg"`
	Data       interface{} `json:"data"`
}

type ResponseErrorCodeAndMessage struct {
	ErrorCode    string  `json:"error_code"`
	ErrorMessage *string `json:"error_msg,omitempty"`
}

const (
	ErrorCodeOK                  int = 200
	ErrorCodeTfaRequired         int = 250
	ErrorCodeBadRequest          int = 400
	ErrorCodeUnauthorized        int = 401
	ErrorCodeForbidden           int = 403
	ErrorCodeNotFound            int = 404
	ErrorMethodNotAllowed        int = 405
	ErrorCodeConflict            int = 409
	ErrorCodeExpired             int = 408
	ErrorCodeFileSizeTooLarge    int = 413
	ErrorCodeTooManyRequests     int = 429
	ErrorCodeInternalServerError int = 500
)

var ErrOK = errors.New("OK")
var ErrTfaRequired = errors.New("Tfa required")
var ErrBadRequest = errors.New("Bad request")
var ErrUnauthorized = errors.New("Unauthorized")
var ErrForbidden = errors.New("Forbidden")
var ErrNotFound = errors.New("Not found")
var ErrMethodNotAllowed = errors.New("Method not allowed")
var ErrConflict = errors.New("Conflict")
var ErrExpired = errors.New("Expired")
var ErrFileSizeTooLarge = errors.New("File size too large")
var ErrTooManyRequests = errors.New("OTP retry limit exceeded")
var ErrInternalServerError = errors.New("Internal server error")

const (
	ErrorMessageOK                  = "ok"
	ErrorMessageTfaRequired         = "tfa_required"
	ErrorMessageBadRequest          = "bad_request"
	ErrorMessageUnauthorized        = "unauthorized"
	ErrorMessageForbidden           = "forbidden"
	ErrorMessageNotFound            = "not_found"
	ErrorMessageMethodNotAllowed    = "method_not_allowed"
	ErrorMessageConflict            = "conflict"
	ErrorMessageExpired             = "expired"
	ErrorMessageFileSizeTooLarge    = "file_size_too_large"
	ErrorMessageTooManyRequests     = "otp_retry_limit_exceeded"
	ErrorMessageInternalServerError = "internal_server_error"
)

const (
	PG_CODE_UNIQUE_VIOLATION string = "23505"
)

func GetStatusCodeByError(err error) (statusCode int, errMsg string) {

	statusCode = 200
	errMsg = ErrorMessageOK

	switch err {
	case ErrTfaRequired:
		statusCode = 250
		errMsg = ErrorMessageTfaRequired
	case ErrBadRequest:
		statusCode = 400
		errMsg = ErrorMessageBadRequest
	case ErrUnauthorized:
		statusCode = 401
		errMsg = ErrorMessageUnauthorized
	case ErrForbidden:
		statusCode = 403
		errMsg = ErrorMessageForbidden
	case ErrNotFound:
		statusCode = 404
		errMsg = ErrorMessageNotFound
	case ErrMethodNotAllowed:
		statusCode = 405
		errMsg = ErrorMessageMethodNotAllowed
	case ErrConflict:
		statusCode = 409
		errMsg = ErrorMessageConflict
	case ErrExpired:
		statusCode = 408
		errMsg = ErrorMessageExpired
	case ErrFileSizeTooLarge:
		statusCode = 413
		errMsg = ErrorMessageFileSizeTooLarge
	case ErrTooManyRequests:
		statusCode = 429
		errMsg = ErrorMessageTooManyRequests
	case ErrInternalServerError:
		statusCode = 500
		errMsg = ErrorMessageInternalServerError
	}

	return
}
