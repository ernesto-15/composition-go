package customer

//Customer struct
type Customer struct {
	name string
	addres string
	phone string
}

//New is a constructor for Customer, returns a Customer
func New(name, address, phone string) Customer {
	return Customer{name, address, phone,}
}