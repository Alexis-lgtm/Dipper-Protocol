package auth

import sdk "github.com/Dipper-Protocol/types"

type contextKey int // local to the auth module

const (
	contextKeySigners contextKey = iota
)

func WithSigners(ctx sdk.Context, account Account) sdk.Context {
	return ctx.WithValue(contextKeySigners, account)
}

func GetSigners(ctx sdk.Context) Account {
	v := ctx.Context().Value(contextKeySigners)
	if v == nil {
		return nil
	}
	return v.(Account)
}
