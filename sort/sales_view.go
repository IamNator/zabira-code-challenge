package sort

import (
	"sort"

	"github.com/IamNator/zabira-code-challenge/model"
)

type SalesToViewRatioSorter struct{} // Sorts products by sales to view ratio

// Sort sorts the products by sales to view ratio
// If desc is true, the products are sorted in descending order
// If desc is false, the products are sorted in ascending order
//
// Bool is used instead of int to make the interface more readable
// and to avoid any confusion about the meaning of the int or string value (e.g. 0 = ascending, 1 = descending) or ("asc" = ascending [ASC, aSc ...], "desc" = descending [DESC, dEsc ...]])
func (svrs *SalesToViewRatioSorter) Sort(products []model.Product, desc bool) {

	//
	// There is duplication here, but I think it's better to have the code duplicated
	//
	// My reasoning is that every time the loop is executed, the compiler has to check if the condition is true
	//
	// Doing it this way, the compiler can optimize the code and only check the condition once
	//

	if desc {

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
