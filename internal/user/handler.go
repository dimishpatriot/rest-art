package user

import (
	"fmt"
	"net/http"

	"github.com/dimishpatriot/rest-art-of-development/internal/handlers"
	"github.com/dimishpatriot/rest-art-of-development/internal/logging"
	"github.com/julienschmidt/httprouter"
)

const (
	usersURL string = "/users"
	userURL  string = "/users/:uuid"
)

type handler struct {
	logger *logging.Logger
}

func NewHandler(logger *logging.Logger) handlers.Handler {
	return &handler{
		logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodDelete, userURL, h.DeleteUser)
	router.HandlerFunc(http.MethodGet, usersURL, h.GetList)
	router.HandlerFunc(http.MethodGet, userURL, h.GetUserByUUID)
	router.HandlerFunc(http.MethodPatch, userURL, h.PartiallyUpdateUser)
	router.HandlerFunc(http.MethodPost, usersURL, h.CreateUser)
	router.HandlerFunc(http.MethodPut, userURL, h.UpdateUser)
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("GetList"))
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("GetUserByUUID %s", r.URL.Query().Get("uuid"))))
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("CreateUser"))
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte(fmt.Sprintf("UpdateUser %s", r.URL.Query().Get("uuid"))))
}

func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte(fmt.Sprintf("PartiallyUpdateUser %s", r.URL.Query().Get("uuid"))))
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte(fmt.Sprintf("DeleteUser %s", r.URL.Query().Get("uuid"))))
}
