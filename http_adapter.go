package main

import (
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	addr = ":9292"
)

type HTTPAdapter struct {
	app         Rateable
	viewsFolder string
}

func (adp HTTPAdapter) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	value := pathAsNumber(req)

	rate, result := adp.app.RateAndResult(value)

	out := map[string]float64{
		"value":  value,
		"rate":   rate,
		"result": result,
	}
	const resultTmpl = "result.tmpl"
	page := HTMLFromTemplateFile(filepath.Join(adp.viewsFolder, resultTmpl), out)
	res.Write(page)
}

func (adp HTTPAdapter) Serve() {
	if err := http.ListenAndServe(addr, adp); err != nil {
		log.Fatalln(err)
	}
}

func NewHTTPAdapter(app Rateable, viewsFolder string) *HTTPAdapter {
	return &HTTPAdapter{app, viewsFolder}
}

func pathAsNumber(request *http.Request) float64 {
	return numberOrZero(pathContent(request))
}

func pathContent(request *http.Request) string {
	const pathSeparator = "/"
	return strings.TrimPrefix(request.URL.Path, pathSeparator)
}

func numberOrZero(s string) float64 {
	number, err := strconv.ParseFloat(s, 0)
	if err != nil {
		return 0
	}
	return number
}
