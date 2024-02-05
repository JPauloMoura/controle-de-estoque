package entity

import "errors"

type Product struct {
	Id                int     `json:"id"`
	Name              string  `json:"name"`
	Price             float64 `json:"price"`
	Description       string  `json:"description"`
	AvailableQuantity int     `json:"availableQuantity"`
}

func (p Product) Validate() error {
	if p.Name == "" {
		return errors.New("product name is required")
	}

	if len(p.Name) > 50 {
		return errors.New("product name must be less than 50 characters")
	}

	if p.Price <= 0 {
		return errors.New("product price must be greater than 0")
	}

	if p.Description == "" {
		return errors.New("product description is required")
	}

	if len(p.Description) > 255 {
		return errors.New("product description must be less than 255 characters")
	}

	if p.AvailableQuantity <= 0 {
		return errors.New("product availableQuantity be greater than 0")
	}

	return nil
}
