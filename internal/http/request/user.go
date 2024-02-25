package request

import (
	"github.com/ffajarpratama/go-chat/pkg/constant"
	base_request "github.com/ffajarpratama/go-chat/pkg/http/request"
)

type ListUserQuery struct {
	base_request.Query
	Role constant.UserRole
}
