package main

import (
	"reflect"
	"testing"
)

func TestSorting(t *testing.T) {
	products := []Product{
		{
			ID:          1,
			Name:        "Alabaster Table",
			Price:       12.99,
			Created:     "2019-01-04",
			SalesCount:  32,
			ViewsCount:  730,
		},
		{
			ID:          2,
			Name:        "Zebra Table",
			Price:       44.49,
			Created:     "2012-01-04",
			SalesCount:  301,
			ViewsCount:  3279,
		},
		{
			ID:           3,
			Name:         "Coffee Table",
			Price:        10.00,
			Created:      "2014-05-28",
			SalesCount:   1048,
			ViewsCount:   20123,
		},
	}

	// Test sorting by price
	expectedPriceOrder := []Product{products[2], products[0], products[1]}
	priceSorter := &PriceSorter{}
	priceSorter.Sort(products)
	if !reflect.DeepEqual(products, expectedPriceOrder) {
		t.Errorf("Product list not sorted by price. Expected %v, got %v", expectedPriceOrder, products)
	}

	// Test sorting by sales to view ratio
	expectedSalesToViewRatioOrder := []Product{products[1], products[0], products[2]}
	salesToViewRatioSorter := &SalesToViewRatioSorter{}
	salesToViewRatioSorter.Sort(products)
	if !reflect.DeepEqual(products, expectedSalesToViewRatioOrder) {
		t.Errorf("Product list not sorted by sales to view ratio. Expected %v, got %v", expectedSalesToViewRatioOrder, products)
	}
}



