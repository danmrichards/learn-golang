package shopping

import (
    "learn-golang/little-go-book/23-shopping/db"
)

// PriceCheck - Check the price for a given item by it's id.
// Returns the items price and true if it exists. Otherwise 0 and false.
func PriceCheck(itemID int) (float64, bool) {
    item := db.LoadItem(itemID)

    if item == nil {
        return 0, false
    }

    return item.Price, true
}
