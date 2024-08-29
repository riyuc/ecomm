package storer

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

func TestCreateProduct(t *testing.T) {
	mockDB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer mockDB.Close()
	db := sqlx.NewDb(mockDB, "sqlmock")
	ms := NewMySQLStorer(db)

	p := &Product{
		Name:        "test",
		Image:       "test",
		Category:    "test",
		Description: "test",
		Rating:      5,
		NumReviews:  10,
		Price:       100,
		CountInStock: 10,
	}

	mock.ExpectExec("INSERT INTO products(name, image, category, description, rating, num_reviews, price, count_in_stock) VALUES(?, ?, ?, ?, ?, ?, ?, ?)").WillReturnResult(sqlmock.NewResult(1, 1))
}