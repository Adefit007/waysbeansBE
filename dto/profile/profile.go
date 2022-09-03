package profiledto

import "waysbeans/models"

type CreateProfile struct {
	Address    string             `json:"address" form:"address" validate:"required"`
	Phone      string             `json:"phone" form:"phone" validate:"required"`
	Image      string             `json:"image" form:"image" validate:"required"`
	City       string             `json:"city" form:"city" validate:"required"`
	PostalCode int                `json:"postal_code" form:"postal_code" validate:"required"`
	UserID     int                `json:"user_id"`
	User       models.UserProfile `json:"user"`
}

type UpdateProfile struct {
	Address    string `json:"address" form:"address"`
	Phone      string `json:"phone" form:"phone"`
	Image      string `json:"image" form:"image"`
	City       string `json:"city" form:"city"`
	PostalCode int    `json:"postal_code" form:"postal_code"`
	UserID     int    `json:"user_id"`
}

type ProfileResponse struct {
	Address    string             `json:"address" form:"address"`
	Phone      string             `json:"phone" form:"phone"`
	Image      string             `json:"image" form:"image"`
	City       string             `json:"city" form:"city"`
	PostalCode int                `json:"postal_code" form:"postal_code"`
	UserID     int                `json:"user_id"`
	User       models.UserProfile `json:"user"`
}