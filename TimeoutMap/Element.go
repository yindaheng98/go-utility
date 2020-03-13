package TimeoutMap

import "github.com/yindaheng98/go-utility/TimeoutMap/TimeoutValue"

//The element used in TimeoutMap.
type Element interface {
	TimeoutValue.Element

	//Return the unique id of the Element.
	GetID() string

	//This method will be called when the Element is added to TimeoutMap.
	NewAddedHandler()

	//This method will be called when the Element is deleted from the TimeoutMap
	DeletedHandler()
}
