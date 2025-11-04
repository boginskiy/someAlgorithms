package main

import "fmt"

// /////////// LinkedList2 //////////////////
type Node struct {
	value int
	next  *Node
	prev  *Node
}
type LinkedList struct {
	Head *Node
	Tail *Node
	Len  int
}

func (l *LinkedList) Print() {
	node := l.Head
	fmt.Print("HEAD: ")
	for node != nil {
		fmt.Printf("%v ", node.value)
		node = node.next
	}
	fmt.Println()
}
func (l *LinkedList) Put(elems ...int) {
	for _, elem := range elems {
		newNode := &Node{value: elem}
		// Вставка первого элемента
		if l.Len == 0 {
			l.Head, l.Tail = newNode, newNode
		} else {
			l.Tail.next = newNode
			newNode.prev = l.Tail
			l.Tail = newNode
		}
		l.Len++
	}
}
func (l *LinkedList) Take() *Node {
	node := l.Head
	if l.Len > 1 {
		l.Head = l.Head.next
		l.Head.prev = nil
		node.next = nil
		l.Len--
	} else if l.Len == 1 {
		l.Head, l.Tail = nil, nil
		l.Len--
	}
	return node
}

// ///////////// Graph //////////////////

// Матрица смежности представляет собой двунаправленную структуру связей между вершинами
var adjacencyMatrix = [][]int{
	{0, 1, 0, 0}, // A
	{0, 0, 1, 0}, // B
	{0, 0, 0, 1}, // C
	{0, 0, 0, 0}, // D
}

// Число вершин в нашем примере
const verticesCount = 4

func bfs2(startVertex int) ([]bool, []int) {
	visited := make([]bool, verticesCount) // Массив посещённых вершин
	distance := make([]int, verticesCount) // Массив расстояний от startVertex

	queue := LinkedList{}       // Очередь для вершин
	queue.Put(startVertex)      // Начало с заданной вершины
	visited[startVertex] = true // Отмечаем, что начали с этой вершины
	distance[startVertex] = 0   // Расстояние от себя самого равно нулю

	for queue.Len > 0 {
		valueNode := queue.Take().value // Получаем индекс текущего узла

		// Просматриваем строку матрицы смежности, соответствующую вершине frontNode
		for nextVertex := 0; nextVertex < verticesCount; nextVertex++ {
			if adjacencyMatrix[valueNode][nextVertex] == 1 && !visited[nextVertex] {

				visited[nextVertex] = true                       // Метим новую вершину как посещённую
				distance[nextVertex] = distance[startVertex] + 1 // Расстояние
				queue.Put(nextVertex)                            // Добавляем её в очередь

			}
		}
	}
	return visited, distance
}

func main() {
	visited, distance := bfs2(0) // Начальная вершина (A)
	fmt.Println(visited, distance)
}
