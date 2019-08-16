package home

import (
	"net/http"

	"github.com/hhly234/passport-tracking/internal/pkg/types"
)

const (
	get = http.MethodGet
)

// NewRouter append new router
func NewRouter(r *[]types.Route) {
	routes := []types.Route{
		{
			Path:    "/",
			Method:  get,
			Handler: newRouter().indexHandler,
		},
		{
			Path:    "/about",
			Method:  get,
			Handler: newRouter().aboutHandler,
		},
	}

	*r = append(*r, routes...)
}
