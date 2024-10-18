package users

import (
	"encoding/json"
	"net/http"
	"strconv"

	"com.redberry.api/domain/entities"
	"com.redberry.api/domain/services/users"
)

// Controller handles user-related endpoints.
type Controller struct {
	service *users.Service
}

// New creates a new instance of the Controller.
func New() *Controller {
	return &Controller{
		service: users.New(),
	}
}

// Register godoc
// @Summary      Register User
// @Description  Registers a new user.
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        body body      entities.User true "User data"
// @Success      201  {object}  entities.User "Registered user"
// @failure      400  {string}  string "Bad Request"
// @Failure      500  {string}  string "Internal Server Error"
// @Router       /api/v1/users [post]
func (controller Controller) Register(w http.ResponseWriter, r *http.Request) {
	var requestBody entities.User

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := controller.service.Register(requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

// GetByEmail godoc
// @Summary      Get User by Email
// @Description  Retrieves a user by their email.
// @Tags         users
// @Produce      json
// @Param        email query     string true "User email"
// @Success      200  {object}  entities.User "User data"
// @Failure      500  {string}  string "Internal Server Error"
// @Router       /api/v1/users [get]
func (controller Controller) GetByEmail(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")

	user, err := controller.service.GetByEmail(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// Update godoc
// @Summary      Update User
// @Description  Updates an existing user.
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user_id path      int true "User ID"
// @Param        body body      entities.User true "Updated user data"
// @Success      200  {object}  entities.User "Updated user"
// @failure      400  {string}  string "Bad Request"
// @Failure      500  {string}  string "Internal Server Error"
// @Router       /api/v1/users/{user_id} [put]
func (controller Controller) Update(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.PathValue("user_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var requestBody entities.User

	err = json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := controller.service.Update(userID, requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// Delete godoc
// @Summary      Delete User
// @Description  Deletes a user by their ID.
// @Tags         users
// @Produce      json
// @Param        user_id path      int true "User ID"
// @Success      200  {object}  entities.User "Deleted user"
// @failure      400  {string}  string "Bad Request"
// @Failure      500  {string}  string "Internal Server Error"
// @Router       /api/v1/users/{user_id} [delete]
func (controller Controller) Delete(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.PathValue("user_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := controller.service.Delete(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
