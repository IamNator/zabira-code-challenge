package main

import (
	"fmt"
	"time"

	"github.com/IamNator/zabira-code-challenge/model"
	"github.com/IamNator/zabira-code-challenge/sort"
)

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

	priceSorter := &sort.PriceSorter{Desc: true} // create a new PriceSorter object
	sort.Sort(products, priceSorter)
	fmt.Println("\n Sorted by price: ", products) // print the sorted products

	salesToViewRatioSorter := &sort.SalesToViewRatioSorter{Desc: false} // create a new SalesToViewRatioSorter object
	sort.Sort(products, salesToViewRatioSorter)
	fmt.Println("\n Sorted by sales to view ratio:", products) // print the sorted products

	sort.Sort(products, priceSorter, salesToViewRatioSorter)                     // sort by price and then by sales to view ratio
	fmt.Println("\n Sorted by price and then by sales to view ratio:", products) // print the sorted products
}
