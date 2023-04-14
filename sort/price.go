package sort

import (
	"sort"

	"github.com/IamNator/zabira-code-challenge/model"
)

// Sorts products by price
//
// you can opt out of using the default algorithm by passing in a custom algorithm to the Alg field
type PriceSorter struct {
	Desc bool
	Alg  func([]model.Product) func(i, j int) bool
}

func (ps PriceSorter) Sort(products []model.Product) {

	if ps.Alg != nil {
		sort.Slice(products, ps.Alg(products))
		return
	}

	//
	// There is duplication here, but I think it's better to have the code duplicated 
	//
	// The alternative would be:
	//
	// sort.Slice(products, func(i, j int) bool { 
	// 	if ps.Desc { //every time the loop is executed, the compiler has to check if the condition is true
	// 		return products[i].Price > products[j].Price // previous price < current price [ascending]
	// 	}
	// 	return products[i].Price < products[j].Price // previous price < current price [ascending]
	// })
	// 
	// 
	//
	// Doing it this way, the compiler can optimize the code and only check the condition once
	//

	if ps.Desc {
		sort.Slice(products, func(i, j int) bool {
			return products[i].Price > products[j].Price // previous price > current price [descending]
		})
		return
	}

	sort.Slice(products, func(i, j int) bool {
		return products[i].Price < products[j].Price // previous price < current price [ascending]
	})

}
