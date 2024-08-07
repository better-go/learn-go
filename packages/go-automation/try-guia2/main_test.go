package main

import (
	"runtime"
	"testing"
)

func Test(t *testing.T) {
	// arch + os
	t.Logf("os arch: %s, os system: %v", runtime.GOARCH, runtime.GOOS)
}
