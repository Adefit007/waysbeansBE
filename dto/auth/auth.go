package authdto

type RegisterRequest struct {
	Name     string `json:"name" gorm:"type: varchar(255) required"`
	Email    string `json:"email" gorm:"type: varchar(255) required"`
	Password string `json:"password" gorm:"type: varchar(255) required"`
}

type LoginRequest struct {
	Email    string `json:"email" gorm:"type: varchar(255) required"`
	Password string `json:"password" gorm:"type: varchar(255) required"`
}

type LoginResponse struct {
	Name   string `json:"name" gorm:"type: varchar(255)"`
	Email  string `json:"email" gorm:"type: varchar(255)"`
	Token  string `json:"token" gorm:"type: varchar(255)"`
	Status string `gorm:"type: varchar(50)"  json:"status"`
}

type CheckAuthResponse struct {
	Id     int    `gorm:"type: int" json:"id"`
	Name   string `gorm:"type: varchar(255)" json:"name"`
	Email  string `gorm:"type: varchar(255)" json:"email"`
	Status string `gorm:"type: varchar(50)"  json:"status"`
}