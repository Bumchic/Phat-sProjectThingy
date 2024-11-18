package routes

import (
	"net/http"
	"server/components/appContext"
)

var (
	BaseRoute = "/v1"
)

func InitRoutes(ctx *appContext.AppContext, v1 *http.ServeMux) {

	v1.Handle(BaseRoute+"/", http.StripPrefix(BaseRoute, ServicesRoutes(ctx.DB)))
	//
}
