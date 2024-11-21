package token

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"orca-service/global/security"
	"testing"
)

var ctx = context.Background()

func TestCreateAccessToken(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	store := NewRedisStore(rdb)
	user := security.UserDetail{
		Id:       "1",
		Username: "testUser",
		Email:    "test@example.com",
	}

	token, err := store.CreateAccessToken(user)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	val, err := rdb.Get(ctx, store.authToAccessKeyPrefix+":"+token).Result()
	assert.NoError(t, err)
	assert.NotEmpty(t, val)
}

func TestRefreshAccessToken(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	store := NewRedisStore(rdb)
	user := security.UserDetail{
		Id:       "1",
		Username: "testUser",
		Email:    "test@example.com",
	}

	token, _ := store.CreateAccessToken(user)
	newToken, err := store.RefreshAccessToken(token)
	assert.NoError(t, err)
	assert.Equal(t, token, newToken)
}

func TestRemoveAccessToken(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	store := NewRedisStore(rdb)
	user := security.UserDetail{
		Id:       "1",
		Username: "testUser",
		Email:    "test@example.com",
	}

	token, _ := store.CreateAccessToken(user)
	err := store.RemoveAccessTokenByUser(user)
	assert.NoError(t, err)

	val, err := rdb.Get(ctx, store.authToAccessKeyPrefix+token).Result()
	assert.Error(t, err)
	assert.Equal(t, "", val)
}

func TestVerifyAccessToken(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	store := NewRedisStore(rdb)
	user := security.UserDetail{
		Id:       "1",
		Username: "testUser",
		Email:    "test@example.com",
	}

	token, _ := store.CreateAccessToken(user)
	userDetails, err := store.VerifyAccessToken(token)
	assert.NoError(t, err)
	assert.Equal(t, user.Username, userDetails.Username)
}
