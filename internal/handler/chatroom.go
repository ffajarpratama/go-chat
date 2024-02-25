package handler

import (
	"net/http"

	"github.com/ffajarpratama/go-chat/internal/http/request"
	"github.com/ffajarpratama/go-chat/pkg/constant"
	base_request "github.com/ffajarpratama/go-chat/pkg/http/request"
	responser "github.com/ffajarpratama/go-chat/pkg/http/response"
	"github.com/ffajarpratama/go-chat/pkg/util"
	"github.com/google/uuid"
)

func (h *handler) CreateChatroom(w http.ResponseWriter, r *http.Request) {
	var req request.CreateChatroomReq

	err := h.v.ValidateStruct(r, &req)
	if err != nil {
		responser.Error(w, err)
		return
	}

	err = h.uc.CreateChatroom(r.Context(), &req)
	if err != nil {
		responser.Error(w, err)
		return
	}

	responser.OK(w, nil)
}

func (h *handler) FindAndCountChatroom(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var params request.ListChatroomQuery
	params.Query = base_request.NewBaseQuery(r)
	params.Topic = constant.ChatroomTopic(r.URL.Query().Get("topic"))
	params.UserID, _ = uuid.Parse(util.UserIDContext(ctx))

	res, cnt, err := h.uc.FindAndCountChatroom(ctx, &params)
	if err != nil {
		responser.Error(w, err)
		return
	}

	responser.Paging(w, res, params.Page, params.Limit, cnt)
}
