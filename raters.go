/*
A Rater produces a multiplier ("rate") given a value
here are two kinds of raters:
- a variable in-code one that can be used when the db is down
- one w the table stored in a file (or db, but I only know files so far)
note: I'm making them  give different rates, so that mistakes show up easier
*/
package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type InCoderRater struct{}

func (in InCoderRater) TaxRate(amount float64) float64 {
	if amount <= 100 {
		return 1.01
	}
	return 1.5
}

func NewInCoderRater() *InCoderRater {
	return &InCoderRater{}
}

type FileRater struct {
	rates [][]float64
}

func (fr FileRater) TaxRate(amount float64) float64 {
	if amount >= fr.rates[0][0] && amount < fr.rates[1][0] {
		return fr.rates[0][1]
	}
	return fr.rates[1][1]
}

func NewFileRater(fn string) *FileRater {
	rates := [][]float64{}
	file, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		rate := make([]float64, len(parts))
		for i, part := range parts {
			value, _ := strconv.ParseFloat(part, 64)
			rate[i] = value
		}
		rates = append(rates, rate)
	}

	return &FileRater{rates: rates}
}
