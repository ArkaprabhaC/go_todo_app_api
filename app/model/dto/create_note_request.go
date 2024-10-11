package dto

type CreateNoteRequest struct {
	Title            string		 `json:"title"` 	
	Description      string      `json:"description"`
}
