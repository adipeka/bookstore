package handler

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

type handler struct {
	db     *sql.DB
	router *mux.Router
}

func New(db *sql.DB, r *mux.Router) handler {
	return handler{
		db:     db,
		router: r,
	}
}

func (h *handler) GetBook(w http.ResponseWriter, r *http.Request) {

}
