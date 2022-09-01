package productsdto

type ProductRequest struct {
	Title string `json:"title" form:"title" gorm:"type:varchar(255)" validate:"required"`
	Price int    `json:"price" gorm:"type: int" form:"price" validate:"required"`
	Image string `json:"image" form:"image" gorm:"type:varchar(255)"`
	Desc  string `json:"desc" form:"desc" gorm:"text"`
	Stock int    `json:"stock" form:"stock" gorm:"type:int"`
}

type UpdateProduct struct {
	Title string `json:"title" form:"title"`
	Price int    `json:"price" gorm:"type: int" form:"price"`
	Image string `json:"image" form:"image"`
	Desc  string `json:"desc" form:"desc" gorm:"text"`
	Stock int    `json:"stock" form:"stock" gorm:"type:int"`
}
