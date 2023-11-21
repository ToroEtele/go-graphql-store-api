package model

type User struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"is_admin"`
}

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Image       string  `json:"image"`
	Description string  `json:"description"`
}

type OrderItem struct {
	ID       string   `json:"id"`
	Quantity int      `json:"quantity"`
	Product  *Product `json:"product"`
}

type Order struct {
	ID         string       `json:"id"`
	Name       string       `json:"name"`
	Email      string       `json:"email"`
	Country    string       `json:"country"`
	Address    string       `json:"address"`
	CreatedAt  string       `json:"created_at"`
	UpdatedAt  string       `json:"updated_at"`
	IsPaid     bool         `json:"is_paid"`
	User       *User        `json:"user"`
	OrderItems []*OrderItem `json:"order_items"`
}
