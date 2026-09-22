package main

// Lesson 3 exercise
// Fill in each TODO. Run `go vet ./...` from the project root to check your code,
// then `go run ./lessons/03-type-switches/` to run it.
//
// Remember: unused imports and unused local variables are compile errors in Go —
// add to the import block below only once you actually use it.

import (
)

// TODO 1: write a function `formatValue(value any) string` using a type switch:
//   - string -> uppercased (strings.ToUpper)
//   - bool -> "yes" or "no"
//   - int -> formatted with fmt.Sprintf("%d", v)
//   - default -> fmt.Sprintf("unhandled type: %T", v)

// TODO 3: define an interface `Shape` with one method: `Area() float64`.

// TODO 3 (continued): define three structs implementing Shape:
//   Circle    { Radius float64 }
//   Rectangle { Width, Height float64 }
//   Triangle  { Base, Height float64 }
// with Area() methods (circle: math.Pi * r * r; rectangle: w * h;
// triangle: 0.5 * b * h). You'll need "math" in the imports.

// TODO 5: write a function `describeShape(s Shape) string` using a type switch
// over the concrete Circle/Rectangle/Triangle types, returning a description
// string for each (format however you like), with a `default` branch that
// panics via fmt.Sprintf("unhandled shape type: %T", v) — since Go can't catch
// a missed case at compile time, this is the idiomatic "fail loudly instead of
// silently" fallback.

func main() {
	// TODO 2: call `formatValue` with a string, a bool, and an int, printing
	// each result.

	// TODO 4: create one Circle, one Rectangle, and one Triangle. For each,
	// print the result of calling .Area() on it directly (through the Shape
	// interface) — no switch needed here, this is the idiomatic
	// "let each type answer for itself" approach from the README.

	// TODO 5 (continued): call `describeShape` with each of your three shapes
	// and print the results.
}
