package messages

import (
	"encoding/json"
	"net/http"
	"strconv"

	"com.redberry.api/domain/entities"
	"com.redberry.api/domain/services/messages"
)

// Controller handles message-related endpoints.
type Controller struct {
	service *messages.Service
}

// New creates a new instance of the Controller.
func New() *Controller {
	return &Controller{
		service: messages.New(),
	}
}

// Create godoc
// @Summary      Create Message
// @Description  Creates a new message.
// @Tags         messages
// @Accept       json
// @Produce      json
// @Param        body body      entities.Message true "Message data"
// @Success      201  {object}  entities.Message "Created message"
// @failure      400  {string}  string "Bad Request"
// @Failure      500  {string}  string "Internal Server Error"
// @Router       /api/v1/messages [post]
func (controller Controller) Create(w http.ResponseWriter, r *http.Request) {
	var requestBody entities.Message

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	message, err := controller.service.Create(requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

// GetByID godoc
// @Summary      Get Message by ID
// @Description  Retrieves a message by its ID.
// @Tags         messages
// @Produce      json
// @Param        message_id path      int true "Message ID"
// @Success      200  {object}  entities.Message "Message data"
// @failure      400  {string}  string "Bad Request"
// @Failure      500  {string}  string "Internal Server Error"
// @Router       /api/v1/messages/{message_id} [get]
func (controller Controller) GetByID(w http.ResponseWriter, r *http.Request) {
	messageID, err := strconv.Atoi(r.PathValue("message_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	message, err := controller.service.GetByID(messageID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// Update godoc
// @Summary      Update Message
// @Description  Updates an existing message.
// @Tags         messages
// @Accept       json
// @Produce      json
// @Param        message_id path      int true "Message ID"
// @Param        body body      entities.Message true "Updated message data"
// @Success      200  {object}  entities.Message "Updated message"
// @failure      400  {string}  string "Bad Request"
// @Failure      500  {string}  string "Internal Server Error"
// @Router       /api/v1/messages/{message_id} [put]
func (controller Controller) Update(w http.ResponseWriter, r *http.Request) {
	messageID, err := strconv.Atoi(r.PathValue("message_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var requestBody entities.Message

	err = json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	message, err := controller.service.Update(messageID, requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// Delete godoc
// @Summary      Delete Message
// @Description  Deletes a message by its ID.
// @Tags         messages
// @Produce      json
// @Param        message_id path      int true "Message ID"
// @Success      200  {object}  entities.Message "Deleted message"
// @failure      400  {string}  string "Bad Request"
// @Failure      500  {string}  string "Internal Server Error"
// @Router       /api/v1/messages/{message_id} [delete]
func (controller Controller) Delete(w http.ResponseWriter, r *http.Request) {
	messageID, err := strconv.Atoi(r.PathValue("message_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	message, err := controller.service.Delete(messageID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
