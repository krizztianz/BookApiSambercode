package models

type Book struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	CategoryID  int     `json:"category_id"`
	Description string  `json:"description"`
	ImageURL    *string `json:"image_url"`
	ReleaseYear int     `json:"release_year"`
	Price       float64 `json:"price"`
	TotalPage   int     `json:"total_page"`
	Thickness   string  `json:"thickness"`
	CreatedAt   string  `json:"created_at"`
	CreatedBy   *string `json:"created_by"`
	ModifiedAt  *string `json:"modified_at"`
	ModifiedBy  *string `json:"modified_by"`
}

type BookRequest struct {
	Title       string  `json:"title" binding:"required"`
	CategoryID  int     `json:"category_id" binding:"required"`
	Description string  `json:"description"`
	ImageURL    string  `json:"image_url"`
	ReleaseYear int     `json:"release_year" binding:"gte=1980,lte=2024"`
	Price       float64 `json:"price"`
	TotalPage   int     `json:"total_page"`
}
