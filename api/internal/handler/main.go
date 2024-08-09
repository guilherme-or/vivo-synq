package handler

import (
	"net/http"

	"github.com/guilherme-or/vivo-synq/api/internal/database"
)

type UserProductsHandler struct {
	conn database.DatabaseConn
}

func NewUserProductsHandler(c database.DatabaseConn) *UserProductsHandler {
	return &UserProductsHandler{conn: c}
}

func (h *UserProductsHandler) GetUserProductsByID(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
