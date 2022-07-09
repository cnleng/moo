package main

import "fmt"

type Foo struct {
	s string
	i int
}

func foo(f Foo) {
	fmt.Printf("%p\n", &f.s)
}

func bar(f *Foo) {
	fmt.Printf("%p\n", &f.s)
}

func main() {
	f := Foo{"Hello", 100}
	fmt.Printf("%p\n", &f.s)
	foo(f)
	bar(&f)
}
