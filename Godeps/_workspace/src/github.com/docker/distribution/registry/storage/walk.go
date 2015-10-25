package storage

import (
	"errors"
	"fmt"

	"github.com/emerald-ci/test-runner/Godeps/_workspace/src/github.com/docker/distribution/context"
	storageDriver "github.com/emerald-ci/test-runner/Godeps/_workspace/src/github.com/docker/distribution/registry/storage/driver"
)

// SkipDir is used as a return value from onFileFunc to indicate that
// the directory named in the call is to be skipped. It is not returned
// as an error by any function.
var ErrSkipDir = errors.New("skip this directory")

// WalkFn is called once per file by Walk
// If the returned error is ErrSkipDir and fileInfo refers
// to a directory, the directory will not be entered and Walk
// will continue the traversal.  Otherwise Walk will return
type WalkFn func(fileInfo storageDriver.FileInfo) error

// Walk traverses a filesystem defined within driver, starting
// from the given path, calling f on each file
func Walk(ctx context.Context, driver storageDriver.StorageDriver, from string, f WalkFn) error {
	children, err := driver.List(ctx, from)
	if err != nil {
		return err
	}
	for _, child := range children {
		fileInfo, err := driver.Stat(ctx, child)
		if err != nil {
			return err
		}
		err = f(fileInfo)
		skipDir := (err == ErrSkipDir)
		if err != nil && !skipDir {
			return err
		}

		if fileInfo.IsDir() && !skipDir {
			Walk(ctx, driver, child, f)
		}
	}
	return nil
}

// pushError formats an error type given a path and an error
// and pushes it to a slice of errors
func pushError(errors []error, path string, err error) []error {
	return append(errors, fmt.Errorf("%s: %s", path, err))
}
