package models

import "time"

type Cart struct {
	ID        int       `json:"id" gorm:"primary_key:auto_increment"`
	QTY       int       `json:"qty"`
	SubTotal  int       `json:"subtotal"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}