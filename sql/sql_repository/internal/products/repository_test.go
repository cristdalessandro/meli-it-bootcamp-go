package products_test

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cristdalessandro/meli-it-bootcamp-go/sql/internal/domain"
	"github.com/cristdalessandro/meli-it-bootcamp-go/sql/internal/products"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	expected := []domain.Product{
		{
			Id:           1,
			Name:         "Corn Shoots",
			Quantity:     244,
			Code_value:   "0009-1111",
			Is_published: false,
			Expiration:   "2022-01-08",
			Price:        23.27,
		},
		{
			Id:           2,
			Name:         "Shrimp - Baby, Cold Water",
			Quantity:     174,
			Code_value:   "49288-0877",
			Is_published: false,
			Expiration:   "2022-08-04",
			Price:        52.12,
		},
	}

	rows := mock.NewRows(
		[]string{
			"Id",
			"Name",
			"Quantity",
			"Code_value",
			"Is_published",
			"Expiration",
			"Price",
		})
	for _, d := range expected {
		rows.AddRow(
			d.Id,
			d.Name,
			d.Quantity,
			d.Code_value,
			d.Is_published,
			d.Expiration,
			d.Price,
		)
	}

	mock.ExpectQuery(
		`SELECT id, name, quantity, code_value, is_published, expiration, price 
		FROM products`).
		WillReturnRows(rows)

	rp := products.NewRepository(db)

	products, err := rp.GetAll()

	fmt.Println(products, err)

	assert.NoError(t, err)
	assert.Equal(t, expected, products)
	assert.NoError(t, mock.ExpectationsWereMet())

}
