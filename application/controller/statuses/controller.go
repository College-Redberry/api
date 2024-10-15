package statuses

import (
	"encoding/json"
	"net/http"
	"strconv"

	"com.redberry.api/domain/entities"
	"com.redberry.api/domain/services/statuses"
)

type Controller struct {
	service *statuses.Service
}

func New() *Controller {
	return &Controller{
		service: statuses.New(),
	}
}

func (controller Controller) Create(w http.ResponseWriter, r *http.Request) {
	var requestBody entities.Status

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	status, err := controller.service.Create(requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(response)
	w.WriteHeader(http.StatusCreated)
}

func (controller Controller) GetByID(w http.ResponseWriter, r *http.Request) {
	statusID, err := strconv.Atoi(r.PathValue("status_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	status, err := controller.service.GetByID(statusID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(response)
	w.WriteHeader(http.StatusOK)
}

func (controller Controller) Update(w http.ResponseWriter, r *http.Request) {
	statusID, err := strconv.Atoi(r.PathValue("status_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var requestBody entities.Status

	err = json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	status, err := controller.service.Update(statusID, requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(response)
	w.WriteHeader(http.StatusOK)
}

func (controller Controller) Delete(w http.ResponseWriter, r *http.Request) {
	statusID, err := strconv.Atoi(r.PathValue("status_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	status, err := controller.service.Delete(statusID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(response)
	w.WriteHeader(http.StatusOK)
}
