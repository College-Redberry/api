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

// Create godoc
// @Summary      Create Board
// @Description  Creates a new board.
// @Tags         boards
// @Accept       json
// @Produce      json
// @Param        body body      entities.Board true "Board data"
// @Success      201  {object}  entities.Board "Created board"
// @failure      400  {string}  string "Bad Request"
// @Failure      500  {string}  string "Internal Server Error"
// @Router       /api/v1/boards [post]
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

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

// GetByID godoc
// @Summary      Get Board by ID
// @Description  Retrieves a board by its ID.
// @Tags         boards
// @Produce      json
// @Param        board_id path      int true "Board ID"
// @Success      200  {object}  entities.Board "Board data"
// @failure      400  {string}  string "Bad Request"
// @Failure      500  {string}  string "Internal Server Error"
// @Router       /api/v1/boards/{board_id} [get]
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

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// Update godoc
// @Summary      Update Board
// @Description  Updates an existing board.
// @Tags         boards
// @Accept       json
// @Produce      json
// @Param        board_id path      int true "Board ID"
// @Param        body body      entities.Board true "Updated board data"
// @Success      200  {object}  entities.Board "Updated board"
// @failure      400  {string}  string "Bad Request"
// @Failure      500  {string}  string "Internal Server Error"
// @Router       /api/v1/boards/{board_id} [put]
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

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// Delete godoc
// @Summary      Delete Board
// @Description  Deletes a board by its ID.
// @Tags         boards
// @Produce      json
// @Param        board_id path      int true "Board ID"
// @Success      200  {object}  entities.Board "Deleted board"
// @failure      400  {string}  string "Bad Request"
// @Failure      500  {string}  string "Internal Server Error"
// @Router       /api/v1/boards/{board_id} [delete]
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

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
