package queueexample

import (
	"fmt"
	"github.com/Quasar777/go-algo/data_structures/queue"
)

func main() {
	qu := &queue.Queue{
		Nums: []int{},
	}

	qu.Enque(1)
	qu.Enque(2)
	qu.Enque(3)

	val, err := qu.Deque()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("dequed:", val)
}