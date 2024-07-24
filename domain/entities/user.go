package entities

type User struct {
	ID           int32  `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password,omitempty"`
	IsAdmin      bool   `json:"is_admin"`
	ProfileImage string `json:"profile_image"`
}
