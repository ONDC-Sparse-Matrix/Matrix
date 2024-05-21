package models

type Merchant struct {
	Name     string             `json:"name,omnitempty" validate:"required"`
	Email    string             `json:"email,omnitempty" validate:"required"`
}
