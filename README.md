# Smaller Web Hexagon

Based on Alistair Cockburn [example repo](https://github.com/totheralistair/SmallerWebHexagon/).

Illustration of a simple hexagon with one user (left) port and one database (right) port.

It simply calculates the formula:

    result = input * rate

e.g.

    result = 100 * 1.1 = 110

The user port is connected to either the web or to a test harness, with or without the server adapter.

The database port looks up the rate in a database, either in-the-code database, or from a file.

The startup decides how the wiring goes.

## Running the Application

Run `go run .` to get the web UI on port 9292

## Running Tests

Run `go test` to run tests.
