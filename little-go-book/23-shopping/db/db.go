package db

// Item - A shopping item.
type Item struct {
	Price float64
}

// LoadItem - Loads an item and returns it.
func LoadItem(id int) *Item {
	return &Item{
		Price: 9.001,
	}
}
