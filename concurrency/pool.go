package concurrency

import (
	"fmt"
	"sync"
)

type Person struct {
	Name string
	Age  int
}

func NewPerson() interface{} {
	return &Person{}
}

func PoolDemo() {
	pool := sync.Pool{
		New: NewPerson,
	}

	person := pool.Get().(*Person)
	person.Name = "John"
	person.Age = 30
	fmt.Printf("Person 1: %+v\n", person)

	pool.Put(person)

	person = pool.Get().(*Person)
	fmt.Printf("Person 2: %+v\n", person)

	pool.Put(person)

	person = pool.Get().(*Person)
	fmt.Printf("Person 3: %+v\n", person)
}
