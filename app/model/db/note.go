package db_model

type Note struct {
	Title       string `db:"title"`
	Description string `db:"description"`
}
