package sort

import (
	"sort"

	"github.com/IamNator/zabira-code-challenge/model"
)

// Sorts products by sales to view ratio
//
// you can opt out of using the default algorithm by passing in a custom algorithm to the Alg field
type SalesToViewRatioSorter struct {
	Desc bool
	Alg  func([]model.Product) func(i, j int) bool
}

func (svrs SalesToViewRatioSorter) Sort(products []model.Product) {

	if svrs.Alg != nil {
		sort.Slice(products, svrs.Alg(products))
		return
	}

	//
	// There is duplication here, but I think it's better to have the code duplicated
	//
	// The alternative would be:
	//
	// sort.Slice(products, func(i, j int) bool { 
	// 		salesToViewRatioI := float64(products[i].SalesCount) / float64(products[i].ViewsCount)
	// 		salesToViewRatioJ := float64(products[j].SalesCount) / float64(products[j].ViewsCount)
	//
	// 		if ps.Desc { //every time the loop is executed, the compiler has to check if the condition is true
	// 			return salesToViewRatioI > salesToViewRatioJ // previous sales to view ratio > current sales to view ratio [descending]
	// 		}
	// 		return salesToViewRatioI < salesToViewRatioJ // previous sales to view ratio < current sales to view ratio [ascending]
	// 	})
	// 
	// Doing it this way, the compiler can optimize the code and only check the condition once
	//

	if svrs.Desc {

		sort.Slice(products, func(i, j int) bool {

			salesToViewRatioI := float64(products[i].SalesCount) / float64(products[i].ViewsCount)
			salesToViewRatioJ := float64(products[j].SalesCount) / float64(products[j].ViewsCount)

			return salesToViewRatioI > salesToViewRatioJ // previous sales to view ratio > current sales to view ratio [descending]
		})

		return
	}

	sort.Slice(products, func(i, j int) bool {

		salesToViewRatioI := float64(products[i].SalesCount) / float64(products[i].ViewsCount)
		salesToViewRatioJ := float64(products[j].SalesCount) / float64(products[j].ViewsCount)

		return salesToViewRatioI < salesToViewRatioJ // previous sales to view ratio < current sales to view ratio [ascending]
	})
}
