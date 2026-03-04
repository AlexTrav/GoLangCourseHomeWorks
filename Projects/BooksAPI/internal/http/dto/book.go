package dto

type CreateBookRequest struct {
	Title  string `json:"title" example:"Clean Code"`
	Author string `json:"author" example:"Robert C. Martin"`
	ISBN   string `json:"isbn" example:"9780132350884"`
}

type UpdateBookRequest struct {
	Title  string `json:"title" example:"Clean Code 2"`
	Author string `json:"author" example:"Robert C. Martin"`
	ISBN   string `json:"isbn" example:"9780132350884"`
}

type BookResponse struct {
	ID     int    `json:"id" example:"1"`
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}
