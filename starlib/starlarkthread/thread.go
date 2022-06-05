// Copyright 2021 Edward McFarlane. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package starlarkthread

import (
	"context"
	"fmt"
	"io"
	"runtime"
	"sync"
	"testing"

	"github.com/emcfarlane/larking/starlib/starlarkstruct"
	"go.starlark.net/starlark"
)

func NewModule() *starlarkstruct.Module {
	return &starlarkstruct.Module{
		Name: "thread",
		Members: starlark.StringDict{
			"os":   starlark.String(runtime.GOOS),
			"arch": starlark.String(runtime.GOARCH),
		},
	}
}

const ctxkey = "context"

// SetContext sets the thread context.
func SetContext(thread *starlark.Thread, ctx context.Context) {
	thread.SetLocal(ctxkey, ctx)
}

// GetContext gets the thread context or returns a TODO context.
// ctx := starlarkthread.Context(thread)
func GetContext(thread *starlark.Thread) context.Context {
	if ctx, ok := thread.Local(ctxkey).(context.Context); ok {
		return ctx
	}
	return context.TODO()
}

// Resource is a starlark.Value that requires close handling.
type Resource interface {
	starlark.Value
	io.Closer
}

const rsckey = "resources"

// ResourceStore is a thread local storage map for adding resources.
// It is thread safe.
type ResourceStore struct {
	mu sync.Mutex
	m  map[Resource]bool // resource whether it's living
}

// WithResourceStore returns a cleanup function. It is required for
// packages that add resources.
func WithResourceStore(thread *starlark.Thread) func() error {
	store := &ResourceStore{
		m: make(map[Resource]bool),
	}
	SetResourceStore(thread, store)
	// TODO: runtime.SetFinalizer?
	return func() error { return CloseResources(thread) }
}

func SetResourceStore(thread *starlark.Thread, store *ResourceStore) {
	thread.SetLocal(rsckey, store)
}
func GetResourceStore(thread *starlark.Thread) (*ResourceStore, error) {
	store, ok := thread.Local(rsckey).(*ResourceStore)
	if !ok {
		return nil, fmt.Errorf("thread missing resource store")
	}
	return store, nil
}

func AddResource(thread *starlark.Thread, rsc Resource) error {
	store, err := GetResourceStore(thread)
	if err != nil {
		return err
	}
	store.mu.Lock()
	store.m[rsc] = true
	store.mu.Unlock()
	return nil
}

func (s *ResourceStore) Close() (firstErr error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for rsc, open := range s.m {
		if !open {
			continue
		}
		if err := rsc.Close(); err != nil && firstErr == nil {
			// TODO: chain errors?
			firstErr = err
		}
		delete(s.m, rsc)
	}
	return
}

func CloseResources(thread *starlark.Thread) (firstErr error) {
	store, err := GetResourceStore(thread)
	if err != nil {
		return err
	}
	return store.Close()
}

// AssertOption implements starlarkassert.TestOption
// Add like so:
//
// 	starlarkassert.RunTests(t, "*.star", globals, starlarkthread.AssertOption)
//
func AssertOption(t testing.TB, thread *starlark.Thread) func() {
	close := WithResourceStore(thread)
	return func() {
		if err := close(); err != nil {
			t.Error(err, "failed to close resources")
		}
	}
}