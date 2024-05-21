package models

import ()

type Map struct {
	PIN_CODE string   `json:"pin_code,omitempty" validate:"required"`
	MERCHANT_IDS    []string `json:"merchant_ids,omitempty" bson:"merchant_ids,omitempty"`
}
