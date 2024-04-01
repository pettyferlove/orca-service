package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestElapsedTimeString(t *testing.T) {
	timeString := ElapsedTimeString(time.Now().Unix() - 10000)
	assert.Equal(t, "10秒", timeString)

	timeString = ElapsedTimeString(time.Now().Unix() - 100000)
	assert.Equal(t, "1分钟", timeString)

	timeString = ElapsedTimeString(time.Now().Unix() - 1000000)
	assert.Equal(t, "16分钟", timeString)

	timeString = ElapsedTimeString(time.Now().Unix() - 10000000)
	assert.Equal(t, "2小时", timeString)

	timeString = ElapsedTimeString(time.Now().Unix() - 100000000)
	assert.Equal(t, "1天", timeString)

	timeString = ElapsedTimeString(time.Now().Unix() - 1000000000)
	assert.Equal(t, "11天", timeString)

	timeString = ElapsedTimeString(time.Now().Unix() - 10000000000)
	assert.Equal(t, "3月", timeString)

	timeString = ElapsedTimeString(time.Now().Unix() - 100000000000)
	assert.Equal(t, "3年", timeString)
}
