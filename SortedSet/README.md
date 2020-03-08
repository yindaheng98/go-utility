# github.com/yindaheng98/go-utility/SortedSet

## Introduction

Sorted set is a special type set in which all the element can be sorted. This package implement a sorted set based on `map` in go and the package `Skiplist`.

## Usage

This package has a interface `Element` and a struct `SortedSet`

### Implement a interface

The interface `Element` defines the element that can be stored in `SortedSet`. This interface is very simple:

```go
type Element interface {
	GetName() string
}
```

The method `GetName()` defined the unique name of an element. If two elements' `GetName()` returns the same string, they will be regarded as the same element.

### Create a `SortedSet`

Just call the function `New(...)` and input a initial size, you can get a sorted set:

```go
set := New(100)
```

The initial size "100" does not mean that this set can not store No.101 elements. It only determines the index level of the skip list inside the sorted set.

### Add and change an element

The sorted set use a float number for sorting, called "weight". You can add an element with weight 1.2 like this:

```go
set.Update(element, 1.2)
```

Or you can increase the weight of an element by 0.2:

```go
set.UpdateDelta(element, 0.2)
```

Or even increase all the weight in the sorted set:

```go
set.DeltaUpdateAll(0.2)
```

### Get the weight of an element

```go
weight, ok := set.GetWeight(element)
```

if the element exists, the weight will be returned and ok will be true. If not, ok will be false.

### Output sorted element

This methods returns a list of the element, with the ascending order of their weight:

```go
func (set *SortedSet) SortedAll() []Element
```

Or if you do not want all the elements, you can get the smallest `n`(`uint64`) elements in the list:

```go
func (set *SortedSet) Sorted(n uint64) []Element
```

### How much elements?

This method returns how much element are in the sorted set.

```go
func (set *SortedSet) Count() uint64
```