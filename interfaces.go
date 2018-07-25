package courier

import (
	"context"
)

type Client interface {
	Do(operationID string, req interface{}, metas ...Metadata) Result
}

type Result interface {
	Into(v interface{}) (Metadata, error)
}

type Transport interface {
	Serve(router *Router) error
}

type Operator interface {
	Output(ctx context.Context) (interface{}, error)
}

type MetadataCarrier interface {
	Meta() Metadata
}

type OperatorWithParams interface {
	OperatorParams() map[string][]string
}

type OperatorWithoutOutput interface {
	Operator
	NoOutput()
}

type ContextProvider interface {
	Operator
	ContextKey() string
}

type DefaultsSetter interface {
	SetDefaults()
}
