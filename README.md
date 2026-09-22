# learn-go

Structured, self-checking lessons for learning Go, written from the perspective
of someone who already knows TypeScript and Python — comparisons to both show up
throughout. Each lesson has a `README.md` (concepts) and an `exercise.go` (TODOs
you fill in yourself). Your work is checked two ways: `go vet` for
compile-correctness, and actually running the file to see real output.

## Prerequisites

- **Go 1.21 or later** (built and verified on 1.27.1). Check your version with
  `go version`. If you need to install or upgrade, get it from
  [go.dev/dl](https://go.dev/dl/) or `brew install go` on macOS.

## Setup

```
git clone https://github.com/ajskala/learn-go.git
cd learn-go
go vet ./...   # should report nothing — confirms your setup works
```

No separate dependency-install step — the lessons only use Go's standard
library, and `go.mod` (Go's `package.json`/`requirements.txt` equivalent) is
already committed.

## Workflow for every lesson

```
go vet ./...                                # checks every lesson for compile errors
go run ./lessons/<lesson-folder>/            # runs one lesson's exercise
```

Read a lesson's `README.md` first, then fill in the `TODO`s in its
`exercise.go`. Iterate with `go vet ./...` until it's clean, then run the file
to see the actual output.

**A Go-specific sharp edge worth knowing up front**: an unused import or an
unused local variable is a hard *compile error* in Go, not a lint warning like
in TS or Python. `go vet`/`go run` simply won't run at all until every import
you write is used and every variable you declare is read somewhere. This is
deliberate — Go's designers wanted to prevent dead code and stray imports from
silently accumulating.

We deliberately don't use `go build ./...` as the check step here: it compiles
everything correctly, but also writes an actual binary per `package main`
directory into your current folder — clutter you don't want in a lessons repo.
`go vet ./...` catches the same compile errors with zero file output.

## Lessons

1. **`01-basics-and-types`** — variables, type inference, zero values (Go has no
   `undefined`/`None`), slices, multiple return values, Go's `(value, error)`
   idiom for error handling, and `any` with type assertions.
2. **`02-structs-methods-interfaces`** — structs, exported vs. unexported
   (capitalization as the access modifier), methods, value vs. pointer
   receivers, embedding (Go's answer to inheritance), and interfaces satisfied
   implicitly — structural typing, always, no nominal option.
3. **`03-type-switches`** — type switches as Go's `switch`/`match` for `any`,
   modeling discriminated unions as one interface with many implementing
   structs, and the honest gap: Go has no compile-time exhaustiveness checking.
4. **`04-generics-and-variadics`** — variadic params, generic functions/structs
   with constraints, and what Go deliberately omits: default parameters and
   function overloading, plus the idiomatic workarounds for both.
