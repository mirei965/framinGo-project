package handlers

import (
	"myapp/data"
	"net/http"
	"time"

	"github.com/CloudyKit/jet/v6"
	"github.com/mirei965/framingo"
)

type Handlers struct {
	App    *framingo.Framingo
	Models data.Models
}

//const

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	defer h.App.LoadTime(time.Now())
	err := h.render(w, r, "home", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}

