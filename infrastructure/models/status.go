package models

type Status struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Color string `db:"color"`
}
