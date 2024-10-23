package entities

type Status struct {
	ID    int    `json:"id" swaggerignore:"true"`
	Name  string `json:"name"`
	Color string `json:"color"`
}
