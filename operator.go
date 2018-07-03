package courier

import (
	"net/url"
	"reflect"
	"fmt"
)

func NewOperatorFactory(op Operator, last bool) *OperatorFactory {
	opType := typeOfOperator(reflect.TypeOf(op))
	if opType.Kind() != reflect.Struct {
		panic(fmt.Errorf("operator must be a struct type, got %#v", op))
	}

	info := &OperatorFactory{}
	info.IsLast = last

	info.Operator = op
	info.Type = typeOfOperator(reflect.TypeOf(op))

	if operatorWithParams, ok := op.(OperatorWithParams); ok {
		info.Params = url.Values(operatorWithParams.OperatorParams())
	}

	if !info.IsLast {
		if ctxKey, ok := op.(ContextProvider); ok {
			info.ContextKey = ctxKey.ContextKey()
		} else {
			info.ContextKey = info.Type.String()
		}
	}

	return info
}

func typeOfOperator(tpe reflect.Type) reflect.Type {
	for tpe.Kind() == reflect.Ptr {
		return typeOfOperator(tpe.Elem())
	}
	return tpe
}

type OperatorFactory struct {
	Type             reflect.Type
	ContextKey       string
	Params           url.Values
	IsLast           bool
	Operator
}

func (o *OperatorFactory) String() string {
	if o.Params != nil {
		return o.Type.String() + "?" + o.Params.Encode()
	}
	return o.Type.String()
}

func (o *OperatorFactory) New() (Operator, error) {
	rv := reflect.New(o.Type)
	op := rv.Interface().(Operator)

	if defaultsSetter, ok := op.(DefaultsSetter); ok {
		defaultsSetter.SetDefaults()
	}

	return op, nil
}

type EmptyOperator struct {
	OperatorWithoutOutput
}
