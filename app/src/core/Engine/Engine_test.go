package Engine

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetEngine(t *testing.T) {
	eng, _ := GetEngine("./data")
	assert.NotEqual(t, nil, eng, "not equal")
}
