package cards

import (
	"encoding/json"
	"net/http"
	"strconv"

	"com.redberry.api/domain/entities"
	"com.redberry.api/domain/services/cards"
)

type Controller struct {
	service *cards.Service
}

func New() *Controller {
	return &Controller{
		service: cards.New(),
	}
}

func (controller Controller) Create(w http.ResponseWriter, r *http.Request) {
	var requestBody entities.Card

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	card, err := controller.service.Create(requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(card)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(response)
	w.WriteHeader(http.StatusCreated)
}

func (controller Controller) GetByID(w http.ResponseWriter, r *http.Request) {
	cardID, err := strconv.Atoi(r.PathValue("card_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	card, err := controller.service.GetByID(cardID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(card)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(response)
	w.WriteHeader(http.StatusOK)
}

func (controller Controller) Update(w http.ResponseWriter, r *http.Request) {
	cardID, err := strconv.Atoi(r.PathValue("card_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var requestBody entities.Card

	err = json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	card, err := controller.service.Update(cardID, requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(card)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(response)
	w.WriteHeader(http.StatusOK)
}

func (controller Controller) Delete(w http.ResponseWriter, r *http.Request) {
	cardID, err := strconv.Atoi(r.PathValue("card_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	card, err := controller.service.Delete(cardID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(card)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(response)
	w.WriteHeader(http.StatusOK)
}
