package genworld

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LoadNodeMgr(t *testing.T) {
	mgr, err := LoadNodeMgr("./data/core.yaml")
	assert.NoError(t, err)
	assert.NotNil(t, mgr)

	t.Logf("Test_LoadNodeMgr OK")
}
