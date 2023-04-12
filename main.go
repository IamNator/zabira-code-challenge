package main

import (
	"fmt"
	"sort"
	"time"
)

type Product struct {
	ID         int
	Name       string
	Price      float64
	Created    time.Time
	SalesCount int
	ViewsCount int
}

func (p Product) String() string {
	return fmt.Sprintf(`
	ID: %d, 
	Name: %s, 
	Price: %.2f, 
	Created: %s, 
	SalesCount: %d, 
	ViewsCount: %d 

	metadata: {
		sales/views: %.2f
	}
	`,
		p.ID, p.Name, p.Price, p.Created.Format("2006-01-02"), p.SalesCount, p.ViewsCount, float64(p.SalesCount)/float64(p.ViewsCount))
}

type ProductSorter interface {
	Sort(products []Product, desc bool)
}

type PriceSorter struct{}

func (ps *PriceSorter) Sort(products []Product, desc bool) {
	sort.Slice(products, func(i, j int) bool {
		if desc {
			return products[i].Price > products[j].Price
		}
		return products[i].Price < products[j].Price
	})
}

type SalesToViewRatioSorter struct{}

func (svrs *SalesToViewRatioSorter) Sort(products []Product, desc bool) {
	sort.Slice(products, func(i, j int) bool {

		salesToViewRatioI := float64(products[i].SalesCount) / float64(products[i].ViewsCount)
		salesToViewRatioJ := float64(products[j].SalesCount) / float64(products[j].ViewsCount)

		if desc {
			return salesToViewRatioI > salesToViewRatioJ
		}
		return salesToViewRatioI < salesToViewRatioJ
	})
}

func main() {
	products := []Product{
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

	priceSorter := &PriceSorter{}
	priceSorter.Sort(products, true)

	fmt.Println("\n Sorted by price: ", products)

	salesToViewRatioSorter := &SalesToViewRatioSorter{}
	salesToViewRatioSorter.Sort(products, true)

	fmt.Println("\n Sorted by sales to view ratio:", products)
}
