package courier

import (
	"context"
)

type Transport interface {
	Serve(router *Router) error
}

type Request interface {
	Do() RequestResult
}

type RequestResult interface {
	Into(v interface{}, metadata Metadata) error
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
