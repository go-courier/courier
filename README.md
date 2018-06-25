## Courier

[![GoDoc Widget](https://godoc.org/github.com/go-courier/courier?status.svg)](https://godoc.org/github.com/go-courier/courier)
[![Build Status](https://travis-ci.org/go-courier/courier.svg?branch=master)](https://travis-ci.org/go-courier/courier)
[![codecov](https://codecov.io/gh/go-courier/courier/branch/master/graph/badge.svg)](https://codecov.io/gh/go-courier/courier)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-courier/courier)](https://goreportcard.com/report/github.com/go-courier/courier)


### Motivation

Courier is inspired by [Go kit](https://github.com/go-kit/kit) with lots of enhancements on software engineering:
 
* Reducing codes by common processing units:
    * enumeration rules
    * parameters decoding and validating
    * normalized responses or errors and theirs metadata in different protocol like HTTP and gRPC
    * request sender with parameters encoder.
* API documents automatically generate [OpenAPI spec](https://www.openapis.org/) json file from codes.
    * generate client of target service by the generated OpenAPI spec json file.

### Core Concepts

#### `Operator`

`Operator` is the least processing unit in `Courier`. 
Any struct can be an `Operator` with a method `Output(ctx context.Context) (interface{}, error)`:

```go
type Operator interface {
    Output(ctx context.Context) (interface{}, error) 
}
```

The struct is used to define input parameters of the `Operator`, which will be unmarshal values by `Transport`.
The method `Output(ctx context.Context) (interface{}, error)` will return success or failed results of this `Operator` with matched input parameters.
We can also pick some value form `context.Context` which be set by upstream，`Operator` or `Transport`

#### `Router` and `Route`

`Router` is a carrier of `Operator`, so one `Router` have to contain at least one `Operator`.
`Router` can register another `Router`s as children, and the `Router`s will form multiple `Route`s which are organized and distributed by `Transport`

```go
var RouterRoot  = courier.NewRouter(&OperatorRoot{})
var RouterA = courier.NewRouter(&OperatorA{})
var RouterB = courier.NewRouter(&OperatorB{})

func init() {
    RouterRoot.Register(RouterA)
    RouterRoot.Register(RouterB)
    
    RouterA.Register(courier.NewRouter(&OperatorA1{}, &OperatorA2{}))
    RouterA.Register(courier.NewRouter(&OperatorA3{}, &OperatorA4{}))
    RouterB.Register(courier.NewRouter(&OperatorB1{}, &OperatorB2{}))
    
    // We will get routes:
    // OperatorRoot -> OperatorA -> OperatorA1 -> OperatorA2
    // OperatorRoot -> OperatorA -> OperatorA3 -> OperatorA4
    // OperatorRoot -> OperatorB -> OperatorB1 -> OperatorB2
}
```

#### `Transport`

`Transport` will do routing distribution and transform incoming data to `Operator`s on matched `Route`.
And send response which is transformed from results of last `Operator` with format of corresponding protocol.

##### HTTP 

TODO

##### gRPC

TODO
