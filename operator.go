package courier

import (
	"fmt"
	"net/url"
	"reflect"
)

// @deprecated
type OperatorMeta = OperatorFactory

func NewOperatorFactory(op Operator, last bool) *OperatorFactory {
	opType := typeOfOperator(reflect.TypeOf(op))
	if opType.Kind() != reflect.Struct {
		panic(fmt.Errorf("operator must be a struct type, got %#v", op))
	}

	meta := &OperatorFactory{}
	meta.IsLast = last

	meta.Operator = op

	if _, isOperatorWithoutOutput := op.(OperatorWithoutOutput); isOperatorWithoutOutput {
		meta.NoOutput = true
	}

	meta.Type = typeOfOperator(reflect.TypeOf(op))

	if operatorWithParams, ok := op.(OperatorWithParams); ok {
		meta.Params = operatorWithParams.OperatorParams()
	}

	if !meta.IsLast {
		if ctxKey, ok := op.(ContextProvider); ok {
			meta.ContextKey = ctxKey.ContextKey()
		} else {
			if ctxKey, ok := op.(oldContextProvider); ok {
				meta.ContextKey = ctxKey.ContextKey()
			} else {
				meta.ContextKey = meta.Type.String()
			}
		}
	}

	return meta
}

type oldContextProvider interface {
	ContextKey() string
}

func typeOfOperator(tpe reflect.Type) reflect.Type {
	for tpe.Kind() == reflect.Ptr {
		return typeOfOperator(tpe.Elem())
	}
	return tpe
}

type OperatorFactory struct {
	Type       reflect.Type
	ContextKey interface{}
	NoOutput   bool
	Params     url.Values
	IsLast     bool
	Operator   Operator
}

func (o *OperatorFactory) String() string {
	if o.Params != nil {
		return o.Type.String() + "?" + o.Params.Encode()
	}
	return o.Type.String()
}

func (o *OperatorFactory) New() Operator {
	rv := reflect.New(o.Type)
	op := rv.Interface().(Operator)

	if defaultsSetter, ok := op.(DefaultsSetter); ok {
		defaultsSetter.SetDefaults()
	}

	return op
}

type EmptyOperator struct {
	OperatorWithoutOutput
}
