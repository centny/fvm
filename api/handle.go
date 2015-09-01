package api

import (
	"github.com/Centny/gwf/routing"
	"github.com/Centny/gwf/routing/filter"
)

func Handle(mux *routing.SessionMux) {
	mux.HFilterFunc("^/api/uload(\\?.*)?$", filter.ParseQuery)
	mux.HFunc("^/api/uload(\\?.*)?$", ULoad)
	mux.HFunc("^/raw.*$", Raw)
}
