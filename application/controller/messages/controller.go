package messages

import (
	"encoding/json"
	"net/http"
	"strconv"

	"com.redberry.api/domain/entities"
	"com.redberry.api/domain/services/messages"
)

type MessagesController struct {
	service *messages.Service
}

func NewMessagesController() *MessagesController {
	return &MessagesController{
		service: messages.New(),
	}
}

// Create lida com a criação de uma nova mensagem.
func (controller *MessagesController) Create(w http.ResponseWriter, r *http.Request) {
	var requestBody entities.Message

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := controller.service.Create(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}

// GetByID lida com a recuperação de uma mensagem pelo ID.
func (controller *MessagesController) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	message, err := controller.service.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(message)
}

// Update lida com a atualização de uma mensagem existente.
func (controller *MessagesController) Update(w http.ResponseWriter, r *http.Request) {
	var requestBody entities.Message

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = controller.service.Update(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Delete lida com a remoção de uma mensagem pelo ID.
func (controller *MessagesController) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	err = controller.service.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetAll lida com a recuperação de todas as mensagens.
func (controller *MessagesController) GetAll(w http.ResponseWriter, r *http.Request) {
	messages, err := controller.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(messages)
}
