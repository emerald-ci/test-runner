package lookup

import (
	"testing"

	"github.com/emerald-ci/test-runner/Godeps/_workspace/src/github.com/docker/libcompose/project"
)

func TestOsEnvLookup(t *testing.T) {
	// Putting bare minimun value for serviceName and config as there are
	// not important on this test.
	serviceName := "anything"
	config := &project.ServiceConfig{}

	osEnvLookup := &OsEnvLookup{}

	envs := osEnvLookup.Lookup("PATH", serviceName, config)
	if len(envs) != 1 {
		t.Fatalf("Expected envs to contains one element, but was %v", envs)
	}

	envs = osEnvLookup.Lookup("path", serviceName, config)
	if len(envs) != 0 {
		t.Fatalf("Expected envs to be empty, but was %v", envs)
	}

	envs = osEnvLookup.Lookup("DOES_NOT_EXIST", serviceName, config)
	if len(envs) != 0 {
		t.Fatalf("Expected envs to be empty, but was %v", envs)
	}
}
