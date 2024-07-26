package optional

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOptional(t *testing.T) {
	assert := assert.New(t)

	opt := Optional[int]{}

	assert.Equal(0,
		opt.OrElseGet(func() int {
			return 0
		}), "should be equal")

	assert.Panics(func() { opt.Get() }, "calling this function without value defined must panic")
}
