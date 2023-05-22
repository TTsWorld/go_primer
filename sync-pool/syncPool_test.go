package sync_pool

import (
	"fmt"
	"sync"
	"testing"
)

var pool *sync.Pool

type Person struct {
	Name string
}

func init() {
	pool = &sync.Pool{
		New: func() interface{} {
			fmt.Println("creating a new person")
			return new(Person)
		},
	}
}

func TestSyncPool(t *testing.T) {

	person := pool.Get().(*Person)
	fmt.Println("Get Pool Object：", person)

	person.Name = "first"
	pool.Put(person)

	fmt.Println("Get Pool Object：", pool.Get().(*Person))
	fmt.Println("Get Pool Object：", pool.Get().(*Person))

}
