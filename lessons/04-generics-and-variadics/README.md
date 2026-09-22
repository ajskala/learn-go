# Lesson 4 — Generics, Variadics, and What Go Leaves Out On Purpose

Go's other three lessons kept finding close parallels to TS/Python. This one is
different: Go deliberately **does not have** default parameters or function
overloading — full stop, not "different spelling," actually absent. Knowing
what's missing and why is as useful as knowing what's there.

## Variadic parameters — Go's rest parameter

```go
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

sum(1, 2, 3)   // nums is []int{1, 2, 3} inside the function
sum()           // nums is []int{}
```

`...int` collects trailing arguments into a real `[]int` — same job as TS's
`...nums: number[]` / Python's `*nums`. `for _, n := range nums` is Go's
for-each loop; `_` is the **blank identifier**, discarding a value you're not
using (here, the index) — required because, unlike TS/Python, an unused
declared variable is a compile error (Lesson 1), and `range` always gives you
both an index and a value.

## No default parameters — at all

There's no equivalent of TS's `quantity: number = 1` or Python's `quantity: int
= 1`. Every parameter must be supplied on every call, always. The idiomatic
workarounds, both worth knowing:

**Zero values as implicit defaults** (works when the type's zero value happens
to be a sensible default):

```go
func describeItem(name string, quantity int) string {
	if quantity == 0 {
		quantity = 1
	}
	return fmt.Sprintf("%dx %s", quantity, name)
}
```

**A dedicated options struct**, for functions with several optional settings —
the struct's zero value doubles as "all defaults":

```go
type ItemOptions struct {
	Quantity int
	Note     string
}

func describeItemWithOptions(name string, opts ItemOptions) string {
	quantity := opts.Quantity
	if quantity == 0 {
		quantity = 1
	}
	base := fmt.Sprintf("%dx %s", quantity, name)
	if opts.Note != "" {
		return fmt.Sprintf("%s (%s)", base, opts.Note)
	}
	return base
}

describeItemWithOptions("widget", ItemOptions{})                       // "1x widget"
describeItemWithOptions("widget", ItemOptions{Quantity: 5, Note: "rush"}) // "5x widget (rush)"
```

## No function overloading — at all

Go doesn't allow two functions with the same name in one package, under any
circumstances — not even with different parameter types. TS/Python's `combine`
overload trick (string-pair vs. number-pair, same name) has no Go equivalent.
The idiomatic fix is refreshingly blunt: **give them different names**.

```go
func CombineStrings(a, b string) string { return a + b }
func CombineInts(a, b int) int          { return a + b }
```

This is a real philosophical difference, not a missing feature Go forgot: Go's
designers traded away these features on purpose, favoring one obvious way to
call any given function over TS/Python's more flexible-but-more-implicit call
shapes.

## Generics (Go 1.18+)

```go
func FirstElement[T any](items []T) (T, bool) {
	var zero T
	if len(items) == 0 {
		return zero, false
	}
	return items[0], true
}

FirstElement([]int{1, 2, 3})     // (1, true)
FirstElement([]string{})          // ("", false)
```

`[T any]` declares a type parameter, same job as TS's `<T>` / Python's PEP 695
`[T]`. `var zero T` declares a variable holding `T`'s zero value (Lesson 1,
paying off again) — used here for the "nothing to return" case, since Go has no
`T | undefined`/`T | None` to fall back on for an arbitrary generic type. The
`(value, bool)` return shape is the same idiom you already know from map
lookups and type assertions — "did this actually work" as an explicit second
value, Go's answer to a nullable return.

## Generic constraints

```go
type Number interface {
	int | float64
}

func Sum[T Number](nums []T) T {
	var total T
	for _, n := range nums {
		total += n
	}
	return total
}

Sum([]int{1, 2, 3})          // T = int
Sum([]float64{1.5, 2.5})      // T = float64
```

An interface used as a constraint (instead of describing methods, like Lesson
2's `Shape`) can list a **union of permitted types** with `|` — this is the one
place `|` for unions genuinely exists in Go, scoped specifically to generic
constraints, not available for ordinary variable types the way TS/Python use it
everywhere.

## Generic structs

```go
type Box[T any] struct {
	Contents T
}

func (b Box[T]) GetContents() T {
	return b.Contents
}

stringBox := Box[string]{Contents: "hello"}
numberBox := Box[int]{Contents: 42}
```

Same idea as TS's `class Box<T>` / Python's `class Box[T]:`. Note the type
argument is explicit here (`Box[string]{...}`) — Go's generic type inference for
struct literals is more limited than for function calls, so you'll write `[T]`
by hand more often on the struct/type side than you did with plain generic
functions above.

## Exercise

Fill in `exercise.go`. Check with `go vet ./...`, run with
`go run ./lessons/04-generics-and-variadics/`.
