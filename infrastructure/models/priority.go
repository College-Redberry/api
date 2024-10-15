package models

type Priority struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Color string `db:"color"`
}
