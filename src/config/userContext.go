package config

import (
	"context"

	"github.com/google/uuid"
)

type transactionIDContextKeyType struct{}

var (
	transactionIDContextKey = &transactionIDContextKeyType{}
)

func TransactionIDFromContext(ctx context.Context) string {
	transactionID := ctx.Value(transactionIDContextKey)
	if transactionID == nil {
		return ""
	}
	id, ok := transactionID.(string)
	if !ok {
		return ""
	}
	return id
}

func StartupContext() context.Context {
	ctx := context.Background()
	ctx = NewTransactionIDInContext(ctx)
	return ctx
}

func NewTransactionIDInContext(ctx context.Context) context.Context {
	return SetTransactionIDInContext(ctx, "")
}

func SetTransactionIDInContext(ctx context.Context, transactionID string) context.Context {
	if transactionID == "" {
		transactionID = uuid.New().String()
	}
	return context.WithValue(ctx, transactionIDContextKey, transactionID)
}
