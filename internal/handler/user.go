package handler

import (
	"net/http"

	"github.com/ffajarpratama/go-chat/internal/http/request"
	"github.com/ffajarpratama/go-chat/pkg/constant"
	base_request "github.com/ffajarpratama/go-chat/pkg/http/request"
	responser "github.com/ffajarpratama/go-chat/pkg/http/response"
)

func (h *handler) FindAndCountUser(w http.ResponseWriter, r *http.Request) {
	var params request.ListUserQuery
	params.Query = base_request.NewBaseQuery(r)
	params.Role = constant.UserRole(r.URL.Query().Get("role"))

	res, cnt, err := h.uc.FindAndCountUser(r.Context(), &params)
	if err != nil {
		responser.Error(w, err)
		return
	}

	responser.Paging(w, res, params.Page, params.Limit, cnt)
}
