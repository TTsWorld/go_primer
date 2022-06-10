package go_test_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var assertions *assert.Assertions

func TestSomeThing(t *testing.T) {
	assertions = assert.New(t)
	assertions.True(true, "True is true!")

}
