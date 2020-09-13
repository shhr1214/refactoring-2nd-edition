package main

import (
	"testing"
)

func TestCustomer_Statement(t *testing.T) {
	customer := NewCustomer("hara")
	movie1 := NewMovie("通常", REGULAR)
	rental1 := NewRental(movie1, 3)
	customer.AddRental(rental1)

	movie2 := NewMovie("新作", NEW_RELEASE)
	rental2 := NewRental(movie2, 5)
	customer.AddRental(rental2)

	expected := `Rental Record for hara
	通常	3.5
	新作	15.0
Amount owed is 18.5
You earned 3 frequent renter points
`
	if customer.Statement() != expected {
		t.Errorf("Test Fail\nexpect: %s\nactual: %s\n", expected, customer.Statement())
	}
}
