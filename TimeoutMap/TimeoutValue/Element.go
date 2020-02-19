package TimeoutValue

type Element interface {
	UpdatedHandler()
	TimeoutedHandler()
}
