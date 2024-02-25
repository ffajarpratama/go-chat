package util

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/ffajarpratama/go-chat/pkg/constant"
)

func GetTokenFromHeader(r *http.Request) (token string, err error) {
	header := r.Header.Get("Authorization")
	if header == "" {
		err = errors.New("token is empty")
		return token, err
	}

	s := strings.Split(header, " ")
	if len(s) != 2 {
		err = errors.New("token is invalid")
		return token, err
	}

	token = s[1]
	return token, nil
}

func UserIDContext(ctx context.Context) (userID string) {
	if ctx == nil {
		return
	}

	if userID, ok := ctx.Value(constant.UserIDKey).(string); ok {
		return userID
	}

	return
}

func RoleContext(ctx context.Context) (role string) {
	if ctx == nil {
		return
	}

	if role, ok := ctx.Value(constant.RoleKey).(string); ok {
		return role
	}

	return
}
