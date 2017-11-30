// generated by "charlatan -dir=testdata/embedder -output=testdata/embedder/embedder.go Embedder".  DO NOT EDIT.

package main

import (
	"reflect"
	"testing"
)

// EmbedInvocation represents a single call of FakeEmbedder.Embed
type EmbedInvocation struct {
	Parameters struct {
		Ident3 string
	}
	Results struct {
		Ident4 string
	}
}

// OtherInvocation represents a single call of FakeEmbedder.Other
type OtherInvocation struct {
	Parameters struct {
		Ident1 string
	}
	Results struct {
		Ident2 string
	}
}

/*
FakeEmbedder is a mock implementation of Embedder for testing.
Use it in your tests as in this example:

	package example

	func TestWithEmbedder(t *testing.T) {
		f := &main.FakeEmbedder{
			EmbedHook: func(ident3 string) (ident4 string) {
				// ensure parameters meet expections, signal errors using t, etc
				return
			},
		}

		// test code goes here ...

		// assert state of FakeEmbed ...
		f.AssertEmbedCalledOnce(t)
	}

Create anonymous function implementations for only those interface methods that
should be called in the code under test.  This will force a painc if any
unexpected calls are made to FakeEmbed.
*/
type FakeEmbedder struct {
	EmbedHook func(string) string
	OtherHook func(string) string

	EmbedCalls []*EmbedInvocation
	OtherCalls []*OtherInvocation
}

// NewFakeEmbedderDefaultPanic returns an instance of FakeEmbedder with all hooks configured to panic
func NewFakeEmbedderDefaultPanic() *FakeEmbedder {
	return &FakeEmbedder{
		EmbedHook: func(string) (ident4 string) {
			panic("Unexpected call to Embedder.Embed")
			return
		},
		OtherHook: func(string) (ident2 string) {
			panic("Unexpected call to Embedder.Other")
			return
		},
	}
}

// NewFakeEmbedderDefaultFatal returns an instance of FakeEmbedder with all hooks configured to call t.Fatal
func NewFakeEmbedderDefaultFatal(t *testing.T) *FakeEmbedder {
	return &FakeEmbedder{
		EmbedHook: func(string) (ident4 string) {
			t.Fatal("Unexpected call to Embedder.Embed")
			return
		},
		OtherHook: func(string) (ident2 string) {
			t.Fatal("Unexpected call to Embedder.Other")
			return
		},
	}
}

// NewFakeEmbedderDefaultError returns an instance of FakeEmbedder with all hooks configured to call t.Error
func NewFakeEmbedderDefaultError(t *testing.T) *FakeEmbedder {
	return &FakeEmbedder{
		EmbedHook: func(string) (ident4 string) {
			t.Error("Unexpected call to Embedder.Embed")
			return
		},
		OtherHook: func(string) (ident2 string) {
			t.Error("Unexpected call to Embedder.Other")
			return
		},
	}
}

func (_f1 *FakeEmbedder) Embed(ident3 string) (ident4 string) {
	invocation := new(EmbedInvocation)

	invocation.Parameters.Ident3 = ident3

	ident4 = _f1.EmbedHook(ident3)

	invocation.Results.Ident4 = ident4

	_f1.EmbedCalls = append(_f1.EmbedCalls, invocation)

	return
}

// EmbedCalled returns true if FakeEmbedder.Embed was called
func (f *FakeEmbedder) EmbedCalled() bool {
	return len(f.EmbedCalls) != 0
}

// AssertEmbedCalled calls t.Error if FakeEmbedder.Embed was not called
func (f *FakeEmbedder) AssertEmbedCalled(t *testing.T) {
	t.Helper()
	if len(f.EmbedCalls) == 0 {
		t.Error("FakeEmbedder.Embed not called, expected at least one")
	}
}

// EmbedNotCalled returns true if FakeEmbedder.Embed was not called
func (f *FakeEmbedder) EmbedNotCalled() bool {
	return len(f.EmbedCalls) == 0
}

