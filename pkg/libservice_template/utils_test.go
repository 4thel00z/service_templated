package libservice_template

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGoEnv(t *testing.T) {

	env, err := GetGoEnv()
	assert.Nil(t, err)
	_, found := env["GOMOD"]
	assert.True(t, found)
}
