// generated by "charlatan -dir=testdata/channeler -output=testdata/channeler/channeler.go Channeler".  DO NOT EDIT.

package main

import "reflect"

// ChannelerChannelInvocation represents a single call of FakeChanneler.Channel
type ChannelerChannelInvocation struct {
	Parameters struct {
		Ident1 chan int
	}
	Results struct {
		Ident2 chan int
	}
}

// NewChannelerChannelInvocation creates a new instance of ChannelerChannelInvocation
func NewChannelerChannelInvocation(ident1 chan int, ident2 chan int) *ChannelerChannelInvocation {
	invocation := new(ChannelerChannelInvocation)

	invocation.Parameters.Ident1 = ident1

	invocation.Results.Ident2 = ident2

	return invocation
}

// ChannelerChannelReceiveInvocation represents a single call of FakeChanneler.ChannelReceive
type ChannelerChannelReceiveInvocation struct {
	Parameters struct {
		Ident1 <-chan int
	}
	Results struct {
		Ident2 <-chan int
	}
}

// NewChannelerChannelReceiveInvocation creates a new instance of ChannelerChannelReceiveInvocation
func NewChannelerChannelReceiveInvocation(ident1 <-chan int, ident2 <-chan int) *ChannelerChannelReceiveInvocation {
	invocation := new(ChannelerChannelReceiveInvocation)

	invocation.Parameters.Ident1 = ident1

	invocation.Results.Ident2 = ident2

	return invocation
}

// ChannelerChannelSendInvocation represents a single call of FakeChanneler.ChannelSend
type ChannelerChannelSendInvocation struct {
	Parameters struct {
		Ident1 chan<- int
	}
	Results struct {
		Ident2 chan<- int
	}
}

// NewChannelerChannelSendInvocation creates a new instance of ChannelerChannelSendInvocation
func NewChannelerChannelSendInvocation(ident1 chan<- int, ident2 chan<- int) *ChannelerChannelSendInvocation {
	invocation := new(ChannelerChannelSendInvocation)

	invocation.Parameters.Ident1 = ident1

	invocation.Results.Ident2 = ident2

	return invocation
}

// ChannelerChannelPointerInvocation represents a single call of FakeChanneler.ChannelPointer
type ChannelerChannelPointerInvocation struct {
	Parameters struct {
		Ident1 *chan int
	}
	Results struct {
		Ident2 *chan int
	}
}

// NewChannelerChannelPointerInvocation creates a new instance of ChannelerChannelPointerInvocation
func NewChannelerChannelPointerInvocation(ident1 *chan int, ident2 *chan int) *ChannelerChannelPointerInvocation {
	invocation := new(ChannelerChannelPointerInvocation)

	invocation.Parameters.Ident1 = ident1

	invocation.Results.Ident2 = ident2

	return invocation
}

// ChannelerChannelInterfaceInvocation represents a single call of FakeChanneler.ChannelInterface
type ChannelerChannelInterfaceInvocation struct {
	Parameters struct {
		Ident1 chan interface{}
	}
	Results struct {
		Ident2 chan interface{}
	}
}

// NewChannelerChannelInterfaceInvocation creates a new instance of ChannelerChannelInterfaceInvocation
func NewChannelerChannelInterfaceInvocation(ident1 chan interface{}, ident2 chan interface{}) *ChannelerChannelInterfaceInvocation {
	invocation := new(ChannelerChannelInterfaceInvocation)

	invocation.Parameters.Ident1 = ident1

	invocation.Results.Ident2 = ident2

	return invocation
}

// ChannelerTestingT represents the methods of "testing".T used by charlatan Fakes.  It avoids importing the testing package.
type ChannelerTestingT interface {
	Error(...interface{})
	Errorf(string, ...interface{})
	Fatal(...interface{})
	Helper()
}

/*
FakeChanneler is a mock implementation of Channeler for testing.
Use it in your tests as in this example:

	package example

	func TestWithChanneler(t *testing.T) {
		f := &main.FakeChanneler{
			ChannelHook: func(ident1 chan int) (ident2 chan int) {
				// ensure parameters meet expections, signal errors using t, etc
				return
			},
		}

		// test code goes here ...

		// assert state of FakeChannel ...
		f.AssertChannelCalledOnce(t)
	}

Create anonymous function implementations for only those interface methods that
should be called in the code under test.  This will force a panic if any
unexpected calls are made to FakeChannel.
*/
type FakeChanneler struct {
	ChannelHook          func(chan int) chan int
	ChannelReceiveHook   func(<-chan int) <-chan int
	ChannelSendHook      func(chan<- int) chan<- int
	ChannelPointerHook   func(*chan int) *chan int
	ChannelInterfaceHook func(chan interface{}) chan interface{}

	ChannelCalls          []*ChannelerChannelInvocation
	ChannelReceiveCalls   []*ChannelerChannelReceiveInvocation
	ChannelSendCalls      []*ChannelerChannelSendInvocation
	ChannelPointerCalls   []*ChannelerChannelPointerInvocation
	ChannelInterfaceCalls []*ChannelerChannelInterfaceInvocation
}

