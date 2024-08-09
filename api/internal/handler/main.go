package handler

import (
	"database/sql"
	"net/http"

	"github.com/guilherme-or/vivo-synq/api/internal/database"
)

type UserProductsHandler struct {
	db *sql.DB
}

func NewUserProductsHandler(c database.SQLConn) *UserProductsHandler {
	return &UserProductsHandler{db: c.GetDatabase()}
}

func (h *UserProductsHandler) GetUserProductsByID(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("user_id")
	if userID == "" {
		writeAppErr(w, ErrInvalidArgument)
		return
	}

	w.WriteHeader(http.StatusNotImplemented)
}
