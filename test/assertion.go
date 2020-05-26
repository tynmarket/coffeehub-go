package test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// AssertJSONEq is assert.JSONEq
func AssertJSONEq(t *testing.T, exp string, actual string) {
	assert.JSONEq(t, strings.ReplaceAll(exp, "\n", ""), actual, "error message %s", "formatted")
}
