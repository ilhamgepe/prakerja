package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasingPassword(t *testing.T) {
	password := "password"
	hash, err := HashPassword(password)
	// fmt.Println(hash)
	// fmt.Println(err)
	assert := assert.New(t)
	assert.Nil(err)
	assert.NotEqual(password, hash)
}
