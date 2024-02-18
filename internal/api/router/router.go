package router

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/stooezy/dymplexgo/internal/api"
	"github.com/stooezy/dymplexgo/web"
)

func Init(s *api.Server) {
	mux := http.NewServeMux()
	pageIndex := web.PageIndex()
	mux.Handle("/", templ.Handler(pageIndex))

	s.Http.Handler = mux
}
