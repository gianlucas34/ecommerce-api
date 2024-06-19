package domain

import (
	"testing"
	"time"

	"github.com/gianlucas34/ecommerce-api/internal/errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	t.Run("Should create product entity correctly", func(t *testing.T) {
		product, _ := NewProduct("Produto", 10)
		_, err := uuid.Parse(product.ID)

		if err != nil {
			t.Errorf(err.Error())
		}

		assert.Equal(t, "Produto", product.Name)
		assert.Equal(t, 10.00, product.Price)

		_, err = time.Parse(time.RFC3339, product.CreatedAt)

		if err != nil {
			t.Errorf("A data gerada para o campo CreatedAt é inválida, forneça uma data do tipo RFC3339!")
		}

		_, err = time.Parse(time.RFC3339, product.UpdatedAt)

		if err != nil {
			t.Errorf("A data gerada para o campo UpdatedAt é inválida, forneça uma data do tipo RFC3339!")
		}
	})

	t.Run("Should return EMPTY_PRODUCT_NAME exception if name is empty", func(t *testing.T) {
		_, err := NewProduct("", 10)

		assert.EqualError(t, err, errors.EMPTY_PRODUCT_NAME)
	})

	t.Run("Should return INVALID_PRODUCT_PRICE exception if price is invalid", func(t *testing.T) {
		_, err := NewProduct("Produto", 0)

		assert.EqualError(t, err, errors.INVALID_PRODUCT_PRICE)
	})
}
