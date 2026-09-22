package main

// Lesson 4 exercise
// Fill in each TODO. Run `go vet ./...` from the project root to check your code,
// then `go run ./lessons/04-generics-and-variadics/` to run it.
//
// Remember: unused imports and unused local variables are compile errors in Go —
// add to the import block below only once you actually use it.

import (
)

// TODO 1: write a variadic function `sum(nums ...int) int` returning the total.

// TODO 3: write a function `describeItem(name string, quantity int) string`
// that treats a zero `quantity` as meaning 1 (see the README's "zero value as
// default" idiom), returning a string like "5x widget".

// TODO 5: write a generic function `FirstElement[T any](items []T) (T, bool)`
// returning the first element and true, or the zero value and false if empty.

// TODO 6: define an interface `Number` permitting `int | float64`. Write a
// generic function `Sum[T Number](nums []T) T` returning the total.

// TODO 7: define a generic struct `Box[T any]` with field `Contents T`, and a
// method `GetContents() T` on it.

// TODO 8 (no overloading): write two separately-named functions,
// `CombineStrings(a, b string) string` and `CombineInts(a, b int) int`, each
// just concatenating/adding their arguments. (There's no way to name them both
// `Combine` in Go — that's the point of this TODO.)

func main() {
	// TODO 2: call `sum` with 0, 2, and 4 arguments, printing each result.

	// TODO 4: call `describeItem` once with quantity 0 and once with quantity
	// 5, printing both results.

	// TODO 5 (continued): call `FirstElement` once with a non-empty []int and
	// once with an empty []string, printing both results (both return values).

	// TODO 6 (continued): call `Sum` once with a []int and once with a
	// []float64, printing both results.

	// TODO 7 (continued): create one Box[string] and one Box[int], printing
	// both `.GetContents()` results.

	// TODO 8 (continued): call `CombineStrings` and `CombineInts` once each,
	// printing both results.
}
