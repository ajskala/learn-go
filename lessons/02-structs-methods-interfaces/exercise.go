package main

// Lesson 2 exercise
// Fill in each TODO. Run `go vet ./...` from the project root to check your code,
// then `go run ./lessons/02-structs-methods-interfaces/` to run it.
//
// Remember: unused imports and unused local variables are compile errors in Go —
// add to the import block below only once you actually use it.

import (
)

// TODO 1: define a struct `Book` with exported fields:
//   Title  string
//   Author string
//   Year   int

// TODO 2: write a method `Describe` on `Book` (value receiver) returning a
// string like "Dune by Frank Herbert (1965)".

// TODO 3: define a struct `EBook` that embeds `Book` (no field name) and adds:
//   FileSizeMB float64

// TODO 6: define an interface `Shape` with one method: `Area() float64`.

// TODO 7: define a struct `Square` with field `Side float64`, and a method
// `Area() float64` on it (value receiver) returning Side * Side. Do NOT write
// anything mentioning `Shape` in Square's definition — it should satisfy the
// interface purely by having a matching method.

// TODO 8: define a struct `Counter` with an unexported field `count int`, and
// two methods on it: `Increment()` (POINTER receiver — it needs to mutate) and
// `Value() int` (value receiver is fine — it only reads).

func main() {
	// TODO 4: create a Book, call .Describe() on it, print the result.

	// TODO 5: create an EBook, print its promoted .Title field and call its
	// promoted .Describe() method — both should work without referencing
	// EBook.Book explicitly.

	// TODO 7 (continued): declare a variable of type Shape, assign a Square to
	// it, and print the result of calling .Area() on it through the Shape
	// variable.

	// TODO 8 (continued): create a Counter (its zero value is already usable —
	// count starts at 0), call Increment() three times, then print Value().
	// Confirm it prints 3, not 0 — this is the value-vs-pointer-receiver
	// distinction actually mattering.
}
