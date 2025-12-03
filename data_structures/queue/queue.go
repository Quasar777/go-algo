package queue

type Queue[T any] struct {
	Nums []T
}

func (q *Queue[T]) Enque(val T) {
	q.Nums = append(q.Nums, val)
}

func (q *Queue[T]) Deque() (T, bool) {
	var zero T

	if len(q.Nums) == 0 {
		return zero, false 
	}

	val := q.Nums[0]
	q.Nums = q.Nums[1:]
	return val, true
}