package main

type Queue struct {
	queue []string
	len   int
}

func NewQueue() *Queue {
	queue := &Queue{}
	queue.queue = make([]string, 0)
	queue.len = 0
	return queue
}

func (queue *Queue) Len() int {
	return queue.len
}
func (queue *Queue) IsEmpty() bool {
	return queue.len == 0
}

func (q *Queue) Dequeue() string {
	tmp := q.queue[0]
	q.queue = q.queue[1:]
	q.len -= 1
	return tmp
}

func (q *Queue) Enqueue(value string) {
	q.len += 1
	q.queue = append(q.queue, value)
}