// NewFakeChannelerDefaultPanic returns an instance of FakeChanneler with all hooks configured to panic
func NewFakeChannelerDefaultPanic() *FakeChanneler {
	return &FakeChanneler{
		ChannelHook: func(chan int) (ident2 chan int) {
			panic("Unexpected call to Channeler.Channel")
		},
		ChannelReceiveHook: func(<-chan int) (ident2 <-chan int) {
			panic("Unexpected call to Channeler.ChannelReceive")
		},
		ChannelSendHook: func(chan<- int) (ident2 chan<- int) {
			panic("Unexpected call to Channeler.ChannelSend")
		},
		ChannelPointerHook: func(*chan int) (ident2 *chan int) {
			panic("Unexpected call to Channeler.ChannelPointer")
		},
		ChannelInterfaceHook: func(chan interface{}) (ident2 chan interface{}) {
			panic("Unexpected call to Channeler.ChannelInterface")
		},
	}
}

// NewFakeChannelerDefaultFatal returns an instance of FakeChanneler with all hooks configured to call t.Fatal
func NewFakeChannelerDefaultFatal(t_sym1 ChannelerTestingT) *FakeChanneler {
	return &FakeChanneler{
		ChannelHook: func(chan int) (ident2 chan int) {
			t_sym1.Fatal("Unexpected call to Channeler.Channel")
			return
		},
		ChannelReceiveHook: func(<-chan int) (ident2 <-chan int) {
			t_sym1.Fatal("Unexpected call to Channeler.ChannelReceive")
			return
		},
		ChannelSendHook: func(chan<- int) (ident2 chan<- int) {
			t_sym1.Fatal("Unexpected call to Channeler.ChannelSend")
			return
		},
		ChannelPointerHook: func(*chan int) (ident2 *chan int) {
			t_sym1.Fatal("Unexpected call to Channeler.ChannelPointer")
			return
		},
		ChannelInterfaceHook: func(chan interface{}) (ident2 chan interface{}) {
			t_sym1.Fatal("Unexpected call to Channeler.ChannelInterface")
			return
		},
	}
}

// NewFakeChannelerDefaultError returns an instance of FakeChanneler with all hooks configured to call t.Error
func NewFakeChannelerDefaultError(t_sym2 ChannelerTestingT) *FakeChanneler {
	return &FakeChanneler{
		ChannelHook: func(chan int) (ident2 chan int) {
			t_sym2.Error("Unexpected call to Channeler.Channel")
			return
		},
		ChannelReceiveHook: func(<-chan int) (ident2 <-chan int) {
			t_sym2.Error("Unexpected call to Channeler.ChannelReceive")
			return
		},
		ChannelSendHook: func(chan<- int) (ident2 chan<- int) {
			t_sym2.Error("Unexpected call to Channeler.ChannelSend")
			return
		},
		ChannelPointerHook: func(*chan int) (ident2 *chan int) {
			t_sym2.Error("Unexpected call to Channeler.ChannelPointer")
			return
		},
		ChannelInterfaceHook: func(chan interface{}) (ident2 chan interface{}) {
			t_sym2.Error("Unexpected call to Channeler.ChannelInterface")
			return
		},
	}
}

func (f *FakeChanneler) Reset() {
	f.ChannelCalls = []*ChannelerChannelInvocation{}
	f.ChannelReceiveCalls = []*ChannelerChannelReceiveInvocation{}
	f.ChannelSendCalls = []*ChannelerChannelSendInvocation{}
	f.ChannelPointerCalls = []*ChannelerChannelPointerInvocation{}
	f.ChannelInterfaceCalls = []*ChannelerChannelInterfaceInvocation{}
}

func (f_sym3 *FakeChanneler) Channel(ident1 chan int) (ident2 chan int) {
	if f_sym3.ChannelHook == nil {
		panic("Channeler.Channel() called but FakeChanneler.ChannelHook is nil")
	}

	invocation_sym3 := new(ChannelerChannelInvocation)
	f_sym3.ChannelCalls = append(f_sym3.ChannelCalls, invocation_sym3)

	invocation_sym3.Parameters.Ident1 = ident1

	ident2 = f_sym3.ChannelHook(ident1)

	invocation_sym3.Results.Ident2 = ident2

	return
}

// SetChannelStub configures Channeler.Channel to always return the given values
func (f_sym4 *FakeChanneler) SetChannelStub(ident2 chan int) {
	f_sym4.ChannelHook = func(chan int) chan int {
		return ident2
	}
}

// SetChannelInvocation configures Channeler.Channel to return the given results when called with the given parameters
// If no match is found for an invocation the result(s) of the fallback function are returned
func (f_sym5 *FakeChanneler) SetChannelInvocation(calls_sym5 []*ChannelerChannelInvocation, fallback_sym5 func() chan int) {
	f_sym5.ChannelHook = func(ident1 chan int) (ident2 chan int) {
		for _, call_sym5 := range calls_sym5 {
			if reflect.DeepEqual(call_sym5.Parameters.Ident1, ident1) {
				ident2 = call_sym5.Results.Ident2

				return
			}
		}

		return fallback_sym5()
	}
}

// ChannelCalled returns true if FakeChanneler.Channel was called
func (f *FakeChanneler) ChannelCalled() bool {
	return len(f.ChannelCalls) != 0
}

// AssertChannelCalled calls t.Error if FakeChanneler.Channel was not called
func (f *FakeChanneler) AssertChannelCalled(t ChannelerTestingT) {
	t.Helper()
	if len(f.ChannelCalls) == 0 {
		t.Error("FakeChanneler.Channel not called, expected at least one")
	}
}

