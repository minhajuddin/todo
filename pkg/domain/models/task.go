package models

type Task struct {
	ID    int64  `db:"id"`
	Title string `db:"title"`
}
