package forwarding

import (
	"github.com/rollicks-c/kgate/internal/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHash(t *testing.T) {
	f1 := managedForwarder{
		group:       config.PortGroup{Name: "test-group"},
		serviceName: "test-service-2",
		namespace:   "default",
		localPort:   "8080",
		remotePort:  "80",
	}

	f2 := f1
	assert.Equal(t, f1.hash(), f2.hash())

	f3 := f1
	f3.localPort = "8081"
	assert.NotEqual(t, f1.hash(), f3.hash())

	f4 := f1
	f4.remotePort = "81"
	assert.NotEqual(t, f1.hash(), f4.hash())

	f5 := f1
	f5.serviceName = "test-service-3"
	assert.NotEqual(t, f1.hash(), f5.hash())

	f6 := f1
	f6.namespace = "test-namespace"
	assert.NotEqual(t, f1.hash(), f6.hash())

	f7 := f1
	f7.group.Name = "test-group-2"
	assert.NotEqual(t, f1.hash(), f7.hash())
}
