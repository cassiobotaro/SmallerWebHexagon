/*
Welcome to Smallerwebhexagon, the simple web hexagon implementation
Alistair Cockburn and a couple of really nice friends

this hexagon has one primary port and one secondary
the user/test/web side is the primary
the rates "database" mechanism is the secondary
the app itself just returns value * rate(as a function of value)
*/
package main

type ForGettingTaxRates interface {
	TaxRate(amount float64) float64
}

type ForCalculatingTaxes interface {
	TaxOn(amount float64) (float64, float64)
}

type SmallerWebHexagon struct {
	taxRateRepository ForGettingTaxRates
}

func (app SmallerWebHexagon) TaxOn(amount float64) (float64, float64) {
	rate := app.taxRateRepository.TaxRate(amount)
	result := amount * rate
	return rate, result
}

func NewSmallerWebHexagon(taxRateRepository ForGettingTaxRates) *SmallerWebHexagon {
	return &SmallerWebHexagon{taxRateRepository: taxRateRepository}
}
