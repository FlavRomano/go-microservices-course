package main

import (
	"log"
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"
)

func Test_routest_exists(t *testing.T) {
	// we can ignore DB and Models because routes doesn't interact with DB
	testApp := Config{}
	testRoutest := testApp.routes()
	chiRoutes := testRoutest.(chi.Router)

	routes := []string{
		"/auth", // we got just this route
	}

	for _, route := range routes {
		routeExists(t, chiRoutes, route)
	}
}

func routeExists(t *testing.T, routes chi.Router, route string) {
	found := false

	_ = chi.Walk(
		routes,
		func(method, foundRoute string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
			log.Println("foundRoute =", foundRoute)
			if route == foundRoute {
				found = true
			}
			return nil
		},
	)

	if !found {
		t.Errorf("did not find %s in registered routes", route)
	}
}
