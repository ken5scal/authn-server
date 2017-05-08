package handlers

import (
	"encoding/json"
	"net/http"
)

type health struct {
	Http bool `json:"http"`
	Db   bool `json:"db"`
}

func (app App) Health(w http.ResponseWriter, req *http.Request) {
	db := true
	err := app.Db.Ping()
	if err != nil {
		db = false
	}

	h := health{true, db}

	j, err := json.Marshal(h)
	if err != nil {
		panic("TODO: gorilla RecoveryHandler")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}