// ChannelNotCalled returns true if FakeChanneler.Channel was not called
func (f *FakeChanneler) ChannelNotCalled() bool {
	return len(f.ChannelCalls) == 0
}

// AssertChannelNotCalled calls t.Error if FakeChanneler.Channel was called
func (f *FakeChanneler) AssertChannelNotCalled(t ChannelerTestingT) {
	t.Helper()
	if len(f.ChannelCalls) != 0 {
		t.Error("FakeChanneler.Channel called, expected none")
	}
}

// ChannelCalledOnce returns true if FakeChanneler.Channel was called exactly once
func (f *FakeChanneler) ChannelCalledOnce() bool {
	return len(f.ChannelCalls) == 1
}

// AssertChannelCalledOnce calls t.Error if FakeChanneler.Channel was not called exactly once
func (f *FakeChanneler) AssertChannelCalledOnce(t ChannelerTestingT) {
	t.Helper()
	if len(f.ChannelCalls) != 1 {
		t.Errorf("FakeChanneler.Channel called %d times, expected 1", len(f.ChannelCalls))
	}
}

// ChannelCalledN returns true if FakeChanneler.Channel was called at least n times
func (f *FakeChanneler) ChannelCalledN(n int) bool {
	return len(f.ChannelCalls) >= n
}

// AssertChannelCalledN calls t.Error if FakeChanneler.Channel was called less than n times
func (f *FakeChanneler) AssertChannelCalledN(t ChannelerTestingT, n int) {
	t.Helper()
	if len(f.ChannelCalls) < n {
		t.Errorf("FakeChanneler.Channel called %d times, expected >= %d", len(f.ChannelCalls), n)
	}
}

// ChannelCalledWith returns true if FakeChanneler.Channel was called with the given values
func (f_sym6 *FakeChanneler) ChannelCalledWith(ident1 chan int) bool {
	for _, call_sym6 := range f_sym6.ChannelCalls {
		if reflect.DeepEqual(call_sym6.Parameters.Ident1, ident1) {
			return true
		}
	}

	return false
}

// AssertChannelCalledWith calls t.Error if FakeChanneler.Channel was not called with the given values
func (f_sym7 *FakeChanneler) AssertChannelCalledWith(t ChannelerTestingT, ident1 chan int) {
	t.Helper()
	var found_sym7 bool
	for _, call_sym7 := range f_sym7.ChannelCalls {
		if reflect.DeepEqual(call_sym7.Parameters.Ident1, ident1) {
			found_sym7 = true
			break
		}
	}

	if !found_sym7 {
		t.Error("FakeChanneler.Channel not called with expected parameters")
	}
}

// ChannelCalledOnceWith returns true if FakeChanneler.Channel was called exactly once with the given values
func (f_sym8 *FakeChanneler) ChannelCalledOnceWith(ident1 chan int) bool {
	var count_sym8 int
	for _, call_sym8 := range f_sym8.ChannelCalls {
		if reflect.DeepEqual(call_sym8.Parameters.Ident1, ident1) {
			count_sym8++
		}
	}

	return count_sym8 == 1
}

// AssertChannelCalledOnceWith calls t.Error if FakeChanneler.Channel was not called exactly once with the given values
func (f_sym9 *FakeChanneler) AssertChannelCalledOnceWith(t ChannelerTestingT, ident1 chan int) {
	t.Helper()
	var count_sym9 int
	for _, call_sym9 := range f_sym9.ChannelCalls {
		if reflect.DeepEqual(call_sym9.Parameters.Ident1, ident1) {
			count_sym9++
		}
	}

	if count_sym9 != 1 {
		t.Errorf("FakeChanneler.Channel called %d times with expected parameters, expected one", count_sym9)
	}
}

// ChannelResultsForCall returns the result values for the first call to FakeChanneler.Channel with the given values
func (f_sym10 *FakeChanneler) ChannelResultsForCall(ident1 chan int) (ident2 chan int, found_sym10 bool) {
	for _, call_sym10 := range f_sym10.ChannelCalls {
		if reflect.DeepEqual(call_sym10.Parameters.Ident1, ident1) {
			ident2 = call_sym10.Results.Ident2
			found_sym10 = true
			break
		}
	}

	return
}

func (f_sym11 *FakeChanneler) ChannelReceive(ident1 <-chan int) (ident2 <-chan int) {
	if f_sym11.ChannelReceiveHook == nil {
		panic("Channeler.ChannelReceive() called but FakeChanneler.ChannelReceiveHook is nil")
	}

	invocation_sym11 := new(ChannelerChannelReceiveInvocation)
	f_sym11.ChannelReceiveCalls = append(f_sym11.ChannelReceiveCalls, invocation_sym11)

	invocation_sym11.Parameters.Ident1 = ident1

	ident2 = f_sym11.ChannelReceiveHook(ident1)

	invocation_sym11.Results.Ident2 = ident2

	return
}

// SetChannelReceiveStub configures Channeler.ChannelReceive to always return the given values
func (f_sym12 *FakeChanneler) SetChannelReceiveStub(ident2 <-chan int) {
	f_sym12.ChannelReceiveHook = func(<-chan int) <-chan int {
		return ident2
	}
}

