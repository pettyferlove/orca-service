package token

import (
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"orca-service/global/security/model"
	"testing"
)

func TestCreateAccessToken(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       1,  // use default DB
	})

	store := NewRedisStore(rdb)
	user := model.UserDetail{Username: "testUser", Password: "testPass"}

	token, err := store.CreateAccessToken(user)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// You can add more assertions here to check the state of the Redis store
	// For example, you might want to check that the token was correctly stored in Redis
}
