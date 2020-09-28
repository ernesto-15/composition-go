package invoiceitem

//Item struct
type Item struct {
	id uint
	product string
	value float64
}

//New is a constructor for Item, returns an Item
func New(id uint, product string, value float64) Item {
	return Item{id, product, value,}
}

//Value is a getter
func (i Item) Value() float64 {
	return i.value
}