// SetChannelReceiveInvocation configures Channeler.ChannelReceive to return the given results when called with the given parameters
// If no match is found for an invocation the result(s) of the fallback function are returned
func (f_sym13 *FakeChanneler) SetChannelReceiveInvocation(calls_sym13 []*ChannelerChannelReceiveInvocation, fallback_sym13 func() <-chan int) {
	f_sym13.ChannelReceiveHook = func(ident1 <-chan int) (ident2 <-chan int) {
		for _, call_sym13 := range calls_sym13 {
			if reflect.DeepEqual(call_sym13.Parameters.Ident1, ident1) {
				ident2 = call_sym13.Results.Ident2

				return
			}
		}

		return fallback_sym13()
	}
}

// ChannelReceiveCalled returns true if FakeChanneler.ChannelReceive was called
func (f *FakeChanneler) ChannelReceiveCalled() bool {
	return len(f.ChannelReceiveCalls) != 0
}

// AssertChannelReceiveCalled calls t.Error if FakeChanneler.ChannelReceive was not called
func (f *FakeChanneler) AssertChannelReceiveCalled(t ChannelerTestingT) {
	t.Helper()
	if len(f.ChannelReceiveCalls) == 0 {
		t.Error("FakeChanneler.ChannelReceive not called, expected at least one")
	}
}

// ChannelReceiveNotCalled returns true if FakeChanneler.ChannelReceive was not called
func (f *FakeChanneler) ChannelReceiveNotCalled() bool {
	return len(f.ChannelReceiveCalls) == 0
}

// AssertChannelReceiveNotCalled calls t.Error if FakeChanneler.ChannelReceive was called
func (f *FakeChanneler) AssertChannelReceiveNotCalled(t ChannelerTestingT) {
	t.Helper()
	if len(f.ChannelReceiveCalls) != 0 {
		t.Error("FakeChanneler.ChannelReceive called, expected none")
	}
}

// ChannelReceiveCalledOnce returns true if FakeChanneler.ChannelReceive was called exactly once
func (f *FakeChanneler) ChannelReceiveCalledOnce() bool {
	return len(f.ChannelReceiveCalls) == 1
}

// AssertChannelReceiveCalledOnce calls t.Error if FakeChanneler.ChannelReceive was not called exactly once
func (f *FakeChanneler) AssertChannelReceiveCalledOnce(t ChannelerTestingT) {
	t.Helper()
	if len(f.ChannelReceiveCalls) != 1 {
		t.Errorf("FakeChanneler.ChannelReceive called %d times, expected 1", len(f.ChannelReceiveCalls))
	}
}

// ChannelReceiveCalledN returns true if FakeChanneler.ChannelReceive was called at least n times
func (f *FakeChanneler) ChannelReceiveCalledN(n int) bool {
	return len(f.ChannelReceiveCalls) >= n
}

// AssertChannelReceiveCalledN calls t.Error if FakeChanneler.ChannelReceive was called less than n times
func (f *FakeChanneler) AssertChannelReceiveCalledN(t ChannelerTestingT, n int) {
	t.Helper()
	if len(f.ChannelReceiveCalls) < n {
		t.Errorf("FakeChanneler.ChannelReceive called %d times, expected >= %d", len(f.ChannelReceiveCalls), n)
	}
}

// ChannelReceiveCalledWith returns true if FakeChanneler.ChannelReceive was called with the given values
func (f_sym14 *FakeChanneler) ChannelReceiveCalledWith(ident1 <-chan int) bool {
	for _, call_sym14 := range f_sym14.ChannelReceiveCalls {
		if reflect.DeepEqual(call_sym14.Parameters.Ident1, ident1) {
			return true
		}
	}

	return false
}

// AssertChannelReceiveCalledWith calls t.Error if FakeChanneler.ChannelReceive was not called with the given values
func (f_sym15 *FakeChanneler) AssertChannelReceiveCalledWith(t ChannelerTestingT, ident1 <-chan int) {
	t.Helper()
	var found_sym15 bool
	for _, call_sym15 := range f_sym15.ChannelReceiveCalls {
		if reflect.DeepEqual(call_sym15.Parameters.Ident1, ident1) {
			found_sym15 = true
			break
		}
	}

	if !found_sym15 {
		t.Error("FakeChanneler.ChannelReceive not called with expected parameters")
	}
}

// ChannelReceiveCalledOnceWith returns true if FakeChanneler.ChannelReceive was called exactly once with the given values
func (f_sym16 *FakeChanneler) ChannelReceiveCalledOnceWith(ident1 <-chan int) bool {
	var count_sym16 int
	for _, call_sym16 := range f_sym16.ChannelReceiveCalls {
		if reflect.DeepEqual(call_sym16.Parameters.Ident1, ident1) {
			count_sym16++
		}
	}

	return count_sym16 == 1
}

// AssertChannelReceiveCalledOnceWith calls t.Error if FakeChanneler.ChannelReceive was not called exactly once with the given values
func (f_sym17 *FakeChanneler) AssertChannelReceiveCalledOnceWith(t ChannelerTestingT, ident1 <-chan int) {
	t.Helper()
	var count_sym17 int
	for _, call_sym17 := range f_sym17.ChannelReceiveCalls {
		if reflect.DeepEqual(call_sym17.Parameters.Ident1, ident1) {
			count_sym17++
		}
	}

	if count_sym17 != 1 {
		t.Errorf("FakeChanneler.ChannelReceive called %d times with expected parameters, expected one", count_sym17)
	}
}

