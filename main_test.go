package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestSorting(t *testing.T) {
	products := []Product{
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
			SalesCount: 1048,                                         //sales/views = 0.32
			ViewsCount: 20123,
		},
	}

	// Test sorting by price
	expectedPriceOrder := []Product{products[2], products[0], products[1]}
	priceSorter := &PriceSorter{}
	priceSorter.Sort(products, true)
	if !reflect.DeepEqual(products, expectedPriceOrder) {
		t.Errorf("Product list not sorted by price. Expected %v, got %v", expectedPriceOrder, products)
	}

	// Test sorting by sales to view ratio
	expectedSalesToViewRatioOrder := []Product{products[1], products[0], products[2]}
	salesToViewRatioSorter := &SalesToViewRatioSorter{}
	salesToViewRatioSorter.Sort(products, true)
	if !reflect.DeepEqual(products, expectedSalesToViewRatioOrder) {
		t.Errorf("Product list not sorted by sales to view ratio. Expected %v, got %v", expectedSalesToViewRatioOrder, products)
	}
}

func BenchmarkSorters(b *testing.B) {
	products := generateProducts(10000) // Generate 10,000 random products

	b.Run("PriceSorter", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ps := &PriceSorter{}
			ps.Sort(products, true)
		}
	})

	b.Run("SalesToViewRatioSorter", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			stvr := &SalesToViewRatioSorter{}
			stvr.Sort(products, true)
		}
	})
}

// Helper function to generate random products
func generateProducts(n int) []Product {
	products := make([]Product, n)
	for i := 0; i < n; i++ {
		p := Product{
			ID:         i + 1,
			Name:       fmt.Sprintf("Product %d", i+1),
			Price:      rand.Float64() * 100,
			Created:    time.Now().AddDate(0, 0, -rand.Intn(365)),
			SalesCount: rand.Intn(1000),
			ViewsCount: rand.Intn(10000),
		}
		products[i] = p
	}
	return products
}
