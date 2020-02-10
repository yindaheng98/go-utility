package TimeoutMap

import "github.com/yindaheng98/go-utility/TimeoutMap/TimeoutValue"

type Element interface {
	TimeoutValue.Element
	GetID() string
	NewAddedHandler()
}
