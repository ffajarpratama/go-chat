package recover

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"runtime"

	"github.com/ffajarpratama/go-chat/pkg/constant"
)

const (
	MAX_BYTE = 2084
)

type PanicJsonResponse struct {
	Success bool                `json:"success"`
	Paging  interface{}         `json:"paging"`
	Data    interface{}         `json:"data"`
	Errors  *PanicErrorResponse `json:"errors"`
}

type PanicErrorResponse struct {
	Code    int    `json:"code"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func RecoverWrap(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				buf := make([]byte, MAX_BYTE)
				n := runtime.Stack(buf, false)
				buf = buf[:n]

				fmt.Printf("[ERROR::RECOVER] %v\n", err)
				fmt.Printf("[StackTrace::] \n%s\n", buf)

				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(PanicJsonResponse{
					Errors: &PanicErrorResponse{
						Code:    constant.DefaultUnhandledError,
						Status:  http.StatusInternalServerError,
						Message: constant.ErrorMessageMap[constant.DefaultUnhandledError],
					},
				})
			}
		}()

		h.ServeHTTP(w, r)
	})
}

func GetPanicErrorMsg(err interface{}) string {
	v := reflect.ValueOf(err)

	//nolint // no need all case
	switch v.Kind() {
	case reflect.String:
		return err.(string)
	default:
		return err.(error).Error()
	}
}
