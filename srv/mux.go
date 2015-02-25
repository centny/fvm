package srv

import (
	"github.com/Centny/fvm/api"
	"github.com/Centny/gwf/routing"
	"github.com/Centny/gwf/routing/filter"
	"net/http"
)

func HSrvMux(smux *http.ServeMux, pre string, www string) {
	mux := routing.NewSessionMux2(pre)
	// mux.ShowLog = true
	cors := filter.NewCORS()
	cors.AddSite("*")
	mux.HFilter("^/.*$", cors)
	//
	//
	mux.HFunc("^/api/uload(\\?.*)?$", api.ULoad)
	mux.Handler("^/.*$", http.FileServer(http.Dir(www)))
	//
	smux.Handle("/", mux)
}
