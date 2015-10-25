package memory

import (
	"testing"

	"github.com/emerald-ci/test-runner/Godeps/_workspace/src/github.com/docker/distribution/registry/storage/cache"
)

// TestInMemoryBlobInfoCache checks the in memory implementation is working
// correctly.
func TestInMemoryBlobInfoCache(t *testing.T) {
	cache.CheckBlobDescriptorCache(t, NewInMemoryBlobDescriptorCacheProvider())
}
