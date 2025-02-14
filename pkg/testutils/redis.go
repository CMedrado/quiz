package testutils

import (
	"context"
	"fmt"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
)

func NewRedisClient(t *testing.T, ctx context.Context) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", "localhost", "6379"),
	})

	err := rdb.Ping(ctx).Err()
	require.NoError(t, err)

	return rdb
}
