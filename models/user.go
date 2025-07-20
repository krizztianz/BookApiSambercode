package models

type User struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"-"`
	CreatedAt  string `json:"created_at"`
	CreatedBy  string `json:"created_by"`
	ModifiedAt string `json:"modified_at"`
	ModifiedBy string `json:"modified_by"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
