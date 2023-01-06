package main

import "testing"

type Queue struct {
	items chan int
}
type Queue1 interface {
	Push(key interface{})
	Pop() interface{}
	Contains(key interface{}) bool
	Len() int
	Keys() []interface{}
}

func New(size int) Queue1 {
	return nil
}
func (q *Queue) Enqueue(i int) {
	q.items <- i
}

func (q *Queue) Dequeue() int {
	return <-q.items
}

func main() {
	q := Queue{
		items: make(chan int, 16),
	}
	q.Enqueue(2)
	q.Enqueue(3)

	q.Enqueue(1)

	println(q.Dequeue())
	println(q.Dequeue())
	println(q.Dequeue())
}

var testValues = []interface{}{
	"lorem",
	"ipsum",
	1,
	2,
	3,
	"jack",
	"jill",
	"felix",
	"donking",
}

// TestPush validate evict old item policy
func TestEvictPolicy(t *testing.T) {
	size := 5
	q := New(size)

	for i, v := range testValues {
		q.Push(v)

		t.Log("current: ", q.Keys())

		// validate
		// item existence
		if !q.Contains(v) {
			t.Errorf("policy: newly inserted %v must be exists", v)
		}

		if i < 5 && q.Len() != (i+1) {
			t.Errorf("expected length %d but actual: %d", i+1, q.Len())
		} else if i >= 5 && q.Len() != 5 {
			t.Errorf("expexted length: %d but actual: %d", size, q.Len())
		}
	}
}

// TestPop validate pop item policy
func TestPop(t *testing.T) {
	size := 5
	q := New(size)

	for _, v := range testValues {
		q.Push(v)
	}

	for q.Len() > 0 {
		t.Log("current: ", q.Keys())

		v := q.Pop()

		// validate
		expect := testValues[len(testValues)-(q.Len()+1)]
		if v != expect {
			t.Error("expected %v but recevied %v", expect, v)
		}
	}

}
