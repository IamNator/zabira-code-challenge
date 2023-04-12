package model

import (
	"fmt"
	"time"
)

type Product struct {
	ID         uint // unique identifier
	Name       string
	Price      float64
	Created    time.Time // used the time.Time type to make it easier to sort by date and to make it easier to format the date
	SalesCount int       //made it an int to make it easier to sort by sales to view ratio  but it could be a uint
	ViewsCount int       //made it an int to make it easier to sort by sales to view ratio  but it could be a uint
}

// helps for debugging
func (p Product) String() string {
	return fmt.Sprintf(`
	--------------------
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
