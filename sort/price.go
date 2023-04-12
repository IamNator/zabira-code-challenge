package sort

import (
	"sort"

	"github.com/IamNator/zabira-code-challenge/model"
)

type PriceSorter struct{} // Sorts products by price

// Sort sorts the products by price
// If desc is true, the products are sorted in descending order
// If desc is false, the products are sorted in ascending order
//
// Bool is used instead of int to make the interface more readable
// and to avoid any confusion about the meaning of the int or string value (e.g. 0 = ascending, 1 = descending) or ("asc" = ascending [ASC, aSc ...], "desc" = descending [DESC, dEsc ...]])
func (ps *PriceSorter) Sort(products []model.Product, desc bool) {
	sort.Slice(products, func(i, j int) bool {
		if desc {
			return products[i].Price > products[j].Price
		}
		return products[i].Price < products[j].Price
	})
}
