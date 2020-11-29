package errorutils_test

import (
	"errors"
	"github.com/parmaanu/goutils/errorutils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPanicOnErr(t *testing.T) {
	assert.Panics(t, func() { errorutils.PanicOnErr(errors.New("fatal error")) })
	assert.NotPanics(t, func() { errorutils.PanicOnErr(nil) })
}
