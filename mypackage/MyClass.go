package mypackage

type MyClass struct {
	x int
}

func (this MyClass) X() int {
	return this.x
}

func (this *MyClass) SetX(x int) {
	this.x = x
}