package models

type Product struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	Category_id   int     `json:"category_id"`
	Monthly_price float32 `json:"monthly_prince"`
	Users         []User  `json:"users,omitempty"`
}
