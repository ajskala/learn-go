# Lesson 3 — Type Switches, and Go's Real Lack of Exhaustiveness Checking

## No union types in Go

TS has `string | number | boolean`. Python has the same via `|`. Go has neither
— there is no built-in syntax for "this value is one of exactly these types."
The nearest tool is `any` (Lesson 1) plus runtime type inspection.

## Type switches — Go's `switch`/`match` for `any`

```go
func formatValue(value any) string {
	switch v := value.(type) {
	case string:
		return strings.ToUpper(v)
	case bool:
		if v {
			return "yes"
		}
		return "no"
	case int:
		return fmt.Sprintf("%d", v)
	default:
		return fmt.Sprintf("unhandled type: %T", v)
	}
}
```

`switch v := value.(type)` is special `switch` syntax found nowhere else in Go
— `value.(type)` only means something inside exactly this construct (Lesson 1's
`value.(string)` type *assertion* checks for one type; this type *switch*
checks a value against several at once). Inside each `case`, `v` is
automatically narrowed to that case's type — same narrowing idea as TS's
`typeof`/Python's `isinstance`/`match case str():`, Go's own spelling of it.
`%T` in a format string prints a value's concrete type — a handy debugging tool
for exactly this kind of code.

**Same case-ordering trap as Python's `bool`/`int`, different cause**: Go's
`case bool:` and `case int:` here don't have Python's "bool is a subclass of
int" issue (Go's `bool` and `int` are genuinely unrelated types, and a type
switch matches the *exact* concrete type stored in the `any`) — but always be
deliberate about `case` order in any type switch, since a `default` or an
interface-typed `case` earlier in the list can still shadow more specific cases
below it.

## No union types for your own types either — model it with an interface instead

TS's discriminated union (`Circle | Rectangle | Triangle`, tagged with `kind`)
has no direct Go equivalent. The idiomatic Go replacement: **one interface, many
implementing structs** — you already have all the pieces from Lesson 2:

```go
type Shape interface {
	Area() float64
}

type Circle struct{ Radius float64 }
type Rectangle struct{ Width, Height float64 }
type Triangle struct{ Base, Height float64 }

func (c Circle) Area() float64    { return math.Pi * c.Radius * c.Radius }
func (r Rectangle) Area() float64 { return r.Width * r.Height }
func (t Triangle) Area() float64  { return 0.5 * t.Base * t.Height }
```

Calling `.Area()` on a `Shape` variable is usually all you need — no `switch` at
all, because each type computes its own area (this is the idiomatic answer, and
it's arguably nicer than a `switch`: adding a new shape can't forget a case,
because there's no case list to forget). But sometimes you genuinely need to
branch differently per concrete type from *outside* the type — for that, a type
switch works on an interface value exactly like it does on `any`:

```go
func describeShape(s Shape) string {
	switch v := s.(type) {
	case Circle:
		return fmt.Sprintf("circle, radius %v", v.Radius)
	case Rectangle:
		return fmt.Sprintf("rectangle, %v x %v", v.Width, v.Height)
	case Triangle:
		return fmt.Sprintf("triangle, base %v height %v", v.Base, v.Height)
	default:
		return "unknown shape"
	}
}
```

## The honest gap: Go has no compile-time exhaustiveness checking

This is a real, deliberate difference from both your other tracks, not a minor
detail. TS's `never`-typed variable trick and Python's `assert_never` both give
you a **compile-time error** if you add a new union member and forget to handle
it somewhere. Go has **nothing built into the standard toolchain** that does
this. Add a `Pentagon` type implementing `Shape` and forget to add a `case
Pentagon:` to `describeShape` — it compiles cleanly, and `Pentagon` just falls
through to `default` silently at runtime. Nothing warns you at build time.

The idiomatic response to this gap is *not* silence — it's making the `default`
case fail loudly at runtime instead of compile time, so a missed case is at
least a crash you'll notice during testing, not a quietly wrong answer:

```go
default:
	panic(fmt.Sprintf("unhandled shape type: %T", v))
```

(There is a third-party static analysis tool, the `exhaustive` linter, that
some teams add to CI specifically to claw back compile-time-ish exhaustiveness
checking for type switches over a known, closed set of types — worth knowing it
exists, not something we're installing for this lesson.)

## Exercise

Fill in `exercise.go`. Check with `go vet ./...`, run with
`go run ./lessons/03-type-switches/`.
