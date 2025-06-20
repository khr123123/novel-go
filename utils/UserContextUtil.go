package utils

import (
	"context"
)

type contextKey string

const (
	userIDKey   contextKey = "userID"
	authorIDKey contextKey = "authorID"
)

// 设置用户 ID
func SetUserID(ctx context.Context, userID int64) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}

// 获取用户 ID
func GetUserID(ctx context.Context) (int64, bool) {
	val := ctx.Value(userIDKey)
	if id, ok := val.(int64); ok {
		return id, true
	}
	return 0, false
}

// 设置作者 ID
func SetAuthorID(ctx context.Context, authorID int64) context.Context {
	return context.WithValue(ctx, authorIDKey, authorID)
}

// 获取作者 ID
func GetAuthorID(ctx context.Context) (int64, bool) {
	val := ctx.Value(authorIDKey)
	if id, ok := val.(int64); ok {
		return id, true
	}
	return 0, false
}