// ChannelReceiveResultsForCall returns the result values for the first call to FakeChanneler.ChannelReceive with the given values
func (f_sym18 *FakeChanneler) ChannelReceiveResultsForCall(ident1 <-chan int) (ident2 <-chan int, found_sym18 bool) {
	for _, call_sym18 := range f_sym18.ChannelReceiveCalls {
		if reflect.DeepEqual(call_sym18.Parameters.Ident1, ident1) {
			ident2 = call_sym18.Results.Ident2
			found_sym18 = true
			break
		}
	}

	return
}

func (f_sym19 *FakeChanneler) ChannelSend(ident1 chan<- int) (ident2 chan<- int) {
	if f_sym19.ChannelSendHook == nil {
		panic("Channeler.ChannelSend() called but FakeChanneler.ChannelSendHook is nil")
	}

	invocation_sym19 := new(ChannelerChannelSendInvocation)
	f_sym19.ChannelSendCalls = append(f_sym19.ChannelSendCalls, invocation_sym19)

	invocation_sym19.Parameters.Ident1 = ident1

	ident2 = f_sym19.ChannelSendHook(ident1)

	invocation_sym19.Results.Ident2 = ident2

	return
}

// SetChannelSendStub configures Channeler.ChannelSend to always return the given values
func (f_sym20 *FakeChanneler) SetChannelSendStub(ident2 chan<- int) {
	f_sym20.ChannelSendHook = func(chan<- int) chan<- int {
		return ident2
	}
}

// SetChannelSendInvocation configures Channeler.ChannelSend to return the given results when called with the given parameters
// If no match is found for an invocation the result(s) of the fallback function are returned
func (f_sym21 *FakeChanneler) SetChannelSendInvocation(calls_sym21 []*ChannelerChannelSendInvocation, fallback_sym21 func() chan<- int) {
	f_sym21.ChannelSendHook = func(ident1 chan<- int) (ident2 chan<- int) {
		for _, call_sym21 := range calls_sym21 {
			if reflect.DeepEqual(call_sym21.Parameters.Ident1, ident1) {
				ident2 = call_sym21.Results.Ident2

				return
			}
		}

		return fallback_sym21()
	}
}

// ChannelSendCalled returns true if FakeChanneler.ChannelSend was called
func (f *FakeChanneler) ChannelSendCalled() bool {
	return len(f.ChannelSendCalls) != 0
}

// AssertChannelSendCalled calls t.Error if FakeChanneler.ChannelSend was not called
func (f *FakeChanneler) AssertChannelSendCalled(t ChannelerTestingT) {
	t.Helper()
	if len(f.ChannelSendCalls) == 0 {
		t.Error("FakeChanneler.ChannelSend not called, expected at least one")
	}
}

// ChannelSendNotCalled returns true if FakeChanneler.ChannelSend was not called
func (f *FakeChanneler) ChannelSendNotCalled() bool {
	return len(f.ChannelSendCalls) == 0
}

// AssertChannelSendNotCalled calls t.Error if FakeChanneler.ChannelSend was called
func (f *FakeChanneler) AssertChannelSendNotCalled(t ChannelerTestingT) {
	t.Helper()
	if len(f.ChannelSendCalls) != 0 {
		t.Error("FakeChanneler.ChannelSend called, expected none")
	}
}

// ChannelSendCalledOnce returns true if FakeChanneler.ChannelSend was called exactly once
func (f *FakeChanneler) ChannelSendCalledOnce() bool {
	return len(f.ChannelSendCalls) == 1
}

// AssertChannelSendCalledOnce calls t.Error if FakeChanneler.ChannelSend was not called exactly once
func (f *FakeChanneler) AssertChannelSendCalledOnce(t ChannelerTestingT) {
	t.Helper()
	if len(f.ChannelSendCalls) != 1 {
		t.Errorf("FakeChanneler.ChannelSend called %d times, expected 1", len(f.ChannelSendCalls))
	}
}

// ChannelSendCalledN returns true if FakeChanneler.ChannelSend was called at least n times
func (f *FakeChanneler) ChannelSendCalledN(n int) bool {
	return len(f.ChannelSendCalls) >= n
}

// AssertChannelSendCalledN calls t.Error if FakeChanneler.ChannelSend was called less than n times
func (f *FakeChanneler) AssertChannelSendCalledN(t ChannelerTestingT, n int) {
	t.Helper()
	if len(f.ChannelSendCalls) < n {
		t.Errorf("FakeChanneler.ChannelSend called %d times, expected >= %d", len(f.ChannelSendCalls), n)
	}
}

// ChannelSendCalledWith returns true if FakeChanneler.ChannelSend was called with the given values
func (f_sym22 *FakeChanneler) ChannelSendCalledWith(ident1 chan<- int) bool {
	for _, call_sym22 := range f_sym22.ChannelSendCalls {
		if reflect.DeepEqual(call_sym22.Parameters.Ident1, ident1) {
			return true
		}
	}

	return false
}

// AssertChannelSendCalledWith calls t.Error if FakeChanneler.ChannelSend was not called with the given values
func (f_sym23 *FakeChanneler) AssertChannelSendCalledWith(t ChannelerTestingT, ident1 chan<- int) {
	t.Helper()
	var found_sym23 bool
	for _, call_sym23 := range f_sym23.ChannelSendCalls {
		if reflect.DeepEqual(call_sym23.Parameters.Ident1, ident1) {
			found_sym23 = true
			break
		}
	}

	if !found_sym23 {
		t.Error("FakeChanneler.ChannelSend not called with expected parameters")
	}
}

