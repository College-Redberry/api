package auth

import (
	"encoding/json"
	"net/http"

	"com.redberry.api/domain/entities"
	"com.redberry.api/domain/services/auth"
)

type Auth struct {
	service *auth.Service
}

func New() *Auth {
	return &Auth{
		service: auth.New(),
	}
}

// Login godoc
// @Summary      Login
// @Description  Authenticates a user and returns an access token.
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Param        body body      entities.User true "User credentials"
// @Success      200  {object}  string "Access token"
// @Failure      400  {object}  string "Bad request"
// @Failure      404  {object}  string "Not found"
// @Failure      500  {object}  string "Internal server error"
// @Router       /api/v1/auth/login [post]
func (controller Auth) Login(w http.ResponseWriter, r *http.Request) {
	var requestBody entities.User

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := controller.service.Login(requestBody.Email, requestBody.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(token))
}
