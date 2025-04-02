package middleware

import (
	"myapp/data"
	"github.com/mirei965/framingo"
)

type Middleware struct {
	App *framingo.Framingo
	Models data.Models
}