// AssertEmbedNotCalled calls t.Error if FakeEmbedder.Embed was called
func (f *FakeEmbedder) AssertEmbedNotCalled(t *testing.T) {
	t.Helper()
	if len(f.EmbedCalls) != 0 {
		t.Error("FakeEmbedder.Embed called, expected none")
	}
}

// EmbedCalledOnce returns true if FakeEmbedder.Embed was called exactly once
func (f *FakeEmbedder) EmbedCalledOnce() bool {
	return len(f.EmbedCalls) == 1
}

// AssertEmbedCalledOnce calls t.Error if FakeEmbedder.Embed was not called exactly once
func (f *FakeEmbedder) AssertEmbedCalledOnce(t *testing.T) {
	t.Helper()
	if len(f.EmbedCalls) != 1 {
		t.Errorf("FakeEmbedder.Embed called %d times, expected 1", len(f.EmbedCalls))
	}
}

// EmbedCalledN returns true if FakeEmbedder.Embed was called at least n times
func (f *FakeEmbedder) EmbedCalledN(n int) bool {
	return len(f.EmbedCalls) >= n
}

// AssertEmbedCalledN calls t.Error if FakeEmbedder.Embed was called less than n times
func (f *FakeEmbedder) AssertEmbedCalledN(t *testing.T, n int) {
	t.Helper()
	if len(f.EmbedCalls) < n {
		t.Errorf("FakeEmbedder.Embed called %d times, expected >= %d", len(f.EmbedCalls), n)
	}
}

// EmbedCalledWith returns true if FakeEmbedder.Embed was called with the given values
func (_f2 *FakeEmbedder) EmbedCalledWith(ident3 string) (found bool) {
	for _, call := range _f2.EmbedCalls {
		if reflect.DeepEqual(call.Parameters.Ident3, ident3) {
			found = true
			break
		}
	}

	return
}

// AssertEmbedCalledWith calls t.Error if FakeEmbedder.Embed was not called with the given values
func (_f3 *FakeEmbedder) AssertEmbedCalledWith(t *testing.T, ident3 string) {
	t.Helper()
	var found bool
	for _, call := range _f3.EmbedCalls {
		if reflect.DeepEqual(call.Parameters.Ident3, ident3) {
			found = true
			break
		}
	}

	if !found {
		t.Error("FakeEmbedder.Embed not called with expected parameters")
	}
}

// EmbedCalledOnceWith returns true if FakeEmbedder.Embed was called exactly once with the given values
func (_f4 *FakeEmbedder) EmbedCalledOnceWith(ident3 string) bool {
	var count int
	for _, call := range _f4.EmbedCalls {
		if reflect.DeepEqual(call.Parameters.Ident3, ident3) {
			count++
		}
	}

	return count == 1
}

// AssertEmbedCalledOnceWith calls t.Error if FakeEmbedder.Embed was not called exactly once with the given values
func (_f5 *FakeEmbedder) AssertEmbedCalledOnceWith(t *testing.T, ident3 string) {
	t.Helper()
	var count int
	for _, call := range _f5.EmbedCalls {
		if reflect.DeepEqual(call.Parameters.Ident3, ident3) {
			count++
		}
	}

	if count != 1 {
		t.Errorf("FakeEmbedder.Embed called %d times with expected parameters, expected one", count)
	}
}

// EmbedResultsForCall returns the result values for the first call to FakeEmbedder.Embed with the given values
func (_f6 *FakeEmbedder) EmbedResultsForCall(ident3 string) (ident4 string, found bool) {
	for _, call := range _f6.EmbedCalls {
		if reflect.DeepEqual(call.Parameters.Ident3, ident3) {
			ident4 = call.Results.Ident4
			found = true
			break
		}
	}

	return
}

func (_f7 *FakeEmbedder) Other(ident1 string) (ident2 string) {
	invocation := new(OtherInvocation)

	invocation.Parameters.Ident1 = ident1

	ident2 = _f7.OtherHook(ident1)

	invocation.Results.Ident2 = ident2

	_f7.OtherCalls = append(_f7.OtherCalls, invocation)

	return
}

