package models

import "time"

type Cart struct {
	ID        		int       		`json:"id" gorm:"primary_key:auto_increment"`
	ProductID		int				`json:"product_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Product			ProductResponse	`json:"product"`
	TransactionID	int				`json:"transaction_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Transaction		Transaction		`json:"-"`
	QTY       		int       		`json:"qty"`
	SubTotal  		int       		`json:"subtotal"`
	CreatedAt 		time.Time 		`json:"-"`
	UpdatedAt 		time.Time 		`json:"-"`
}

type CartResponse struct {
	ID        	int             `json:"id"`
	ProductID 	int             `json:"product_id"`
	Product   	ProductResponse `json:"product"`
	Qty       	int             `json:"qty"`
	SubTotal	int             `json:"subtotal"`
}

func (CartResponse) TableName() string {
	return "carts"
}