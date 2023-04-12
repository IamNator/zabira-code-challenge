package sort

import "github.com/IamNator/zabira-code-challenge/model"

type ProductSorter interface {
	Sort(products []model.Product)
}

func Sort(products []model.Product, sorters ...ProductSorter) {
	for _, sorter := range sorters {
		sorter.Sort(products)
	}
}
