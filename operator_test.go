package courier

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type DoSomeThing struct {
	Param int
}

func (req *DoSomeThing) SetDefaults() {
	if req != nil {
		if req.Param == 0 {
			req.Param = 1
		}
	}
}

func (DoSomeThing) ContextKey() string {
	return "DoSomeThing"
}

func (DoSomeThing) Output(ctx context.Context) (interface{}, error) {
	return nil, nil
}

func TestNewOperatorFactory(t *testing.T) {
	opInfo := NewOperatorFactory(&DoSomeThing{}, true)

	op := opInfo.New()
	assert.Equal(t, 1, op.(*DoSomeThing).Param)
}

func TryCatch(f func()) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()

	f()
	return
}
