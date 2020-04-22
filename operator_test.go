package courier

import (
	"context"
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
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

	NewWithT(t).Expect(op.(*DoSomeThing).Param).To(Equal(1))
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
