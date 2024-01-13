package pulsepropagation

type Queue []Event // FIFO queue

func (q *Queue) Push(e Event) {
	*q = append(*q, e)
}

func (q *Queue) Pop() Event {
	e := (*q)[0]
	*q = (*q)[1:]
	return e
}

func (q *Queue) Empty() bool {
	return len(*q) == 0
}
