package lazy_test

import (
	"errors"
	"fmt"

	"github.com/JeremyLoy/lazy"
)

type server struct {
	expensiveResource         func() string
	expensiveFailableResource func() (string, error)
}

func (s *server) foo() {
	resource := s.expensiveResource()
	_ = resource // use resource in foo
}

func (s *server) bar() error {
	resource, err := s.expensiveFailableResource()
	if err != nil {
		return err
	}
	_ = resource // use resource in foo
	return nil
}

func ExampleLazy() {
	var timesCalled int
	s := server{
		expensiveResource: lazy.Lazy(func() string {
			timesCalled++
			return "expensive"
		}),
	}

	s.foo()
	s.foo()

	fmt.Println(timesCalled)

	// Output:
	// 1
}

func ExampleLazyError() {
	var timesCalled int
	s := server{
		expensiveFailableResource: lazy.LazyError(func() (string, error) {
			timesCalled++
			return "", errors.New("Oops!")
		}),
	}

	err := s.bar()
	_ = s.bar()

	fmt.Println(timesCalled)
	fmt.Println(err)

	// Output:
	// 1
	// Oops!
}
