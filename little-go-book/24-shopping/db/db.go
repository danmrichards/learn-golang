package db

import (
    "learn-golang/little-go-book/24-shopping/models"
)

// Note that the Item structure is no longer in this file. If we moved it to
// the shopping package we'd get an 'import cycle not allowed' error (as the
// shopping package imports db). Instead we've moved it into a new models
// package.

// LoadItem - Loads an item and returns it.
func LoadItem(id int) *models.Item {
    return &models.Item{
        Price: 9.001,
    }
}
