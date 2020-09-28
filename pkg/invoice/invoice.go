package invoice

import (
	"github.com/ernesto-15/composition-go/pkg/customer"
	"github.com/ernesto-15/composition-go/pkg/invoiceitem"
)

//Invoice struct for a invoice
type Invoice struct {
	country string
	city string
	total float64
	customer customer.Customer
	items []invoiceitem.Item
}

//New is a constructor for Invoice, returns a Invoice
func New(country, city string, customer customer.Customer, items []invoiceitem.Item) *Invoice {
	return &Invoice {
		country: country,
		city: city,
		customer: customer,
		items: items,
	}
}

//SetTotal setter de Invoice.total
func (i *Invoice) SetTotal() {
	for _, v := range i.items {
		i.total += v.Value()
	}
}