// ChannelSendCalledOnceWith returns true if FakeChanneler.ChannelSend was called exactly once with the given values
func (f_sym24 *FakeChanneler) ChannelSendCalledOnceWith(ident1 chan<- int) bool {
	var count_sym24 int
	for _, call_sym24 := range f_sym24.ChannelSendCalls {
		if reflect.DeepEqual(call_sym24.Parameters.Ident1, ident1) {
			count_sym24++
		}
	}

	return count_sym24 == 1
}

// AssertChannelSendCalledOnceWith calls t.Error if FakeChanneler.ChannelSend was not called exactly once with the given values
func (f_sym25 *FakeChanneler) AssertChannelSendCalledOnceWith(t ChannelerTestingT, ident1 chan<- int) {
	t.Helper()
	var count_sym25 int
	for _, call_sym25 := range f_sym25.ChannelSendCalls {
		if reflect.DeepEqual(call_sym25.Parameters.Ident1, ident1) {
			count_sym25++
		}
	}

	if count_sym25 != 1 {
		t.Errorf("FakeChanneler.ChannelSend called %d times with expected parameters, expected one", count_sym25)
	}
}

// ChannelSendResultsForCall returns the result values for the first call to FakeChanneler.ChannelSend with the given values
func (f_sym26 *FakeChanneler) ChannelSendResultsForCall(ident1 chan<- int) (ident2 chan<- int, found_sym26 bool) {
	for _, call_sym26 := range f_sym26.ChannelSendCalls {
		if reflect.DeepEqual(call_sym26.Parameters.Ident1, ident1) {
			ident2 = call_sym26.Results.Ident2
			found_sym26 = true
			break
		}
	}

	return
}

func (f_sym27 *FakeChanneler) ChannelPointer(ident1 *chan int) (ident2 *chan int) {
	if f_sym27.ChannelPointerHook == nil {
		panic("Channeler.ChannelPointer() called but FakeChanneler.ChannelPointerHook is nil")
	}

	invocation_sym27 := new(ChannelerChannelPointerInvocation)
	f_sym27.ChannelPointerCalls = append(f_sym27.ChannelPointerCalls, invocation_sym27)

	invocation_sym27.Parameters.Ident1 = ident1

	ident2 = f_sym27.ChannelPointerHook(ident1)

	invocation_sym27.Results.Ident2 = ident2

	return
}

// SetChannelPointerStub configures Channeler.ChannelPointer to always return the given values
func (f_sym28 *FakeChanneler) SetChannelPointerStub(ident2 *chan int) {
	f_sym28.ChannelPointerHook = func(*chan int) *chan int {
		return ident2
	}
}

// SetChannelPointerInvocation configures Channeler.ChannelPointer to return the given results when called with the given parameters
// If no match is found for an invocation the result(s) of the fallback function are returned
func (f_sym29 *FakeChanneler) SetChannelPointerInvocation(calls_sym29 []*ChannelerChannelPointerInvocation, fallback_sym29 func() *chan int) {
	f_sym29.ChannelPointerHook = func(ident1 *chan int) (ident2 *chan int) {
		for _, call_sym29 := range calls_sym29 {
			if reflect.DeepEqual(call_sym29.Parameters.Ident1, ident1) {
				ident2 = call_sym29.Results.Ident2

				return
			}
		}

		return fallback_sym29()
	}
}

// ChannelPointerCalled returns true if FakeChanneler.ChannelPointer was called
func (f *FakeChanneler) ChannelPointerCalled() bool {
	return len(f.ChannelPointerCalls) != 0
}

// AssertChannelPointerCalled calls t.Error if FakeChanneler.ChannelPointer was not called
func (f *FakeChanneler) AssertChannelPointerCalled(t ChannelerTestingT) {
	t.Helper()
	if len(f.ChannelPointerCalls) == 0 {
		t.Error("FakeChanneler.ChannelPointer not called, expected at least one")
	}
}

// ChannelPointerNotCalled returns true if FakeChanneler.ChannelPointer was not called
func (f *FakeChanneler) ChannelPointerNotCalled() bool {
	return len(f.ChannelPointerCalls) == 0
}

// AssertChannelPointerNotCalled calls t.Error if FakeChanneler.ChannelPointer was called
func (f *FakeChanneler) AssertChannelPointerNotCalled(t ChannelerTestingT) {
	t.Helper()
	if len(f.ChannelPointerCalls) != 0 {
		t.Error("FakeChanneler.ChannelPointer called, expected none")
	}
}

// ChannelPointerCalledOnce returns true if FakeChanneler.ChannelPointer was called exactly once
func (f *FakeChanneler) ChannelPointerCalledOnce() bool {
	return len(f.ChannelPointerCalls) == 1
}

// AssertChannelPointerCalledOnce calls t.Error if FakeChanneler.ChannelPointer was not called exactly once
func (f *FakeChanneler) AssertChannelPointerCalledOnce(t ChannelerTestingT) {
	t.Helper()
	if len(f.ChannelPointerCalls) != 1 {
		t.Errorf("FakeChanneler.ChannelPointer called %d times, expected 1", len(f.ChannelPointerCalls))
	}
}

