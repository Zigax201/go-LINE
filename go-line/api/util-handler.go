package api

import (
	"go-line/util"
	"net/http"
)

func (a *App) PingHandler(w http.ResponseWriter, r *http.Request) {
	util.RespondJSON(w, 200, map[string]string{"message": "Pong"})
}
