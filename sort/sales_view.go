package sort

import (
	"sort"

	"github.com/IamNator/zabira-code-challenge/model"
)

type SalesToViewRatioSorter struct{}

func (svrs *SalesToViewRatioSorter) Sort(products []model.Product, desc bool) {
	sort.Slice(products, func(i, j int) bool {

		salesToViewRatioI := float64(products[i].SalesCount) / float64(products[i].ViewsCount)
		salesToViewRatioJ := float64(products[j].SalesCount) / float64(products[j].ViewsCount)

		if desc {
			return salesToViewRatioI > salesToViewRatioJ
		}
		return salesToViewRatioI < salesToViewRatioJ
	})
}
