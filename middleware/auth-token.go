package middleware

import (
	"net/http"
)

func (m *Middleware) AuthToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_,err :=m.Models.Tokens.AuthenticateToken(r)
    if err != nil {
      var payload struct {
        Error bool `json:"error"`
        Success string `json:"message"`
      }
      payload.Error =  true  
      payload.Success = "Invalid authentication credentials"
      _ = m.App.WriteJSON(w, http.StatusUnauthorized, payload)
      return
    }
    next.ServeHTTP(w, r)
	})
}