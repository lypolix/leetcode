package medium

type FooBar struct {
	n int
    fooChan  chan struct{}
    barChan  chan struct{}
}

func NewFooBar(n int) *FooBar {
	fb := &FooBar{
        n: n,
        fooChan: make(chan struct{}, 1),
        barChan: make(chan struct{}, 1),
    }
    fb.fooChan <- struct{}{}
    return fb
}

func (fb *FooBar) Foo(printFoo func()) {
	for i := 0; i < fb.n; i++ {
		// printFoo() outputs "foo". Do not change or remove this line.

        <-fb.fooChan
        printFoo()
        fb.barChan <- struct{}{}
	}
}

func (fb *FooBar) Bar(printBar func()) {
	for i := 0; i < fb.n; i++ {
		// printBar() outputs "bar". Do not change or remove this line.

        <-fb.barChan
        printBar()
        fb.fooChan <- struct{}{}
	}
}

