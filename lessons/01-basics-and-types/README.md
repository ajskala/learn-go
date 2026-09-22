# Lesson 1 — Go Basics, Types, Multiple Returns, and Errors

## Packages, modules, and running code

Every Go file belongs to a `package`. An executable program's entry point is a
file with `package main` and a `func main()`:

```go
package main

import "fmt"

func main() {
	fmt.Println("hello")
}
```

**A sharp edge to know about up front**: an unused import or an unused local
variable is a hard **compile error** in Go, not a lint warning like it is in
TS or Python — `go vet`/`go run` will refuse to run at all until every import
you write is actually used and every local variable you declare is actually
read somewhere. This is deliberate (Go's designers wanted to stop dead code and
stray imports from accumulating silently), but it does mean you'll add imports
to the `import (...)` block only as you actually need them while working
through the exercise below, not all up front.

A `go.mod` file at the project root declares a **module** (this project's is
`github.com/ajskala/learn-go`) — Go's equivalent of `package.json`/`pyproject.toml`,
tracking the module's own import path and its dependencies. No dependencies yet in
these lessons; the standard library covers everything so far.

Two commands you'll use constantly:

- `go run ./lessons/<lesson-folder>/` — compiles and runs one lesson, no leftover
  files. This is your `node file.ts` / `python file.py` equivalent.
- `go vet ./...` — analyzes every package in the module for compile-correctness and
  common mistakes, with **zero** file output. This is your `tsc`/`mypy` equivalent —
  run it constantly. (`go build ./...` also compiles everything, but it writes a
  binary per `package main` directory into your current folder — annoying clutter
  for a lessons repo, so we don't use it here for checking; `go vet` is enough for
  catching type/compile errors without the mess.)

## Variables and type inference

```go
var age int = 34          // explicit type
var name = "AJ"            // inferred as string
count := 0                  // ":=" — short variable declaration, infers the type,
                              // can ONLY be used inside a function body, never for
                              // package-level variables
```

`:=` is what you'll use almost everywhere inside functions — same spirit as TS's
`let x = ...` or Python's bare `x = ...`, just spelled differently. `var` is mostly
reserved for package-level declarations or when you want a type without an initial
value.

## Zero values — Go has no `undefined`/`None`/`null` for basic types

This is a real mental shift from both TS and Python. Every type has a **zero
value** — declaring a variable without initializing it never leaves it in an
"empty"/"undefined" state, it's automatically set to that type's zero value:

```go
var count int      // 0, not undefined
var name string      // "", not None
var ok bool           // false
var items []string     // nil (a nil slice behaves like an empty one for most uses)
```

There's no `unknown`/`undefined` distinction to worry about here — a declared
variable always holds a valid, usable value of its type from the moment it's
declared.

## Basic types and slices

```go
var city string = "Chicago"
var population int = 2700000
var ratio float64 = 3.14
var landmarks []string = []string{"Willis Tower", "Navy Pier"}
```

`[]string` is a **slice** — Go's primary "list of things" type, closest to TS's
`string[]` / Python's `list[str]`. (There's also a lower-level fixed-size `array`
type, `[2]string`, which you'll rarely reach for directly — slices are what
everyone actually uses day to day.)

## Functions, and Go's distinctive multiple return values

```go
func describeCity(city string, population int) string {
	return fmt.Sprintf("%s has %d people.", city, population)
}
```

`fmt.Sprintf` is Go's string-formatting function (`%s` for strings, `%d` for
integers) — the closest thing to a template literal / f-string, just as a function
call instead of special syntax.

Go functions can return **more than one value** — no TS/Python equivalent besides
manually returning a tuple/array. This isn't just a convenience; it's the
foundation of Go's whole approach to error handling:

```go
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide %v by zero", a)
	}
	return a / b, nil
}
```

## The idiom: functions return `(value, error)`, callers check it immediately

TS and Python both handle failure with exceptions — `throw`/`raise`, caught
somewhere up the call stack, possibly far from where the problem happened. Go
mostly doesn't do this (it *has* a `panic`/`recover` mechanism, but reserves it for
truly unrecoverable situations, not routine failure). Instead, anything that can
fail returns an `error` as its last return value, and the caller is expected to
check it **immediately**, right at the call site:

```go
result, err := divide(10, 2)
if err != nil {
	fmt.Println("Error:", err)
} else {
	fmt.Println("Result:", result)
}
```

`nil` is Go's "no error occurred" value (`error` is an interface type, and `nil`
is the zero value for interfaces — same `nil` concept as a nil slice earlier, one
consistent idea across the language). This `if err != nil` block is the single
most common pattern you will type in Go code — get comfortable with it now.

## `any` and type assertions — Go's answer to `unknown`

`any` (an alias for `interface{}`) can hold a value of *any* type — same escape
hatch as TS's `unknown` or Python's `object`. Getting a concrete type back out
requires a **type assertion**, and the two-value form is the safe one (never
panics, just tells you whether it worked):

```go
func safeLength(value any) int {
	s, ok := value.(string)
	if !ok {
		return 0
	}
	return len(s)
}
```

`value.(string)` attempts to treat `value` as a `string`. The two-value form `s,
ok := value.(string)` returns the asserted value plus a `bool` for whether it
actually was that type — `ok` is `false` (not a panic) if it wasn't. This is
directly analogous to TS's `typeof value === "string"` narrowing and Python's
`isinstance(value, str)` — same job, Go-flavored spelling.

## Exercise

Fill in the `TODO`s in `exercise.go`. Check with `go vet ./...` from the project
root, run with `go run ./lessons/01-basics-and-types/`.
