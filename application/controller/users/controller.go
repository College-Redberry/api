package users

import (
	"encoding/json"
	"net/http"

	"com.redberry.api/domain/entities"
	"com.redberry.api/domain/services/users"
)

type Auth struct {
	service *users.Service
}

func New() *Auth {
	return &Auth{
		service: users.New(),
	}
}

func (controller Auth) Register(w http.ResponseWriter, r *http.Request) {
	var requestBody entities.User

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	user, err := controller.service.Register(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	response, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Write(response)
	w.WriteHeader(http.StatusOK)
}
