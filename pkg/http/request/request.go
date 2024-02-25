package request

import (
	"net/http"
	"strconv"
)

type Query struct {
	Page    int    `query:"page"`
	Limit   int    `query:"limit"`
	Keyword string `query:"keyword"`
	Sort    string `query:"sort"`
}

const (
	DefaultPage  int = 1
	DefaultLimit int = 20
)

func NewBaseQuery(r *http.Request) Query {
	page := DefaultPage
	limit := DefaultLimit

	if r.URL.Query().Get("page") != "" {
		page, _ = strconv.Atoi(r.URL.Query().Get("page"))
	}

	if r.URL.Query().Get("limit") != "" {
		limit, _ = strconv.Atoi(r.URL.Query().Get("limit"))
	}

	return Query{
		Page:    page,
		Limit:   limit,
		Keyword: r.URL.Query().Get("keyword"),
		Sort:    r.URL.Query().Get("sort"),
	}
}
