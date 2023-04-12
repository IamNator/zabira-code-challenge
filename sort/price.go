package sort

import (
	"sort"

	"github.com/IamNator/zabira-code-challenge/model"
)

type PriceSorter struct{}

func (ps *PriceSorter) Sort(products []model.Product, desc bool) {
	sort.Slice(products, func(i, j int) bool {
		if desc {
			return products[i].Price > products[j].Price
		}
		return products[i].Price < products[j].Price
	})
}
