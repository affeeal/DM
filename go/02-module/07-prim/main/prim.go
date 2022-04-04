package main

import "fmt"

type PriorityQueue struct {
	heap       []*Vertex
	cap, count int
}

type Vertex struct {
	index, key int
	value      *Vertex
	list       *List
}

type List struct {
	value *Vertex
	key   int
	next  *List
}

func InitPriorityQueue(N int) *PriorityQueue {

	q := new(PriorityQueue)
	q.heap = make([]*Vertex, N)
	for i := 0; i < N; i++ {
		q.heap[i] = nil
	}
	q.cap = N
	q.count = 0
	return q
}

func InitVertex() *Vertex {

	v := new(Vertex)
	v.index = -1
	v.key = 0
	v.value = nil
	v.list = InitList()
	return v
}

func InitList() *List {

	l := new(List)
	l.value = nil
	l.key = 0
	l.next = nil
	return l
}

func InsertList(u, v *Vertex, a int) {

	l := u.list
	for l.next != nil {

		if l.next.key > a {

			t := l.next
			l.next = InitList()
			l.next.value = v
			l.next.key = a
			l.next.next = t
			return
		}
		l = l.next
	}
	l.next = InitList()
	l.next.value = v
	l.next.key = a
}

func EmptyQueue(q *PriorityQueue) bool {

	return q.count == 0
}

func Heapify(i, n int, vs []*Vertex) {

	for {
		l := 2*i + 1
		r := l + 1
		j := i
		if l < n && vs[i].key > vs[l].key {

			i = l
		}
		if r < n && vs[i].key > vs[r].key {

			i = r
		}
		if i == j {
			break
		}
		vs[i], vs[j] = vs[j], vs[i]
		vs[i].index = i
		vs[j].index = j
	}
}

func ExtractMin(q *PriorityQueue) *Vertex {

	u := q.heap[0]
	q.count--
	if q.count > 0 {

		q.heap[0] = q.heap[q.count]
		q.heap[0].index = 0
		Heapify(0, q.count, q.heap)
	}
	return u
}

func InsertQueue(q *PriorityQueue, u *Vertex) {

	i := q.count
	q.count++
	q.heap[i] = u
	for i > 0 && q.heap[(i-1)/2].key > q.heap[i].key {

		q.heap[(i-1)/2], q.heap[i] = q.heap[i], q.heap[(i-1)/2]
		q.heap[i].index = i
		i = (i - 1) / 2
	}
	q.heap[i].index = i
}

func DecreaseKey(q *PriorityQueue, u *Vertex, key int) {

	i := u.index
	u.key = key
	for i > 0 && q.heap[(i-1)/2].key > key {

		q.heap[(i-1)/2], q.heap[i] = q.heap[i], q.heap[(i-1)/2]
		q.heap[i].index = i
		i = (i - 1) / 2
	}
	u.index = i
}

func MSTPrim(vs []*Vertex, N int) []*Vertex {

	mst := make([]*Vertex, 0)
	q := InitPriorityQueue(N)
	v := vs[0]
	for {
		v.index = -2
		for l := v.list.next; l != nil; l = l.next {

			u := l.value
			if u.index == -1 {

				u.key = l.key
				u.value = v
				InsertQueue(q, u)
			} else if u.index != -2 && l.key < u.key {

				u.value = v
				DecreaseKey(q, u, l.key)
			}
		}
		if EmptyQueue(q) {
			break
		}
		v = ExtractMin(q)
		mst = append(mst, v)
	}
	return mst
}

func main() {

	var N, M int
	fmt.Scanf("%d", &N)
	fmt.Scanf("%d", &M)

	vs := make([]*Vertex, N)
	for i := 0; i < N; i++ {

		vs[i] = InitVertex()
	}

	var u, v, a int
	for i := 0; i < M; i++ {

		fmt.Scanf("%d %d %d\n", &u, &v, &a)
		InsertList(vs[u], vs[v], a)
		InsertList(vs[v], vs[u], a)
	}
	mst, sum := MSTPrim(vs, N), 0
	for _, v := range mst {

		sum += v.key
	}
	fmt.Printf("%d\n", sum)
}
