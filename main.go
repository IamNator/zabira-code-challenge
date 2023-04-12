package main

import (
	"fmt"
	"time"

	"github.com/IamNator/zabira-code-challenge/model"
	"github.com/IamNator/zabira-code-challenge/sort"
)

type ProductSorter interface {
	// Sort sorts the products in place
	// If desc is true, the products are sorted in descending order
	// If desc is false, the products are sorted in ascending order
	//
	// Bool is used instead of int to make the interface more readable
	// and to avoid any confusion about the meaning of the int value (e.g. 0 = ascending, 1 = descending)
	// thesame goes for strings (e.g. "asc" = ascending [ASC, aSc ...], "desc" = descending [DESC, dEsc ...]])
	//
	// The interface is defined in the sort package
	// so that the sort package can be used as a library
	// without having to import the main package
	// (which would cause a circular dependency)
	Sort(products []model.Product, desc bool)
}

// Sort sorts the products in place
// The Bool desc determines the sort order
//
// I know others may disagree, but I think it's better to have the Sort Interface take a variable for the sort order
// rather than having to create an object for each sort order
//
// This reduces the number of objects that need to be created and makes the code more readable and easier to use
// e.g.
// sorter := sort.PriceSorter{} - > sorter([]any, desc)  // user can switch between sort orders by changing the value of desc
// sorter := sort.PriceSorter{Desc: false}) - > sorter.Sort([]any) // user has to create a new object for each sort order
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

	priceSorter := &sort.PriceSorter{}            // create a new PriceSorter object
	Sort(products, priceSorter, true)             // sort the products in descending order
	fmt.Println("\n Sorted by price: ", products) // print the sorted products

	salesToViewRatioSorter := &sort.SalesToViewRatioSorter{}   // create a new SalesToViewRatioSorter object
	Sort(products, salesToViewRatioSorter, false)              // sort the products in ascending order
	fmt.Println("\n Sorted by sales to view ratio:", products) // print the sorted products
}
