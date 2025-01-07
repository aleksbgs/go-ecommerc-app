package domain

type Cart struct {
	ID        uint    `json:"id" gorm:"PrimaryKey"`
	UserId    uint    `json:"user_id"`
	ProductId uint    `json:"product_id"`
	Name      string  `json:"name"`
	ImageURL  string  `json:"image_url"`
	SellerId  uint    `json:"seller_id"`
	Price     float64 `json:"price"`
	Qty       uint    `json:"qty"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}
