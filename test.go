package main

type Deque struct {
	max   int
	size  int
	queue []int
	Head  int
	Tail  int
}

func NewDeque(n int) *Deque {
	return &Deque{max: n, queue: make([]int, n)}
}

//    def push_front(self, element: str):
//        if self.size == self.m_max:
//            raise Exception(self.massage_push)

//        self.head = (self.head - 1) % self.m_max
//        self.dek_arr[self.head] = element
//        self.size += 1
//        return element

//    def pop_front(self):
//        if self.is_empty():
//            raise Exception(self.massage_pop)

//        element = self.dek_arr[self.head]
//        self.dek_arr[self.head] = None
//        self.head = (self.head + 1) % self.m_max
//        self.size -= 1
//        return element

func (d *Deque) PushFront(elem int) {
	if d.size != d.max {
		d.queue[d.Head] = elem
		d.Head = (d.Head + 1) % d.max
		d.size++
	}
}

func (d *Deque) PopFront() (elem int) {
	if d.size != 0 {
		d.Head = (d.Head - 1) % d.max

		if d.Head < 0 {
			d.Head = (d.max + d.Head)
		}

		elem = d.queue[d.Head]
		d.queue[d.Head] = 0
		d.size--
	}
	return elem
}

func main() {
	queue := NewDeque(4)
}
