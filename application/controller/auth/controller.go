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
// @Description  get user access token
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object}  string
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
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

	w.Write([]byte(token))
	w.WriteHeader(http.StatusOK)
}
