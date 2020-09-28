package main

import (
	"fmt"

	"github.com/ernesto-15/composition-go/pkg/customer"
	"github.com/ernesto-15/composition-go/pkg/invoice"
	"github.com/ernesto-15/composition-go/pkg/invoiceitem"
)

func main() {
	i := invoice.New(
		"Colombia",
		"Bogota",
		customer.New("Ernesto", "Av Nose 521", "958642255"),
		invoiceitem.NewItems(
			invoiceitem.New(2, "Go OOP", 52.34),
			invoiceitem.New(3, "Go DB", 90.34),
			invoiceitem.New(1, "Go Course", 12.34),
		),
	)

	i.SetTotal()

	fmt.Printf("%+v", i)
}
