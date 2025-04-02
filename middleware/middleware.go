package middleware

import (
	"myapp/data"

	"github.com/mirei965/framinGo"
)

type Middleware struct {
	App    *framinGo.FraminGo
	Models data.Models
}
