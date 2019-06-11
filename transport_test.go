package courier

import (
	"fmt"
)

type FakeTransport struct {
}

func (FakeTransport) UnmarshalOperator(op Operator) error {
	return nil
}

func (FakeTransport) Serve(router *Router) error {
	return fmt.Errorf("some thing wrong")
}

func ExampleRun() {
	var RouterRoot = NewRouter(&EmptyOperator{})

	Run(RouterRoot, &FakeTransport{})
	// Output:
}
