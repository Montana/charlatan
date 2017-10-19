# Charlatan

[![Circle CI](https://circleci.com/gh/percolate/charlatan.svg?style=svg)](https://circleci.com/gh/percolate/charlatan)
[![codecov.io](https://codecov.io/github/percolate/charlatan/coverage.svg?branch=master)](https://codecov.io/github/percolate/charlatan?branch=master)

Percolate's Go Interface Mocking Tool.

## Installation

    go get github.com/percolate/charlatan

## Usage

```
  charlatan [options] <interface> ...
  charlatan -h | --help

Options:

  -dir string
        input package directory [default: current package directory]
  -file value
        name of input file, may be repeated, ignored if -dir is present
  -output string
        output file path [default: ./charlatan.go]
  -package string
        output package name [default: "<current package>"]
```

If you would like the mock implementations to live in the same package
as the interace definition then use the simplest invocation as a
directive:

    //go:generate charlatan Interface

or from the command line:

    charlatan -file=path/to/file.go Interface

You can chose the output path using `-output`, which must include the
name of the generated source file.  Any intermediate directories in the
path that don't exist will be created.  The package used in the
generated file's `pacakge` directive can be set using `-package`.

## Example

Given the following interface:

```go
package example

//go:generate charlatan Service

type Service interface {
	Query(filter *QueryFilter) ([]*Thing, error)
	Fetch(id string) (*Thing, error)
}
```

Running `go generate ...` for the above package/file should produce
the file `charlatan.go`:

```go
package example

type QueryInvocation struct {
	Parameters struct {
		Filter *QueryFilter
	}
	Results struct {
		Ident1 []*Thing
		Ident2 error
	}
}

type FetechInvocation struct {
	Parameters struct {
		Id string
	}
	Results struct {
		Ident3 *Thing
		Ident4 error
	}
}

type FakeService struct {
	QueryHook func(*QueryFilter) ([]*Thing, error)
	FetchHook func(string) (*Thing, error)

	QueryCalls []*QueryInvocation
	FetchCalls []*FetchInvocation
}

func (f *FakeService) Query(filter *QueryFilter) (ident1 []*Thing, ident2 error) {
	invocation := new(QueryInvocation)
	invocation.Parameters.Filter = filter

	ident1, ident2 := f.QueryHook(filter)

	invocation.Results.Ident1 = ident1
	invocation.Results.Ident2 = ident2

	return
}

// other generated code elided ...
```

Now you can use this in your tests by injecting the `FakeService`
implementation instead of the actual one.  A `FakeService` can be used
anywhere a `Service` interface is expected.

```go
func TestUsingService(t *testing.T) {
	// expectedThings := ...
	// expectedCriteria := ...
	svc := &example.FakeService{
		QueryHook: func(filter *QueryFilter) ([]*Thing, error) {
			if filter.Criteria != expectedCriteria {
				t.Errorf("expected criteria value: %v, have: %v", filter.Criteria, expectedCriteria)
				return nil, errors.New("unexpected criteria")
			}
			return expectedThings, nil
		}
	}

	// use the `svc` instance in the code under text ...

	// assert state of FakeService ...
	svc.AssertQueryCalledOnce(t)
}
```

Create anonymous function implementations for only those interface
methods that should be called in the code under test.  This will force
a painc if any unexpected calls are made the mock implementation.

The generated code has `godoc` formatted comments explaining the use
of the mock and its methods.
