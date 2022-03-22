package model

// Product struct
type Product struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
}
type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
