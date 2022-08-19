package list_test

import (
	"fmt"

	"github.com/zendesk/generic/list"
)

func Example() {
	l := list.New[int]()
	l.PushBack(0)
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	l.Front.Each(func(i int) {
		fmt.Println(i)
	})
	// Output:
	// 0
	// 1
	// 2
	// 3
}
