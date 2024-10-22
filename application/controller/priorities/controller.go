package priorities

import (
	"encoding/json"
	"net/http"
	"strconv"

	"com.redberry.api/domain/entities"
	"com.redberry.api/domain/services/priorities"
)

// Controller handles priority-related endpoints.
type Controller struct {
	service *priorities.Service
}

// New creates a new instance of the Controller.
func New() *Controller {
	return &Controller{
		service: priorities.New(),
	}
}

// Create godoc
// @Summary      Create Priority
// @Description  Creates a new priority.
// @Tags         priorities
// @Accept       json
// @Produce      json
// @Param        body body      entities.Priority true "Priority data"
// @Success      201  {object}  entities.Priority "Created priority"
// @failure      400  {string}  string "Bad Request"
// @Failure      500  {string}  string "Internal Server Error"
// @Router       /api/v1/priorities [post]
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

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

// GetByID godoc
// @Summary      Get Priority by ID
// @Description  Retrieves a priority by its ID.
// @Tags         priorities
// @Produce      json
// @Param        priority_id path      int true "Priority ID"
// @Success      200  {object}  entities.Priority "Priority data"
// @failure      400  {string}  string "Bad Request"
// @Failure      500  {string}  string "Internal Server Error"
// @Router       /api/v1/priorities/{priority_id} [get]
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

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// Update godoc
// @Summary      Update Priority
// @Description  Updates an existing priority.
// @Tags         priorities
// @Accept       json
// @Produce      json
// @Param        priority_id path      int true "Priority ID"
// @Param        body body      entities.Priority true "Updated priority data"
// @Success      200  {object}  entities.Priority "Updated priority"
// @failure      400  {string}  string "Bad Request"
// @Failure      500  {string}  string "Internal Server Error"
// @Router       /api/v1/priorities/{priority_id} [put]
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

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// Delete godoc
// @Summary      Delete Priority
// @Description  Deletes a priority by its ID.
// @Tags         priorities
// @Produce      json
// @Param        priority_id path      int true "Priority ID"
// @Success      200  {object}  entities.Priority "Deleted priority"
// @failure      400  {string}  string "Bad Request"
// @Failure      500  {string}  string "Internal Server Error"
// @Router       /api/v1/priorities/{priority_id} [delete]
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

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
