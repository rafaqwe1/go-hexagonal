package application_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/rafaqwe1/go-hexagonal/application"
	"github.com/stretchr/testify/require"
)

func TestApplicationProduct_Enable(t *testing.T) {
	product := application.Product{}

	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()

	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}

func TestApplicationProduct_Disable(t *testing.T) {
	product := application.Product{}

	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()

	require.Equal(t, "the price must be zero in order to have the product disabled", err.Error())
}

func TestApplicationProduct_IsValid(t *testing.T) {
	product := application.Product{}

	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10
	product.ID = uuid.NewString()

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()

	require.Equal(t, "the status must be enabled or disabled", err.Error())

	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()

	require.Equal(t, err.Error(), "the price must be greater or equal zero")
}