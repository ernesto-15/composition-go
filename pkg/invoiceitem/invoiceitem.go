package invoiceitem

//Item struct
type Item struct {
	id      uint
	product string
	value   float64
}

//New is a constructor for Item, returns an Item
func New(id uint, product string, value float64) Item {
	return Item{id, product, value}
}

//Value is a getter
// func (i Item) Value() float64 {
// 	return i.value
// }

//Items is a custom type created from the Item struct
type Items []Item

//NewItems is a constructor for Items type
func NewItems(items ...Item) Items{
	var is Items
	for _, v := range items {
		is = append(is, v)
	}
	return is
}

//Total is a method for Items type
func (is Items) Total() float64 {
	var total float64
	for _, v := range is {
		total += v.value
	}
	return total
}
