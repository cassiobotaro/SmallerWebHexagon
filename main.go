package main

func main() {
	hex := NewSmallerWebHexagon(NewInCoderRater())
	app := NewHTTPAdapter(hex, "views")
	app.Serve()
}
