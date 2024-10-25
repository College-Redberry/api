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
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        body body      entities.LoginRequest true "User credentials"  // Atualização aqui
// @Success      200  {object}  string "Access token"
// @failure      400  {string}  string "Bad Request"
// @Failure      404  {object}  string "Not found"
// @Failure      500  {string}  string "Internal Server Error"
// @Router       /api/v1/auth/login [post]
func (controller Auth) Login(w http.ResponseWriter, r *http.Request) {
	var requestBody entities.LoginRequest  // Atualização aqui

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
