package statuses

import (
	"encoding/json"
	"net/http"
	"strconv"

	"com.redberry.api/domain/entities"
	"com.redberry.api/domain/services/statuses"
)

// Controller handles status-related endpoints.
type Controller struct {
	service *statuses.Service
}

// New creates a new instance of the Controller.
func New() *Controller {
	return &Controller{
		service: statuses.New(),
	}
}

// Create godoc
// @Summary      Create Status
// @Description  Creates a new status.
// @Security BearerAuth
// @Tags         statuses
// @Accept       json
// @Produce      json
// @Param        body body      entities.Status true "Status data"
// @Success      201  {object}  entities.Status "Created status"
// @failure      400  {string}  string "Bad Request"
// @Failure      500  {string}  string "Internal Server Error"
// @Router       /api/v1/statuses [post]
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

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

// GetByID godoc
// @Summary      Get Status by ID
// @Description  Retrieves a status by its ID.
// @Security BearerAuth
// @Tags         statuses
// @Produce      json
// @Param        status_id path      int true "Status ID"
// @Success      200  {object}  entities.Status "Status data"
// @failure      400  {string}  string "Bad Request"
// @Failure      500  {string}  string "Internal Server Error"
// @Router       /api/v1/statuses/{status_id} [get]
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

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// Update godoc
// @Summary      Update Status
// @Description  Updates an existing status.
// @Security BearerAuth
// @Tags         statuses
// @Accept       json
// @Produce      json
// @Param        status_id path      int true "Status ID"
// @Param        body body      entities.Status true "Updated status data"
// @Success      200  {object}  entities.Status "Updated status"
// @failure      400  {string}  string "Bad Request"
// @Failure      500  {string}  string "Internal Server Error"
// @Router       /api/v1/statuses/{status_id} [put]
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

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// Delete godoc
// @Summary      Delete Status
// @Description  Deletes a status by its ID.
// @Security BearerAuth
// @Tags         statuses
// @Produce      json
// @Param        status_id path      int true "Status ID"
// @Success      200  {object}  entities.Status "Deleted status"
// @failure      400  {string}  string "Bad Request"
// @Failure      500  {string}  string "Internal Server Error"
// @Router       /api/v1/statuses/{status_id} [delete]
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

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
