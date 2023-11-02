package main

func main() {
	hex := NewSmallerWebHexagon(NewInCoderRater())
	app := NewHttpAdapter(hex, "views")
	app.Serve()
}
