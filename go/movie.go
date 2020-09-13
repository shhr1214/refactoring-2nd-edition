package main

import "fmt"

const CHILDREN = 2
const REGULAR = 0
const NEW_RELEASE = 1

type Movie struct {
	title     string
	priceCode int
}

func NewMovie(title string, priceCode int) Movie {
	return Movie{
		title:     title,
		priceCode: priceCode,
	}
}

func (m Movie) GetPriceCode() int {
	return m.priceCode
}

func (m *Movie) SetPriceCode(priceCode int) {
	m.priceCode = priceCode
}

func (m Movie) GetTitle() string {
	return m.title
}

type Rental struct {
	movie      Movie
	daysRented int
}

func NewRental(movie Movie, daysRented int) Rental {
	return Rental{
		movie:      movie,
		daysRented: daysRented,
	}
}

func (r Rental) GetDaysRented() int {
	return r.daysRented
}

func (r Rental) GetMovie() Movie {
	return r.movie
}

type Customer struct {
	name    string
	rentals []Rental
}

func NewCustomer(name string) Customer {
	return Customer{
		name: name,
	}
}

func (c *Customer) AddRental(r Rental) {
	c.rentals = append(c.rentals, r)
}

func (c Customer) GetName() string {
	return c.name
}

func (c Customer) Statement() string {
	totalAmount := 0.0
	frequentRenterPoints := 0
	rentals := c.rentals
	result := fmt.Sprintf("Rental Record for %s\n", c.GetName())

	for _, v := range rentals {
		thisAmount := v.amountFor()

		frequentRenterPoints++
		if v.GetMovie().GetPriceCode() == NEW_RELEASE &&
			v.GetDaysRented() > 1 {
			frequentRenterPoints++
		}

		result = fmt.Sprintf("%s\t%s\t%.1f\n", result, v.GetMovie().title, thisAmount)
		totalAmount += thisAmount
	}

	result = fmt.Sprintf("%sAmount owed is %.1f\n", result, totalAmount)
	result = fmt.Sprintf("%sYou earned %d frequent renter points\n", result, frequentRenterPoints)
	return result
}

func (r Rental) amountFor() float64 {
	var result float64
	switch r.GetMovie().GetPriceCode() {
	case REGULAR:
		result += 2.0
		if r.GetDaysRented() > 2 {
			result += float64(r.GetDaysRented()-2) * 1.5
		}
	case NEW_RELEASE:
		result += float64(r.GetDaysRented() * 3)
	case CHILDREN:
		result += 1.5
		if r.GetDaysRented() > 3 {
			result += float64(r.GetDaysRented()-3) * 1.5
		}
	}
	return result
}
