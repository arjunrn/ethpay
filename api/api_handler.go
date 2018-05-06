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
		case *json.UnmarshalTypeError:
			w.WriteHeader(http.StatusBadRequest)
			errorMessage = "Invalid Input"
		default:
			log.Errorf("Unknown Error: %v %T", err, err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		encoder := json.NewEncoder(w)
		encoder.Encode(errorResponse{Error: errorMessage})
		return
	}
}
