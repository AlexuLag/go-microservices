package main

import (
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"
)

func test_routes_exist(t *testing.T) {
	testApp := Config{}
	testRoutes := testApp.routes()
	actualRoutes := testRoutes.(chi.Router)

	routesToTest := []string{"/authenticate"}

	for _, route := range routesToTest {
		routeExists(t, actualRoutes, route)
	}

}

func routeExists(t *testing.T, routes chi.Router, route string) {
	found := false

	_ = chi.Walk(routes, func(method string,
		foundRoute string,
		handler http.Handler,
		middlewares ...func(http.Handler) http.Handler) error {
		if foundRoute == route {
			found = true
		}
		return nil
	})

	if !found {
		t.Errorf("expected to find route %s, but it was not in the routes", route)
	}
}
