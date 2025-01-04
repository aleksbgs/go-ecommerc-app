package domain

type Product struct {
	ID          uint64  `json:"id"`
	Name        string  `json:"name" gorm:"index;"`
	Description string  `json:"description"`
	CategoryId  uint64  `json:"category_id"`
	ImageURL    string  `json:"image_url"`
	Price       float64 `json:"price"`
	UserID      int     `json:"user_id"`
	Stock       uint    `json:"stock"`
	CreatedAt   int64   `json:"created_at"`
	UpdatedAt   int64   `json:"updated_at"`
}
