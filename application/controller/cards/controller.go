package cards

import (
	"encoding/json"
	"net/http"
	"strconv"

	"com.redberry.api/domain/entities"
	"com.redberry.api/domain/services/cards"
)

// Controller handles card-related endpoints.
type Controller struct {
	service *cards.Service
}

// New creates a new instance of the Controller.
func New() *Controller {
	return &Controller{
		service: cards.New(),
	}
}

// Create godoc
// @Summary      Create Card
// @Description  Creates a new card.
// @Tags         cards
// @Accept       json
// @Produce      json
// @Param        body body      entities.Card true "Card data"
// @Success      201  {object}  entities.Card "Created card"
// @failure      400  {string}  string "Bad Request"
// @Failure      500  {string}  string "Internal Server Error"
// @Router       /api/v1/cards [post]
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

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

// GetByID godoc
// @Summary      Get Card by ID
// @Description  Retrieves a card by its ID.
// @Tags         cards
// @Produce      json
// @Param        card_id path      int true "Card ID"
// @Success      200  {object}  entities.Card "Card data"
// @failure      400  {string}  string "Bad Request"
// @Failure      500  {string}  string "Internal Server Error"
// @Router       /api/v1/cards/{card_id} [get]
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

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// Update godoc
// @Summary      Update Card
// @Description  Updates an existing card.
// @Tags         cards
// @Accept       json
// @Produce      json
// @Param        card_id path      int true "Card ID"
// @Param        body body      entities.Card true "Updated card data"
// @Success      200  {object}  entities.Card "Updated card"
// @failure      400  {string}  string "Bad Request"
// @Failure      500  {string}  string "Internal Server Error"
// @Router       /api/v1/cards/{card_id} [put]
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

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// Delete godoc
// @Summary      Delete Card
// @Description  Deletes a card by its ID.
// @Tags         cards
// @Produce      json
// @Param        card_id path      int true "Card ID"
// @Success      200  {object}  entities.Card "Deleted card"
// @failure      400  {string}  string "Bad Request"
// @Failure      500  {string}  string "Internal Server Error"
// @Router       /api/v1/cards/{card_id} [delete]
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

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
