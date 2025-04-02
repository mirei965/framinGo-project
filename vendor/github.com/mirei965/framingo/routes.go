package framingo

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (f *Framingo) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	if f.Debug {
		mux.Use(middleware.Logger)
	}
	mux.Use(middleware.Recoverer)
	mux.Use(f.SessionLoad)
	mux.Use(f.NoSurf)

	return mux

}
