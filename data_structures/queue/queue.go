package queue

import "fmt"

type Queue struct {
	Nums []int
}

func (q *Queue) Enque(val int) {
	q.Nums = append(q.Nums, val)
}

func (q *Queue) Deque() (int, error) {
	if len(q.Nums) == 0 {
		return 0, fmt.Errorf("queue is empty")
	}

	val := q.Nums[0]
	q.Nums = q.Nums[1:]
	return val, nil
}