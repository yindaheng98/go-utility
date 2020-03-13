package TimeoutValue

//Element used in TimeoutValue
type Element interface {

	//This method will be called when the element was updated.
	UpdatedHandler()

	//This method will be called when the element was timeout.
	TimeoutedHandler()
}
