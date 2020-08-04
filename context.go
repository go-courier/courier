package courier

import "context"

type ContextWith = func(ctx context.Context) context.Context

func ComposeContextWith(contextWiths ...ContextWith) ContextWith {
	return func(ctx context.Context) context.Context {
		for i := range contextWiths {
			ctx = contextWiths[i](ctx)
		}
		return ctx
	}
}
