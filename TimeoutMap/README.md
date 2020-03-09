# github.com/yindaheng98/go-utility/TimeoutMap

## Introduction

This package defines a special type of map, which can time each element and delete it if is not updated within a certain period of time. Also, this map will emit the event (call some functions) when an new element was added or an exist element was updated, timeouted, or deleted.

## Usage

This package has an exported interface `Element` and an exported struct `TimeoutMap`.

### Make an element

Before using the `TimeoutMap`, you should make sure the elements you want to put into it implement the interface `Element`:

```go
type Element interface {
	GetID() string
	NewAddedHandler()
    UpdatedHandler()
	TimeoutedHandler()
	DeletedHandler()
}
```

The `Element` stands for the value in `TimeoutMap`, and the method `GetID()` stands for its key. And the other elements will be called when:

* `NewAddedHandler()`: the key of this element is not exist in `TimeoutMap` when update
* `UpdatedHandler()`: the key of this element is exist in `TimeoutMap` when update
* `TimeoutedHandler()`: the element is not updated within a certain period of time
* `DeletedHandler()`: the element is deleted from `TimeoutMap`

### Make an `TimeoutMap`

```go
import "github.com/yindaheng98/go-utility/TimeoutMap"
```

```go
m := TimeoutMap.New()
```

### Add or update an element

Use this method to update an element and change its content and its timing:

```go
func (m *TimeoutMap) UpdateInfo(el Element, timeout time.Duration)
```

Or use this method to only update the element:

```go
func (m *TimeoutMap) UpdateID(id string)
```

### Delete an element

Use this mathod to delete an element by its key:

```go
func (m *TimeoutMap) Delete(id string)
```

Or delete all the elements:

```go
func (m *TimeoutMap) DeleteAll()
```

### Find an element by its key

```go
func (m *TimeoutMap) GetElement(id string) (Element, bool)
```

If the key not exists, `nil,false` will returned.

### Get all the exist elements

```go
func (m *TimeoutMap) GetAll() []Element
```

### Get how much elements are in the `TimeoutMap`

```go
func (m *TimeoutMap) Count() int 
```