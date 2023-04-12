package sort_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/IamNator/zabira-code-challenge/model"
	"github.com/IamNator/zabira-code-challenge/sort"
)

func TestPriceSorter(t *testing.T) {
	products := []model.Product{
		{
			ID:         1,
			Name:       "Alabaster Table",
			Price:      12.99,
			Created:    time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC), //"2019-01-04"
			SalesCount: 32,                                          //sales/views = 0.04
			ViewsCount: 730,
		},
		{
			ID:         2,
			Name:       "Zebra Table",
			Price:      44.49,
			Created:    time.Date(2012, 1, 4, 0, 0, 0, 0, time.UTC), //"2012-01-04"
			SalesCount: 301,                                         //sales/views = 0.09
			ViewsCount: 3279,
		},
		{
			ID:         3,
			Name:       "Coffee Table",
			Price:      10.00,
			Created:    time.Date(2014, 5, 28, 0, 0, 0, 0, time.UTC), //"2014-05-28"
			SalesCount: 1048,                                         //sales/views = 0.05
			ViewsCount: 20123,
		},
	}

	// Test sorting by price
	expectedPriceOrder := []model.Product{products[1], products[0], products[2]}
	priceSorter := &sort.PriceSorter{}
	priceSorter.Sort(products, true) // Sort in descending order
	if !reflect.DeepEqual(products, expectedPriceOrder) {
		t.Errorf("Product list not sorted by price. \nExpected %v, \ngot %v", expectedPriceOrder, products)
	}

}

func BenchmarkPriceSorter(b *testing.B) {
	products := generateProducts(10000) // Generate 10,000 random products

	b.Run("PriceSorter", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ps := &sort.PriceSorter{}
			ps.Sort(products, true)
		}
	})
}
