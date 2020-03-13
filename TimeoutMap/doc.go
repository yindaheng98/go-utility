/*
Package TimeoutMap defines a special type of map, which can time each element and delete it if is not updated within a certain period of time.
Also, this map will emit the event (call some functions) when an new element was added or an exist element was updated, timeouted, or deleted.

Source code and other details for the project are available at GitHub:

	https://github.com/yindaheng98/go-utility

Installation

The only requirement is the Go Programming Language, at least version 1.13.

	$ go get github.com/yindaheng98/go-utility

*/
package TimeoutMap
