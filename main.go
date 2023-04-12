package main

import (
	"fmt"
	"time"

	"github.com/IamNator/zabira-code-challenge/model"
	"github.com/IamNator/zabira-code-challenge/sort"
)

type ProductSorter interface {
	Sort(products []model.Product, desc bool)
}

func Sort(products []model.Product, sorter ProductSorter, desc bool) {
	sorter.Sort(products, desc)
}

func main() {
	products := []model.Product{
		{
			ID:         1,
			Name:       "Alabaster Table",
			Price:      12.99,
			Created:    time.Date(2019, 1, 4, 0, 0, 0, 0, time.UTC), //"2019-01-04"
			SalesCount: 32,
			ViewsCount: 730,
		},
		{
			ID:         2,
			Name:       "Zebra Table",
			Price:      44.49,
			Created:    time.Date(2012, 1, 4, 0, 0, 0, 0, time.UTC), //"2012-01-04"
			SalesCount: 301,
			ViewsCount: 3279,
		},
		{
			ID:         3,
			Name:       "Coffee Table",
			Price:      10.00,
			Created:    time.Date(2014, 5, 28, 0, 0, 0, 0, time.UTC), //"2014-05-28"
			SalesCount: 1048,
			ViewsCount: 20123,
		},
	}

	priceSorter := &sort.PriceSorter{}
	Sort(products, priceSorter, true)
	fmt.Println("\n Sorted by price: ", products)

	salesToViewRatioSorter := &sort.SalesToViewRatioSorter{}
	Sort(products, salesToViewRatioSorter, false)
	fmt.Println("\n Sorted by sales to view ratio:", products)
}
