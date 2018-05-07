package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type APIHandler interface {
	ProcessRequest(w http.ResponseWriter, r *http.Request) error
}

type errorResponse struct {
	Error string `json:"error"`
}
type WrappedErrorHandler struct {
	handler APIHandler
}

func NewWrappedErrorHandler(h APIHandler) WrappedErrorHandler {
	return WrappedErrorHandler{handler: h}
}

type NotFoundError struct {
	model string
}

type ErrNotAvailable struct {
	service string
}

func (e ErrNotAvailable) Error() string {
	return fmt.Sprintf("%s currently unavailble", e.service)
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("%s not found", e.model)
}

func (h WrappedErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := h.handler.ProcessRequest(w, r)
	errorMessage := "Internal Server Error"
	if err != nil {
		switch err.(type) {
		case NotFoundError:
			w.WriteHeader(http.StatusNotFound)
			errorMessage = "Not Found"
			break
		case *json.UnmarshalTypeError:
			w.WriteHeader(http.StatusBadRequest)
			errorMessage = "Invalid Input"
			break
		case *ErrNotAvailable:
			w.WriteHeader(http.StatusServiceUnavailable)
			errorMessage = fmt.Sprintf("Service %s currently Unavailable", err.(ErrNotAvailable).service)
		default:
			log.Errorf("Unknown Error: %v %T", err, err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		encoder := json.NewEncoder(w)
		encoder.Encode(errorResponse{Error: errorMessage})
		return
	}
}
