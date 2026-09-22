# Lesson 2 — Structs, Methods, and Go's Implicit Interfaces

## Structs — Go's plain data type

```go
type Book struct {
	Title  string
	Author string
	Year   int
}

b := Book{Title: "Dune", Author: "Frank Herbert", Year: 1965}
fmt.Println(b.Title)
```

`struct` is Go's closest thing to a TS `interface`-typed object or a Python
`@dataclass` — a named bundle of fields. There's no separate "type" vs "value"
declaration step like TS's `interface Book {...}` plus a variable — the `type
Book struct {...}` line *is* both the shape definition and what you construct
instances of.

## Capitalization is Go's access-modifier system

Notice `Title`, not `title`. Go has no `public`/`private` keywords at all — a
field, function, type, or method is **exported** (visible outside its package)
if its name starts with an uppercase letter, and **unexported** (package-private)
if it starts with lowercase. `Book.title` (lowercase) would be invisible to any
code outside the package that defines `Book` — including, notably, the
`encoding/json` package and similar reflection-based tools, which is why almost
every struct field meant to hold real data is capitalized. This is a real,
load-bearing language rule, not a style preference.

## Methods — functions with a receiver

```go
func (b Book) Describe() string {
	return fmt.Sprintf("%s by %s (%d)", b.Title, b.Author, b.Year)
}

b.Describe()
```

`(b Book)` before the function name is the **receiver** — this makes `Describe`
a method on `Book`, callable as `b.Describe()`. This is Go's version of a class
method, but there's no `class` keyword: any type (not just structs) can have
methods attached to it, declared anywhere in the same package as the type, even
in a separate file.

## Value receivers vs. pointer receivers — a real gotcha

```go
func (b Book) SetTitleWrong(t string) {
	b.Title = t   // mutates a COPY — has no effect on the original
}

func (b *Book) SetTitle(t string) {
	b.Title = t   // mutates through a pointer — this actually works
}
```

`(b Book)` receives a **copy** of the struct — any mutation inside the method is
invisible to the caller once the method returns, a silent no-op bug that won't
raise any error. `(b *Book)` receives a **pointer** to the original — mutations
through it are real. Rule of thumb: if a method needs to modify the receiver, or
the struct is large enough that copying it is wasteful, use a pointer receiver.
Read-only methods on small structs are fine with value receivers. Go
conveniently lets you call either kind the same way (`b.SetTitle(...)`) whether
`b` itself is a value or a pointer — it inserts the `&`/dereference for you — so
the calling code looks identical either way; the receiver type is what decides
whether mutation sticks.

## Embedding — Go's alternative to inheritance

```go
type EBook struct {
	Book          // embedded, no field name — just the type
	FileSizeMB float64
}

e := EBook{Book: Book{Title: "Dune", Author: "Frank Herbert", Year: 1965}, FileSizeMB: 4.2}
fmt.Println(e.Title)      // "promoted" field — reads straight through to e.Book.Title
fmt.Println(e.Describe())  // promoted method too
```

Go has no `class Dog extends Animal`. Instead, you **embed** one struct inside
another (a field with a type but no name). The outer struct gets direct access
to the embedded struct's fields and methods as if they were its own — called
**promotion**. This is composition wearing inheritance's clothes: `EBook`
doesn't "become a" `Book` in the type-system sense TS's `extends` gives you
(there's no `EBook satisfies Book`-style relationship you can check), it just
has one glued inside it whose members are reachable without an extra `.Book.`
in front.

## Interfaces — satisfied implicitly (structural, like TypeScript!)

```go
type Shape interface {
	Area() float64
}

type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

var shape Shape = Square{Side: 4}   // works — no "implements" anywhere
```

This is the big payoff of everything above, and it directly parallels TS
Lesson 2's structural typing: `Square` never mentions `Shape` anywhere. There's
no `implements` keyword in Go at all. `Square` satisfies the `Shape` interface
purely because it has a matching `Area() float64` method — the same "shape
matching is enough" rule TS uses by default. Recall the three-way contrast this
completes across your lessons: **TypeScript** — objects/interfaces structural by
default. **Python** — classes nominal by default (`Protocol` opts back into
structural). **Go** — interfaces are *always* structural; there is no nominal
option for interfaces at all, ever.

## Exercise

Fill in `exercise.go`. Check with `go vet ./...`, run with
`go run ./lessons/02-structs-methods-interfaces/`.