// OtherCalled returns true if FakeEmbedder.Other was called
func (f *FakeEmbedder) OtherCalled() bool {
	return len(f.OtherCalls) != 0
}

// AssertOtherCalled calls t.Error if FakeEmbedder.Other was not called
func (f *FakeEmbedder) AssertOtherCalled(t *testing.T) {
	t.Helper()
	if len(f.OtherCalls) == 0 {
		t.Error("FakeEmbedder.Other not called, expected at least one")
	}
}

// OtherNotCalled returns true if FakeEmbedder.Other was not called
func (f *FakeEmbedder) OtherNotCalled() bool {
	return len(f.OtherCalls) == 0
}

// AssertOtherNotCalled calls t.Error if FakeEmbedder.Other was called
func (f *FakeEmbedder) AssertOtherNotCalled(t *testing.T) {
	t.Helper()
	if len(f.OtherCalls) != 0 {
		t.Error("FakeEmbedder.Other called, expected none")
	}
}

// OtherCalledOnce returns true if FakeEmbedder.Other was called exactly once
func (f *FakeEmbedder) OtherCalledOnce() bool {
	return len(f.OtherCalls) == 1
}

// AssertOtherCalledOnce calls t.Error if FakeEmbedder.Other was not called exactly once
func (f *FakeEmbedder) AssertOtherCalledOnce(t *testing.T) {
	t.Helper()
	if len(f.OtherCalls) != 1 {
		t.Errorf("FakeEmbedder.Other called %d times, expected 1", len(f.OtherCalls))
	}
}

// OtherCalledN returns true if FakeEmbedder.Other was called at least n times
func (f *FakeEmbedder) OtherCalledN(n int) bool {
	return len(f.OtherCalls) >= n
}

// AssertOtherCalledN calls t.Error if FakeEmbedder.Other was called less than n times
func (f *FakeEmbedder) AssertOtherCalledN(t *testing.T, n int) {
	t.Helper()
	if len(f.OtherCalls) < n {
		t.Errorf("FakeEmbedder.Other called %d times, expected >= %d", len(f.OtherCalls), n)
	}
}

// OtherCalledWith returns true if FakeEmbedder.Other was called with the given values
func (_f8 *FakeEmbedder) OtherCalledWith(ident1 string) (found bool) {
	for _, call := range _f8.OtherCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			found = true
			break
		}
	}

	return
}

// AssertOtherCalledWith calls t.Error if FakeEmbedder.Other was not called with the given values
func (_f9 *FakeEmbedder) AssertOtherCalledWith(t *testing.T, ident1 string) {
	t.Helper()
	var found bool
	for _, call := range _f9.OtherCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			found = true
			break
		}
	}

	if !found {
		t.Error("FakeEmbedder.Other not called with expected parameters")
	}
}

// OtherCalledOnceWith returns true if FakeEmbedder.Other was called exactly once with the given values
func (_f10 *FakeEmbedder) OtherCalledOnceWith(ident1 string) bool {
	var count int
	for _, call := range _f10.OtherCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			count++
		}
	}

	return count == 1
}

// AssertOtherCalledOnceWith calls t.Error if FakeEmbedder.Other was not called exactly once with the given values
func (_f11 *FakeEmbedder) AssertOtherCalledOnceWith(t *testing.T, ident1 string) {
	t.Helper()
	var count int
	for _, call := range _f11.OtherCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			count++
		}
	}

	if count != 1 {
		t.Errorf("FakeEmbedder.Other called %d times with expected parameters, expected one", count)
	}
}

// OtherResultsForCall returns the result values for the first call to FakeEmbedder.Other with the given values
func (_f12 *FakeEmbedder) OtherResultsForCall(ident1 string) (ident2 string, found bool) {
	for _, call := range _f12.OtherCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			ident2 = call.Results.Ident2
			found = true
			break
		}
	}

	return
}
