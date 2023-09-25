package z_test

import (
	"testing"

	"github.com/mkafonso/go-schema/z"
	"github.com/stretchr/testify/assert"
)

func TestStringMin_withValidString(t *testing.T) {
	// don't expect error when string is valid
	schema := z.NewStringMinSchema().Min(5).Message("Must be 5 or more characters long")
	result, err := schema.Parse("hello")

	assert.Equal(t, result, "hello")
	assert.NoError(t, err)
}