// ChannelPointerCalledN returns true if FakeChanneler.ChannelPointer was called at least n times
func (f *FakeChanneler) ChannelPointerCalledN(n int) bool {
	return len(f.ChannelPointerCalls) >= n
}

// AssertChannelPointerCalledN calls t.Error if FakeChanneler.ChannelPointer was called less than n times
func (f *FakeChanneler) AssertChannelPointerCalledN(t ChannelerTestingT, n int) {
	t.Helper()
	if len(f.ChannelPointerCalls) < n {
		t.Errorf("FakeChanneler.ChannelPointer called %d times, expected >= %d", len(f.ChannelPointerCalls), n)
	}
}

// ChannelPointerCalledWith returns true if FakeChanneler.ChannelPointer was called with the given values
func (f_sym30 *FakeChanneler) ChannelPointerCalledWith(ident1 *chan int) bool {
	for _, call_sym30 := range f_sym30.ChannelPointerCalls {
		if reflect.DeepEqual(call_sym30.Parameters.Ident1, ident1) {
			return true
		}
	}

	return false
}

// AssertChannelPointerCalledWith calls t.Error if FakeChanneler.ChannelPointer was not called with the given values
func (f_sym31 *FakeChanneler) AssertChannelPointerCalledWith(t ChannelerTestingT, ident1 *chan int) {
	t.Helper()
	var found_sym31 bool
	for _, call_sym31 := range f_sym31.ChannelPointerCalls {
		if reflect.DeepEqual(call_sym31.Parameters.Ident1, ident1) {
			found_sym31 = true
			break
		}
	}

	if !found_sym31 {
		t.Error("FakeChanneler.ChannelPointer not called with expected parameters")
	}
}

// ChannelPointerCalledOnceWith returns true if FakeChanneler.ChannelPointer was called exactly once with the given values
func (f_sym32 *FakeChanneler) ChannelPointerCalledOnceWith(ident1 *chan int) bool {
	var count_sym32 int
	for _, call_sym32 := range f_sym32.ChannelPointerCalls {
		if reflect.DeepEqual(call_sym32.Parameters.Ident1, ident1) {
			count_sym32++
		}
	}

	return count_sym32 == 1
}

// AssertChannelPointerCalledOnceWith calls t.Error if FakeChanneler.ChannelPointer was not called exactly once with the given values
func (f_sym33 *FakeChanneler) AssertChannelPointerCalledOnceWith(t ChannelerTestingT, ident1 *chan int) {
	t.Helper()
	var count_sym33 int
	for _, call_sym33 := range f_sym33.ChannelPointerCalls {
		if reflect.DeepEqual(call_sym33.Parameters.Ident1, ident1) {
			count_sym33++
		}
	}

	if count_sym33 != 1 {
		t.Errorf("FakeChanneler.ChannelPointer called %d times with expected parameters, expected one", count_sym33)
	}
}

// ChannelPointerResultsForCall returns the result values for the first call to FakeChanneler.ChannelPointer with the given values
func (f_sym34 *FakeChanneler) ChannelPointerResultsForCall(ident1 *chan int) (ident2 *chan int, found_sym34 bool) {
	for _, call_sym34 := range f_sym34.ChannelPointerCalls {
		if reflect.DeepEqual(call_sym34.Parameters.Ident1, ident1) {
			ident2 = call_sym34.Results.Ident2
			found_sym34 = true
			break
		}
	}

	return
}

func (f_sym35 *FakeChanneler) ChannelInterface(ident1 chan interface{}) (ident2 chan interface{}) {
	if f_sym35.ChannelInterfaceHook == nil {
		panic("Channeler.ChannelInterface() called but FakeChanneler.ChannelInterfaceHook is nil")
	}

	invocation_sym35 := new(ChannelerChannelInterfaceInvocation)
	f_sym35.ChannelInterfaceCalls = append(f_sym35.ChannelInterfaceCalls, invocation_sym35)

	invocation_sym35.Parameters.Ident1 = ident1

	ident2 = f_sym35.ChannelInterfaceHook(ident1)

	invocation_sym35.Results.Ident2 = ident2

	return
}

// SetChannelInterfaceStub configures Channeler.ChannelInterface to always return the given values
func (f_sym36 *FakeChanneler) SetChannelInterfaceStub(ident2 chan interface{}) {
	f_sym36.ChannelInterfaceHook = func(chan interface{}) chan interface{} {
		return ident2
	}
}

// SetChannelInterfaceInvocation configures Channeler.ChannelInterface to return the given results when called with the given parameters
// If no match is found for an invocation the result(s) of the fallback function are returned
func (f_sym37 *FakeChanneler) SetChannelInterfaceInvocation(calls_sym37 []*ChannelerChannelInterfaceInvocation, fallback_sym37 func() chan interface{}) {
	f_sym37.ChannelInterfaceHook = func(ident1 chan interface{}) (ident2 chan interface{}) {
		for _, call_sym37 := range calls_sym37 {
			if reflect.DeepEqual(call_sym37.Parameters.Ident1, ident1) {
				ident2 = call_sym37.Results.Ident2

				return
			}
		}

		return fallback_sym37()
	}
}

