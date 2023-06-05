package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	api "daeuniverse/functions/api/v2"
)

type Router struct {
	*chi.Mux
}

type Route struct {
	Method    string
	Pattern   string
	HandlerFn http.HandlerFunc
}

type Routes []Route

func (r *Router) Serve(port int) error {
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), r)
	if err != nil {
		return err
	}

	return nil
}

func (r *Router) RegisterMiddlewares() {
	r.Use(middleware.Logger)
}

func (r *Router) RegisterRoutes(routes Routes) {
	for _, route := range routes {
		r.MethodFunc(route.Method, route.Pattern, route.HandlerFn)
	}
}

func NewRouter() *Router {
	router := &Router{Mux: chi.NewRouter()}
	router.RegisterMiddlewares()
	router.RegisterRoutes(Routes{
		Route{"GET", "/api/v2/geodata", api.GeodataHandler},
		Route{"GET", "/api/v2/health", api.HealthcheckHandler},
		Route{"GET", "/api/v2/index", api.IndexHandler},
	})
	return router
}
