// sorting4
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"container/heap"
	"slices"
	"testing"
)

type task struct {
	name     string
	priority int // 1 is the most urgent
}

// queue is a priority queue: heap.Pop returns the most urgent task first.
type queue []task

func (q queue) Len() int      { return len(q) }
func (q queue) Swap(i, j int) { q[i], q[j] = q[j], q[i] }
func (q *queue) Push(x any)   { *q = append(*q, x.(task)) }
func (q *queue) Pop() any {
	old := *q
	last := old[len(old)-1]
	*q = old[:len(old)-1]
	return last
}

// TODO: heap.Pop returns the element that Less puts first, and this Less puts the biggest priority number first. The most urgent task, priority 1, has to come out first: change the comparison.
func (q queue) Less(i, j int) bool { return q[i].priority > q[j].priority }

func TestTheMostUrgentTaskComesFirst(t *testing.T) {
	q := &queue{}
	heap.Push(q, task{"write the report", 3})
	heap.Push(q, task{"fix the build", 1})
	heap.Push(q, task{"answer the mail", 2})

	var got []string
	for q.Len() > 0 {
		got = append(got, heap.Pop(q).(task).name)
	}

	want := []string{"fix the build", "answer the mail", "write the report"}
	if !slices.Equal(got, want) {
		t.Errorf("tasks came out as %q, want %q", got, want)
	}
}
