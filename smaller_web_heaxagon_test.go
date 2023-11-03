package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func valueShouldProduceRate(t *testing.T, app *SmallerWebHexagon, value float64, expRate float64) {
	t.Helper()
	rate, result := app.RateAndResult(value)
	if rate != expRate {
		t.Errorf("value %f got rate %f, want %f", value, rate, expRate)
	}
	if result != value*expRate {
		t.Errorf("value %f got result %f, want %f", value, rate, expRate)
	}
}

func TestItWorksWithInCodeRater(t *testing.T) {
	app := NewSmallerWebHexagon(NewInCoderRater())
	valueShouldProduceRate(t, app, 100, 1.01)
	valueShouldProduceRate(t, app, 200, 1.5)
}

func TestItWorksWithFileRater(t *testing.T) {
	app := NewSmallerWebHexagon(NewFileRater("testdata/file_rater.txt"))
	valueShouldProduceRate(t, app, 10, 1.00)
	valueShouldProduceRate(t, app, 100, 2)
}

func TestRunsViaHttpAdapter(t *testing.T) {
	hex := NewSmallerWebHexagon(NewInCoderRater())
	viewsFolder := "views/"
	app := NewHTTPAdapter(hex, viewsFolder)
	req := httptest.NewRequest(http.MethodGet, "/100", strings.NewReader(""))
	res := httptest.NewRecorder()
	app.ServeHTTP(res, req)

	out := map[string]float64{
		"value":  100,
		"rate":   1.01,
		"result": 100 * 1.01,
	}
	expBody := string(HTMLFromTemplateFile("views/result.tmpl", out))
	got := res.Body.String()
	if got != expBody {
		t.Errorf("expect %q, want %q", got, expBody)
	}
}
