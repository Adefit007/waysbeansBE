package models

import "time"

type Product struct {
	ID        	int       	`json:"id"`
	Title     	string    	`json:"title" gorm:"type: varchar(255)"`
	Price     	int       	`json:"price" gorm:"type: int"`
	Image     	string    	`json:"image" gorm:"type: varchar(255)"`
	Desc		string		`json:"desc" gorm:"type: text"`
	Stock		int			`json:"stock" gorm:"type: int"`
	CreatedAt 	time.Time 	`json:"-"`
	UpdatedAt 	time.Time 	`json:"-"`
}

type ProductResponse struct {
	ID		int		`json: "id"`
	Title	string	`json: "title"`
	Price	int		`json: "price"`
	Image	string	`json: "image"`
	Desc	string	`json: "desc"`
	Stock	int		`json: "stock"`
}

func (ProductResponse) TableName() string {
	return "products"
}