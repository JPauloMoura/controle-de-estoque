package entity

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProduct_Validate(t *testing.T) {
	tests := []struct {
		name    string
		product Product
		wantErr bool
	}{
		{
			name:    "Success",
			product: Product{Name: "name", Price: 1, Description: "description", AvailableQuantity: 1},
			wantErr: false,
		},
		{
			name:    "should return error when name is empty",
			product: Product{Name: "", Price: 1, Description: "description", AvailableQuantity: 1},
			wantErr: true,
		},
		{
			name:    "should return an error when the name has more than 50 characters",
			product: Product{Name: strings.Repeat("X", 51), Price: 1, Description: "description", AvailableQuantity: 1},
			wantErr: true,
		},
		{
			name:    "should return an error when the price is equal to zero",
			product: Product{Name: "name", Price: 0, Description: "description", AvailableQuantity: 1},
			wantErr: true,
		},
		{
			name:    "should return an error when the price is less than zero",
			product: Product{Name: "name", Price: -1, Description: "description", AvailableQuantity: 1},
			wantErr: true,
		},
		{
			name:    "should return error when description is empty",
			product: Product{Name: "name", Price: 1, Description: "", AvailableQuantity: 1},
			wantErr: true,
		},
		{
			name:    "should return an error when the description has more than 255 characters",
			product: Product{Name: "name", Price: 1, Description: strings.Repeat("X", 256), AvailableQuantity: 1},
			wantErr: true,
		},
		{
			name:    "should be error returned when the available quantity is equal to zero",
			product: Product{Name: "name", Price: 1, Description: "description", AvailableQuantity: 0},
			wantErr: true,
		},
		{
			name:    "should be error returned when the available quantity is less than zero",
			product: Product{Name: "name", Price: 1, Description: "description", AvailableQuantity: -1},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Product{
				Id:                tt.product.Id,
				Name:              tt.product.Name,
				Price:             tt.product.Price,
				Description:       tt.product.Description,
				AvailableQuantity: tt.product.AvailableQuantity,
			}

			err := p.Validate()
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
