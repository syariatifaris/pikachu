package handler

import (
	"net/http"

	"github.com/syariatifaris/arkeus/core/net"
	"github.com/syariatifaris/arkeus/core/framework/handler"
	"github.com/syariatifaris/pikachu/app/module/page/repo"
)

func NewPageHandler(repo repo.PageRepository) *PageHandler {
	return &PageHandler{
		repo: repo,
	}
}

type PageHandler struct {
	repo repo.PageRepository
	handler.BaseHandler
}

func (*PageHandler) Name() string {
	return "PageHandler"
}

func (s *PageHandler) RegisterHandlers(router net.Router) {
	pageRouter := router.PathPrefix("/page").Subrouter()
	pageRouter.HandleFunc("/all", s.NoAuthenticate(s.Index)).Methods(http.MethodGet)
}

func (s *PageHandler) Index(r *http.Request) (interface{}, error) {
	return s.repo.GetAll(), nil
}
