package handlers

import (
	"myapp/data"
	"net/http"
	"time"

	"github.com/mirei965/framinGo"
)

type Handlers struct {
	App    *framinGo.FraminGo
	Models *data.Models
}

//const

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	defer h.App.LoadTime(time.Now())
	err := h.render(w, r, "home", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}
