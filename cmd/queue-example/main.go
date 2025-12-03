package queueexample

import (
	"fmt"
	"github.com/Quasar777/go-algo/data_structures/queue"
)

func main() {
	qu := &queue.Queue[int]{
		Nums: []int{},
	}

	qu.Enque(1)
	qu.Enque(2)
	qu.Enque(3)

	val, ok := qu.Deque()
	if !ok {
		fmt.Println("no values")
		return
	}

	fmt.Println("dequed:", val)
}