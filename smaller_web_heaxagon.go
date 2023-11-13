/*
Welcome to Smallerwebhexagon, the simple web hexagon implementation
Alistair Cockburn and a couple of really nice friends

this hexagon has one primary port and one secondary
the user/test/web side is the primary
the rates "database" mechanism is the secondary
the app itself just returns value * rate(as a function of value)
*/
package main

type Rater interface {
	Rate(value float64) float64
}

type Rateable interface {
	RateAndResult(value float64) (float64, float64)
}

type SmallerWebHexagon struct {
	rater Rater
}

func (app SmallerWebHexagon) RateAndResult(value float64) (float64, float64) {
	rate := app.rater.Rate(value)
	result := value * rate
	return rate, result
}

func NewSmallerWebHexagon(rater Rater) *SmallerWebHexagon {
	return &SmallerWebHexagon{rater: rater}
}
