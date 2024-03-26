package token

import (
	"github.com/stretchr/testify/assert"
	"orca-service/global/security/model"
	"testing"
	"time"
)

func TestJwtStore(t *testing.T) {
	key := []byte("your_secret_key")
	store := NewJwtStore(key)

	user := model.UserDetail{
		Id:       "1",
		Username: "testUser",
		Email:    "test@example.com",
	}

	// Test CreateAccessToken
	token, err := store.CreateAccessToken(user)
	assert.Nil(t, err)
	assert.NotEmpty(t, token)

	// Test VerifyAccessToken
	userDetail, err := store.VerifyAccessToken(token)
	assert.Nil(t, err)
	assert.Equal(t, user.Username, userDetail.Username)
	assert.Equal(t, user.Email, userDetail.Email)

	// Test RefreshAccessToken
	// 停顿
	time.Sleep(1 * time.Second)
	newToken, err := store.RefreshAccessToken(token)
	assert.Nil(t, err)
	assert.NotEmpty(t, newToken)
	assert.NotEqual(t, token, newToken)

	// Test RemoveAccessToken
	err = store.RemoveAccessToken(user)
	assert.Nil(t, err)
}