// ChannelInterfaceCalled returns true if FakeChanneler.ChannelInterface was called
func (f *FakeChanneler) ChannelInterfaceCalled() bool {
	return len(f.ChannelInterfaceCalls) != 0
}

// AssertChannelInterfaceCalled calls t.Error if FakeChanneler.ChannelInterface was not called
func (f *FakeChanneler) AssertChannelInterfaceCalled(t ChannelerTestingT) {
	t.Helper()
	if len(f.ChannelInterfaceCalls) == 0 {
		t.Error("FakeChanneler.ChannelInterface not called, expected at least one")
	}
}

// ChannelInterfaceNotCalled returns true if FakeChanneler.ChannelInterface was not called
func (f *FakeChanneler) ChannelInterfaceNotCalled() bool {
	return len(f.ChannelInterfaceCalls) == 0
}

// AssertChannelInterfaceNotCalled calls t.Error if FakeChanneler.ChannelInterface was called
func (f *FakeChanneler) AssertChannelInterfaceNotCalled(t ChannelerTestingT) {
	t.Helper()
	if len(f.ChannelInterfaceCalls) != 0 {
		t.Error("FakeChanneler.ChannelInterface called, expected none")
	}
}

// ChannelInterfaceCalledOnce returns true if FakeChanneler.ChannelInterface was called exactly once
func (f *FakeChanneler) ChannelInterfaceCalledOnce() bool {
	return len(f.ChannelInterfaceCalls) == 1
}

// AssertChannelInterfaceCalledOnce calls t.Error if FakeChanneler.ChannelInterface was not called exactly once
func (f *FakeChanneler) AssertChannelInterfaceCalledOnce(t ChannelerTestingT) {
	t.Helper()
	if len(f.ChannelInterfaceCalls) != 1 {
		t.Errorf("FakeChanneler.ChannelInterface called %d times, expected 1", len(f.ChannelInterfaceCalls))
	}
}

// ChannelInterfaceCalledN returns true if FakeChanneler.ChannelInterface was called at least n times
func (f *FakeChanneler) ChannelInterfaceCalledN(n int) bool {
	return len(f.ChannelInterfaceCalls) >= n
}

// AssertChannelInterfaceCalledN calls t.Error if FakeChanneler.ChannelInterface was called less than n times
func (f *FakeChanneler) AssertChannelInterfaceCalledN(t ChannelerTestingT, n int) {
	t.Helper()
	if len(f.ChannelInterfaceCalls) < n {
		t.Errorf("FakeChanneler.ChannelInterface called %d times, expected >= %d", len(f.ChannelInterfaceCalls), n)
	}
}

// ChannelInterfaceCalledWith returns true if FakeChanneler.ChannelInterface was called with the given values
func (f_sym38 *FakeChanneler) ChannelInterfaceCalledWith(ident1 chan interface{}) bool {
	for _, call_sym38 := range f_sym38.ChannelInterfaceCalls {
		if reflect.DeepEqual(call_sym38.Parameters.Ident1, ident1) {
			return true
		}
	}

	return false
}

// AssertChannelInterfaceCalledWith calls t.Error if FakeChanneler.ChannelInterface was not called with the given values
func (f_sym39 *FakeChanneler) AssertChannelInterfaceCalledWith(t ChannelerTestingT, ident1 chan interface{}) {
	t.Helper()
	var found_sym39 bool
	for _, call_sym39 := range f_sym39.ChannelInterfaceCalls {
		if reflect.DeepEqual(call_sym39.Parameters.Ident1, ident1) {
			found_sym39 = true
			break
		}
	}

	if !found_sym39 {
		t.Error("FakeChanneler.ChannelInterface not called with expected parameters")
	}
}

// ChannelInterfaceCalledOnceWith returns true if FakeChanneler.ChannelInterface was called exactly once with the given values
func (f_sym40 *FakeChanneler) ChannelInterfaceCalledOnceWith(ident1 chan interface{}) bool {
	var count_sym40 int
	for _, call_sym40 := range f_sym40.ChannelInterfaceCalls {
		if reflect.DeepEqual(call_sym40.Parameters.Ident1, ident1) {
			count_sym40++
		}
	}

	return count_sym40 == 1
}

// AssertChannelInterfaceCalledOnceWith calls t.Error if FakeChanneler.ChannelInterface was not called exactly once with the given values
func (f_sym41 *FakeChanneler) AssertChannelInterfaceCalledOnceWith(t ChannelerTestingT, ident1 chan interface{}) {
	t.Helper()
	var count_sym41 int
	for _, call_sym41 := range f_sym41.ChannelInterfaceCalls {
		if reflect.DeepEqual(call_sym41.Parameters.Ident1, ident1) {
			count_sym41++
		}
	}

	if count_sym41 != 1 {
		t.Errorf("FakeChanneler.ChannelInterface called %d times with expected parameters, expected one", count_sym41)
	}
}

// ChannelInterfaceResultsForCall returns the result values for the first call to FakeChanneler.ChannelInterface with the given values
func (f_sym42 *FakeChanneler) ChannelInterfaceResultsForCall(ident1 chan interface{}) (ident2 chan interface{}, found_sym42 bool) {
	for _, call_sym42 := range f_sym42.ChannelInterfaceCalls {
		if reflect.DeepEqual(call_sym42.Parameters.Ident1, ident1) {
			ident2 = call_sym42.Results.Ident2
			found_sym42 = true
			break
		}
	}

	return
}
