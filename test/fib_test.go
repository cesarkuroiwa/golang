package test

import "testing"
import . "launchpad.net/gocheck"

func TestFibonacci(t *testing.T) {
	r := Fibonacci(5)
	e := 5
	if (r != e) {
		t.Errorf("Expected %d, got %d", 5, e)
	}
}

func Test(t *testing.T) { TestingT(t) }

type S struct {}
var _ = Suite(&S{})

func (s *S) TestFibonacci2(c *C) {
	c.Assert(Fibonacci(1), Equals, 1)
	c.Check(Fibonacci(2), Equals, 1)
}

func (s *S) TestFibonacci3(c *C) {
	c.Check(Fibonacci(5), Equals, 5)
}