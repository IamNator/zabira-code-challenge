package main

import (
	"fmt"
	"sort"
)

type Product struct {
	ID          int
	Name        string
	Price       float64
	Created     string
	SalesCount  int
	ViewsCount  int
}

type ProductSorter interface {
	Sort(products []Product)
}

type PriceSorter struct{}

func (ps *PriceSorter) Sort(products []Product) {
	sort.Slice(products, func(i, j int) bool {
		return products[i].Price < products[j].Price
	})
}

type SalesToViewRatioSorter struct{}

func (svrs *SalesToViewRatioSorter) Sort(products []Product) {
	sort.Slice(products, func(i, j int) bool {
		salesToViewRatioI := float64(products[i].SalesCount) / float64(products[i].ViewsCount)
		salesToViewRatioJ := float64(products[j].SalesCount) / float64(products[j].ViewsCount)
		return salesToViewRatioI > salesToViewRatioJ
	})
}

func main() {
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

	priceSorter := &PriceSorter{}
	priceSorter.Sort(products)
	fmt.Println("Sorted by price:", products)

	salesToViewRatioSorter := &SalesToViewRatioSorter{}
	salesToViewRatioSorter.Sort(products)
	fmt.Println("Sorted by sales to view ratio:", products)
}
