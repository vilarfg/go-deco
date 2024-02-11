# Deco

Decorator chaining for [Go](https://go.dev/).

Deco is a lightweight, generic Go package designed to facilitate the chaining of decorators, providing a simple yet powerful API for enhancing or modifying behavior in a composable manner. By reducing chains to a single function (a new decorator), Deco ensures that these chains become immutable and therefore thread-safe, making it ideal for applications requiring high concurrency or immutable data patterns. Whether for HTTP middleware, logging, data processing, or other processing layers, Deco offers a clean and flexible solution.

## Features

- **Generic**: works with any data type, providing flexibility across different use cases.
- **Chaining**: easily chain multiple decorators together to build complex processing pipelines.
- **Extendibility**: dynamically extend existing chains with additional decorators.
- **Simplicity**: minimal and straightforward API that's easy to integrate and use within your projects.

## Installation

```shell
go get github.com/vilarfg/go-deco
```

## Usage

```go
import (
    "github.com/vilarfg/go-deco"
    "strings" 
)

func stop(s string) string      { return s + ". Stop it." }
func name(s string) string      { return "Ford, " + s }
func turning(s string) string   { return "you're turning into " + s }
func uppercase(s string) string { return strings.ToUpper(s) }
func identity(s string) string  { return s }

// An empty chain produces a decorator
// that does not decorate at all.
// Internally is the identity function.
var chain0 = deco.Chain[string]()

// A chain of one or more decorators produces a new decorator.
// Decorators are applied in reverse order,
// `name` will be applied before `stop`.
// Notice how the type of the thing being decorated is inferred.
var chain1 = deco.Chain(stop, name)

// Decorators can be extended to produce a new decorator.
// `turning` will be applied before the original decorator.
var chain2 = chain1.Extend(turning)

// Which, in turn, can be further extended.
var chain3 = chain2.Extend(uppercase)

// nil values are filtered out
var chain4 = chain3.Extend(nil, identity)

println(chain0("whale"))            // "whale"
println(chain1("bowl of petunias")) // "Ford, bowl of petunias. Stop it."
println(chain2("a sofa"))           // "Ford, you're turning into a sofa. Stop it."
println(chain3("yarn"))             // "Ford, you're turning into YARN. Stop it."
println(chain4("a penguin"))        // "Ford, you're turning into A PENGUIN. Stop it.
```

## API

### `Chain`

- `func Chain[T any](decorators ...Decorator[T]) Decorator[T]`
  - Chains multiple decorators into a single decorator.
  - Decorators are applied in reverse order

### `Decorator`

- `type Decorator[T any] func(T) T`
  - A function that takes a value of type `T` and returns a potentially modified value of the same type.

#### `Extend`

- `func (m Decorator[T]) Extend(decorators ...Decorator[T]) Decorator[T]`
  - Extends an existing decorator with additional decorators.
  - Decorators are applied in reverse order

#### `Apply`

- `func (m Decorator[T]) Apply(t T) T`
  - Applies the decorator to a value.

## License

[MIT](LICENSE). Copyright © 2024 Fernando G. Vilar.
