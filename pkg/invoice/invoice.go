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
	items invoiceitem.Items
}

//New is a constructor for Invoice, returns a Invoice
func New(country, city string, customer customer.Customer, items invoiceitem.Items) *Invoice {
	return &Invoice {
		country: country,
		city: city,
		customer: customer,
		items: items,
	}
}

//SetTotal setter de Invoice.total
func (i *Invoice) SetTotal() {
	i.total = i.items.Total()
}