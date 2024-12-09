package dto_model

type GetNotesResponse struct {
	Notes []Note `json:"notes"`
}

type Note struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
