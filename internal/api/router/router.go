package router

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/stooezy/dymplex/internal/api"
	"github.com/stooezy/dymplex/web"
)

func Init(s *api.Server) {
	mux := http.NewServeMux()
	pageIndex := web.PageIndex("Agoy")
	mux.Handle("/", templ.Handler(pageIndex))

	s.Http.Handler = mux
}
