// Package lazy is a light wrapper around sync.Once providing support for return values.
// It removes the burden of capturing return values via closures from the caller.
package lazy

import (
	"sync"
)

// Lazy wraps the provided function in a sync.Once. It does not invoke f itself.
// The function returned is safe for concurrent use.
func Lazy[T any](f func() T) func() T {
	var once sync.Once
	var t T
	fWrapper := func() {
		t = f()
	}
	return func() T {
		once.Do(fWrapper)
		return t
	}
}

// LazyError wraps the provided function in a sync.Once. It does not invoke f itself.
// The function returned is safe for concurrent use.
func LazyError[T any](f func() (T, error)) func() (T, error) {
	var once sync.Once
	var t T
	var err error
	fWrapper := func() {
		t, err = f()
	}
	return func() (T, error) {
		once.Do(fWrapper)
		return t, err
	}
}
