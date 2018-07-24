package courier

import (
	"fmt"
	"net/url"
	"reflect"
)

func NewOperatorFactory(op Operator, last bool) *OperatorMeta {
	opType := typeOfOperator(reflect.TypeOf(op))
	if opType.Kind() != reflect.Struct {
		panic(fmt.Errorf("operator must be a struct type, got %#v", op))
	}

	info := &OperatorMeta{}
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

type OperatorMeta struct {
	Type       reflect.Type
	ContextKey string
	Params     url.Values
	IsLast     bool
	Operator
}

func (o *OperatorMeta) String() string {
	if o.Params != nil {
		return o.Type.String() + "?" + o.Params.Encode()
	}
	return o.Type.String()
}

func (o *OperatorMeta) New() Operator {
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
