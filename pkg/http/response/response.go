package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"

	"github.com/ffajarpratama/go-chat/pkg/constant"
	custom_error "github.com/ffajarpratama/go-chat/pkg/error"
)

const (
	ContentType     string = "Content-Type"
	ApplicationJson string = "application/json"
)

type JsonResponse struct {
	Success bool           `json:"success"`
	Paging  *Pagination    `json:"paging"`
	Data    interface{}    `json:"data"`
	Error   *ErrorResponse `json:"error"`
}

type Pagination struct {
	Page      int   `json:"page"`
	Limit     int   `json:"limit"`
	Count     int64 `json:"count"`
	TotalPage int   `json:"total_page"`
	Next      bool  `json:"next"`
	Prev      bool  `json:"prev"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (er *ErrorResponse) Error() string {
	return "validate.request"
}

func OK(w http.ResponseWriter, data interface{}) {
	w.Header().Set(ContentType, ApplicationJson)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(JsonResponse{
		Data:    data,
		Success: true,
	})
}

func Paging(w http.ResponseWriter, list interface{}, page, limit int, cnt int64) {
	var paging *Pagination

	total := calculateTotalPage(cnt, limit)
	if page > 0 {
		paging = &Pagination{
			Page:      page,
			Limit:     limit,
			Count:     cnt,
			TotalPage: total,
			Next:      hasNext(page, total),
			Prev:      hasPrev(page),
		}
	}

	w.Header().Set(ContentType, ApplicationJson)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(JsonResponse{
		Success: true,
		Paging:  paging,
		Data:    list,
	})
}

func Error(w http.ResponseWriter, err error) {
	v, isValidationErr := err.(*ErrorResponse)
	if isValidationErr {
		w.Header().Set(ContentType, ApplicationJson)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(JsonResponse{
			Error: &ErrorResponse{
				Code:    v.Code,
				Status:  v.Status,
				Message: v.Message,
			},
		})

		return
	}

	e, isCustomErr := err.(*custom_error.CustomError)
	if !isCustomErr {
		if err != nil && !errors.Is(err, context.Canceled) {
			// bugsnag.Notify(err)
			fmt.Println(err.Error(), "[unhandled-error]")
		}

		w.Header().Set(ContentType, ApplicationJson)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(JsonResponse{
			Error: &ErrorResponse{
				Code:    constant.DefaultUnhandledError,
				Status:  http.StatusInternalServerError,
				Message: constant.ErrorMessageMap[constant.DefaultUnhandledError],
			},
		})

		return
	}

	httpCode := http.StatusInternalServerError
	internalCode := constant.DefaultUnhandledError
	msg := constant.ErrorMessageMap[constant.DefaultUnhandledError]

	if e.ErrorContext != nil && e.ErrorContext.HTTPCode > 0 {
		httpCode = e.ErrorContext.HTTPCode
		internalCode = constant.InternalResponseCodeMap[httpCode]
		msg = constant.ErrorMessageMap[internalCode]

		if e.ErrorContext.Message != "" {
			msg = e.ErrorContext.Message
		}
	}

	w.Header().Set(ContentType, ApplicationJson)
	w.WriteHeader(httpCode)
	json.NewEncoder(w).Encode(JsonResponse{
		Error: &ErrorResponse{
			Code:    internalCode,
			Status:  httpCode,
			Message: msg,
		},
	})
}

func UnauthorizedError(w http.ResponseWriter) {
	w.Header().Set(ContentType, ApplicationJson)
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(JsonResponse{
		Data:    nil,
		Success: false,
		Error: &ErrorResponse{
			Code:    constant.DefaultUnauthorizedError,
			Status:  http.StatusUnauthorized,
			Message: constant.ErrorMessageMap[constant.DefaultUnauthorizedError],
		},
	})
}

func hasNext(currentPage, totalPages int) bool {
	return currentPage < totalPages
}

func hasPrev(currentPage int) bool {
	return currentPage > 1
}

func calculateTotalPage(cnt int64, limit int) (total int) {
	return int(math.Ceil(float64(cnt) / float64(limit)))
}
