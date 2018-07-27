package core

import (
	"github.com/syariatifaris/arkeus/core/inject"
    "github.com/syariatifaris/arkeus/core/net"
	fwhandler "github.com/syariatifaris/arkeus/core/framework/handler"
	"github.com/syariatifaris/pikachu/app/handler"
	"github.com/syariatifaris/pikachu/app/module/page/repo"
)

//gets new dependencies
func NewDependencies() inject.Injection {
	var (
		router net.Router
		pageRepo repo.PageRepository
		baseHandler  *fwhandler.BaseHandler
		pageHandler *handler.PageHandler
	)

	di := inject.NewDependencyInjection()
    di.AddDependency(&router, net.NewRouter)
    di.AddDependency(&pageRepo, repo.NewPageRepository)
    di.AddDependency(&baseHandler, fwhandler.NewSimpleBaseHandler)
    di.AddDependency(&pageHandler, handler.NewPageHandler, &pageRepo)

	return di
}
