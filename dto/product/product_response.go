package productsdto

type ProductResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title" form:"title" validate:"required"`
	Price int    `json:"price" form:"price" gorm:"type: int" validate:"required"`
	Image string `json:"image" form:"id" validate:"required"`
	Desc  string `json:"desc" form:"desc" gorm:"type:varchar(255)"`
	Stock int    `json:"stock" form:"stock" gorm:"type:int"`
}