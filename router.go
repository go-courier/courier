package courier

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)

func NewRouter(operators ...Operator) *Router {
	ops := make([]Operator, 0)
	for i := range operators {
		op := operators[i]

		if withMiddleOperators, ok := op.(WithMiddleOperators); ok {
			ops = append(ops, withMiddleOperators.MiddleOperators()...)
		}

		ops = append(ops, op)
	}

	return &Router{
		operators: ops,
	}
}

// Router
type Router struct {
	parent    *Router
	operators []Operator
	children  map[*Router]bool
}

// Register child Router
func (router *Router) Register(r *Router) {
	if router.children == nil {
		router.children = map[*Router]bool{}
	}
	if r.parent != nil {
		panic(fmt.Errorf("router %v already registered to router %v", r, r.parent))
	}
	r.parent = router
	router.children[r] = true
}

func (router *Router) route() *Route {
	parent := router.parent
	operators := router.operators

	for parent != nil {
		operators = append(parent.operators, operators...)
		parent = parent.parent
	}

	return &Route{
		Operators: operators,
		last:      len(router.children) == 0,
	}
}

func (router *Router) Routes() (routes Routes) {
	maybeAppendRoute := func(router *Router) {
		route := router.route()

		if route.last && len(route.Operators) > 0 {
			routes = append(routes, route)
		}

		if len(router.children) > 0 {
			routes = append(routes, router.Routes()...)
		}
	}

	if len(router.children) == 0 {
		maybeAppendRoute(router)
		return
	}

	for childRouter := range router.children {
		maybeAppendRoute(childRouter)
	}

	return
}

type Routes []*Route

func (routes Routes) String() string {
	keys := make([]string, len(routes))
	for i, route := range routes {
		keys[i] = route.String()
	}
	sort.Strings(keys)
	return strings.Join(keys, "\n")
}

type Route struct {
	Operators []Operator
	last      bool
}

func (route *Route) OperatorFactories() (operatorFactories []*OperatorFactory) {
	lenOfOps := len(route.Operators)
	for i, op := range route.Operators {
		operatorFactories = append(operatorFactories, NewOperatorFactory(op, i == lenOfOps-1))
	}
	return
}

func (route *Route) String() string {
	buf := &bytes.Buffer{}
	operatorFactories := route.OperatorFactories()
	for i, operatorFactory := range operatorFactories {
		if i > 0 {
			buf.WriteString(" |> ")
		}
		buf.WriteString(operatorFactory.String())
	}
	return buf.String()
}
