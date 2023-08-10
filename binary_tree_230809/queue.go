package main

type Queue []interface{}

func MakeQueue() *Queue {
	return &Queue{}
}

func (q Queue) Len() int {
	return len(q)
}

func (q Queue) IsEmpty() bool {
	return len(q) == 0
}

func (q *Queue) EnQueue(item interface{}) {
	*q = append(append(*q, item))
}

func (q *Queue) DeQueue() interface{} {
	if (*q).Len() == 0 {
		return nil
	}
	item := (*q)[0]
	*q = (*q)[1:len(*q)]
	return item
}
