package shared

import (
	"errors"
	"fmt"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

var (
	ErrJSONInvalid = errors.New("json invalid")
	ErrDatabase    = errors.New("db error")
	ErrDatabaseTx  = errors.New("db transaction error")
)

type InternalCode int

const (
	BadRequest = iota + 1
	Unauthorized
	Forbidden
	NotFound
	Gone
	Internal
	BadGateway
)

type CodeMap map[int]int

var (
	CodesMap = CodeMap{
		BadRequest:   http.StatusBadRequest,
		Unauthorized: http.StatusUnauthorized,
		Forbidden:    http.StatusForbidden,
		NotFound:     http.StatusNotFound,
		Gone:         http.StatusGone,
		Internal:     http.StatusInternalServerError,
		BadGateway:   http.StatusBadGateway,
	}
	DefaultApiError = NewApiError("", 0, "", "", nil)
	DefaultDBError  = NewDBError("", nil)
)

type AppLevel string

const (
	ServiceLevel        = "Service"
	TransportLevel      = "Transport"
	AuthMiddlewareLevel = "Auth Endpoint Middleware"
	TxMiddlewareLevel   = "Tx Endpoint Middleware"
)

func GetApiError(err error) *ApiError {
	if err != nil {
		thError := DefaultApiError
		if errors.Is(err, thError) && errors.As(err, &thError) {
			return thError.(*ApiError)
		}
	}
	return nil
}

func NewApiError(text string, code InternalCode, op string, level AppLevel, err error) error {
	return &ApiError{
		internalCode: code,
		Message:      text,
		Op:           op,
		Level:        level,
		err:          err,
		ErrStack:     "",
	}
}

type ApiError struct {
	internalCode InternalCode
	Message      string
	Op           string
	Level        AppLevel
	err          error
	ErrStack     string
}

func (ae *ApiError) Error() string {
	return ae.Message
}

func (ae *ApiError) Code() int {
	return int(ae.internalCode)
}

func (ae *ApiError) Unwrap() error {
	return ae.err
}

func (ae *ApiError) Is(e error) bool {
	_, ok := e.(*ApiError)
	return ok
}

func (ae *ApiError) As(e error) bool {
	_, ok := e.(*ApiError)
	return ok
}

func (ae *ApiError) GetError() error {
	return ae.err
}

func (ae *ApiError) Stack() error {
	err := ae.GetError()
	stack := ""
	count := 0
	for err != nil {
		count++
		stack += fmt.Sprintf("%d-: %s ", count, err.Error())
		err = errors.Unwrap(err)
	}
	ae.ErrStack = stack
	return ae
}

func NewDBError(message string, err error) error {
	return &DBError{
		Message: message,
		Err:     err,
	}
}

type DBError struct {
	Message string
	Err     error
}

func (db *DBError) Error() string {
	return db.Message
}

func (db *DBError) Unwrap() error {
	return db.Err
}

func (db *DBError) Is(e error) bool {
	_, ok := e.(*DBError)
	return ok
}

func HandleDbError(err error, message string, op string) error {
	if !errors.Is(err, DefaultDBError) {
		err = NewApiError(message, NotFound, op, ServiceLevel, err)
	} else {
		err = NewApiError(err.Error(), Internal, op, ServiceLevel, err)
	}
	return err
}
