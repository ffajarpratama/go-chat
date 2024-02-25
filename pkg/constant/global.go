package constant

import "net/http"

type JwtKey string

const (
	UserIDKey JwtKey = "userID"
	RoleKey   JwtKey = "role"

	DefaultUnhandledError = 1000 + iota
	DefaultNotFoundError
	DefaultBadRequestError
	DefaultUnauthorizedError
	DefaultDuplicateDataError
)

var ErrorMessageMap = map[int]string{
	DefaultUnhandledError:     "something went wrong with our side, please try again",
	DefaultNotFoundError:      "data not found",
	DefaultUnauthorizedError:  "you are not authorized to access this api",
	DefaultDuplicateDataError: "your request body was unprocessable, please modify it",
}

var InternalResponseCodeMap = map[int]int{
	http.StatusUnprocessableEntity: DefaultDuplicateDataError,
	http.StatusInternalServerError: DefaultUnhandledError,
	http.StatusNotFound:            DefaultNotFoundError,
	http.StatusUnauthorized:        DefaultUnauthorizedError,
}
