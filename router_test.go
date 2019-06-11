package courier

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func ExampleNewRouter() {
	var RouterRoot = NewRouter(&EmptyOperator{})
	var RouterA = NewRouter(&OperatorA{})
	var RouterB = NewRouter(&OperatorB{})

	RouterRoot.Register(RouterA)
	RouterRoot.Register(RouterB)

	RouterA.Register(NewRouter(&OperatorA1{}))
	RouterA.Register(NewRouter(&OperatorA2{}))
	RouterB.Register(NewRouter(&OperatorB1{}, OperatorB2{}))

	fmt.Println(RouterRoot.Routes())
	// Output:
	//courier.OperatorA |> courier.OperatorA1?allowedRoles=ADMIN&allowedRoles=OWNER
	//courier.OperatorA |> courier.OperatorA2
	//courier.OperatorB |> courier.OperatorB1 |> courier.OperatorB2
}

type OperatorA struct{}

func (OperatorA) ContextKey() string {
	return "OperatorA"
}

func (OperatorA) Output(ctx context.Context) (interface{}, error) {
	return nil, nil
}

type OperatorA1 struct{}

func (OperatorA1) OperatorParams() map[string][]string {
	return map[string][]string{
		"allowedRoles": {"ADMIN", "OWNER"},
	}
}

func (OperatorA1) Output(ctx context.Context) (interface{}, error) {
	return nil, nil
}

type OperatorA2 struct{}

func (OperatorA2) Output(ctx context.Context) (interface{}, error) {
	return nil, nil
}

type OperatorB struct{}

func (OperatorB) Output(ctx context.Context) (interface{}, error) {
	return nil, nil
}

type OperatorB1 struct{}

func (OperatorB1) Output(ctx context.Context) (interface{}, error) {
	return nil, nil
}

type OperatorB2 struct{}

func (OperatorB2) Output(ctx context.Context) (interface{}, error) {
	return nil, nil
}

func TestRegister(t *testing.T) {
	var RouterRoot = NewRouter(&EmptyOperator{})
	var RouterA = NewRouter(&OperatorA{})
	RouterRoot.Register(RouterA)

	err := TryCatch(func() {
		RouterRoot.Register(RouterA)
	})
	require.Error(t, err)
}
