package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"daeuniverse/functions/api"
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
		Route{"GET", "/api/geodata", api.GeodataHandler},
		Route{"GET", "/api/health", api.HealthcheckHandler},
	})
	return router
}