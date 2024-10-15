package boards

import (
	"encoding/json"
	"net/http"
	"strconv"

	"com.redberry.api/domain/entities"
	"com.redberry.api/domain/services/boards"
)

type Controller struct {
	service *boards.Service
}

func New() *Controller {
	return &Controller{
		service: boards.New(),
	}
}

func (controller Controller) Create(w http.ResponseWriter, r *http.Request) {
	var requestBody entities.Board

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	board, err := controller.service.Create(requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(board)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(response)
	w.WriteHeader(http.StatusCreated)
}

func (controller Controller) GetByID(w http.ResponseWriter, r *http.Request) {
	boardID, err := strconv.Atoi(r.PathValue("board_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	board, err := controller.service.GetByID(boardID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(board)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(response)
	w.WriteHeader(http.StatusOK)
}

func (controller Controller) Update(w http.ResponseWriter, r *http.Request) {
	boardID, err := strconv.Atoi(r.PathValue("board_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var requestBody entities.Board

	err = json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	board, err := controller.service.Update(boardID, requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(board)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(response)
	w.WriteHeader(http.StatusOK)
}

func (controller Controller) Delete(w http.ResponseWriter, r *http.Request) {
	boardID, err := strconv.Atoi(r.PathValue("board_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	board, err := controller.service.Delete(boardID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(board)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(response)
	w.WriteHeader(http.StatusOK)
}
