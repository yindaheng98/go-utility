package TimeoutValue

import (
	"fmt"
	"testing"
	"time"
)

type TestElement struct {
	id string
}

func (e *TestElement) TimeoutHandler() {
	fmt.Printf("Element %s is timeout.\n", e.id)
}

func TestTimeoutValue(t *testing.T) {
	v := New(&TestElement{"001"}, 1e8)
	go v.Run()
	go v.Update(nil, 1e9)
	time.Sleep(7e8)
	go v.Update(nil, 2e9)
	time.Sleep(3e9)
	go v.Stop()
	go v.Update(nil, 3e9)
	go v.Stop()
	go v.Run()
	go v.Update(nil, 40e9)
	//go v.Stop()
	go v.Update(nil, 50e9)
	//go v.Stop()
	time.Sleep(1e9)
	go v.Update(nil, 1e9)
	time.Sleep(2e9)
	v.Stop()
}
