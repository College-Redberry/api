package priorities

import (
	"encoding/json"
	"net/http"
	"strconv"

	"com.redberry.api/domain/entities"
	"com.redberry.api/domain/services/priorities"
)

type Controller struct {
	service *priorities.Service
}

func New() *Controller {
	return &Controller{
		service: priorities.New(),
	}
}

func (controller Controller) Create(w http.ResponseWriter, r *http.Request) {
	var requestBody entities.Priority

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	priority, err := controller.service.Create(requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(priority)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(response)
	w.WriteHeader(http.StatusCreated)
}

func (controller Controller) GetByID(w http.ResponseWriter, r *http.Request) {
	priorityID, err := strconv.Atoi(r.PathValue("priority_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	priority, err := controller.service.GetByID(priorityID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(priority)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(response)
	w.WriteHeader(http.StatusOK)
}

func (controller Controller) Update(w http.ResponseWriter, r *http.Request) {
	priorityID, err := strconv.Atoi(r.PathValue("priority_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var requestBody entities.Priority

	err = json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	priority, err := controller.service.Update(priorityID, requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(priority)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(response)
	w.WriteHeader(http.StatusOK)
}

func (controller Controller) Delete(w http.ResponseWriter, r *http.Request) {
	priorityID, err := strconv.Atoi(r.PathValue("priority_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	priority, err := controller.service.Delete(priorityID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(priority)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(response)
	w.WriteHeader(http.StatusOK)
}
