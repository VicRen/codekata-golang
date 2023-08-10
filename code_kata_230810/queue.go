package main

type Queue []interface{}

func MakeQueue() *Queue {
	return &Queue{}
}

func (q Queue) IsEmpty() bool {
	return len(q) == 0
}

func (q Queue) Len() int {
	return len(q)
}

func (q *Queue) EnQueue(item interface{}) {
	*q = append(*q, item)
}

func (q *Queue) DeQueue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	ret := (*q)[0]
	*q = (*q)[1:]
	return ret
}
