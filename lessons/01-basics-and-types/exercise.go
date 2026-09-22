package main

// Lesson 1 exercise
// Fill in each TODO. Run `go vet ./...` from the project root to check your code,
// then `go run ./lessons/01-basics-and-types/` to run it.
//
// Remember: unused imports and unused local variables are compile errors in Go —
// add an import to the block below only once you actually use it (you'll need
// "fmt" almost immediately, and "strings" once you get to TODO 5).

import (
)

// TODO 4: write a function `describeCity(city string, population int) string`
// that returns a string like "Chicago has 2700000 people." using fmt.Sprintf.
// Declare it here, at package level (outside main).

// TODO 5: write a function `summarize(landmarks []string) string` that joins
// them with ", " (use strings.Join).

// TODO 6: write a function `divide(a, b float64) (float64, error)` that
// returns an error via fmt.Errorf if b is 0, otherwise returns a/b and a nil
// error.

// TODO 7: write a function `safeLength(value any) int` that returns the
// string's length if value is a string (using the two-value type assertion
// `s, ok := value.(string)`), otherwise 0.

func main() {
	// TODO 1: declare a variable `city` (string) and `population` (int) using
	// `:=`, set to any city name and population.

	// TODO 2: declare a variable `landmarks` ([]string) with at least two
	// landmark names in `city`.

	// TODO 3: call `describeCity` and `summarize` and print both results with
	// fmt.Println.

	// TODO 6 (continued): call `divide` once with a non-zero divisor and once
	// with 0. For each call, check `if err != nil` and print either the error
	// or the result.

	// TODO 7 (continued): call `safeLength` with a string value and with a
	// non-string value (e.g. an int), printing both results.
}
