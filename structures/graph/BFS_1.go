package main

import (
	"fmt"
)

/*
	Breadth-First Search - алгоритм поиска в ширину

	Поиск в ширину позволяет найти кратчайшее расстояние между двумя объ­ектами.

	Алгоритм поиска, который работает для графа
	Отвечает на вопросы двух типов.
		тип 1: существует ли путь от узла А к узлу В?
		тип 2: как выглядит кратчайший путь от узла А к узлу В?
*/

// // Очередь на двунаправленном LinkedList. Интерфейс: put, take
// type Node struct {
// 	value string
// 	next  *Node
// 	prev  *Node
// }
// type LinkedList struct {
// 	Head *Node
// 	Tail *Node
// 	cnt  int
// }

// func (l *LinkedList) Print() {
// 	node := l.Head
// 	fmt.Print("HEAD: ")
// 	for node != nil {
// 		fmt.Printf("%v ", node.value)
// 		node = node.next
// 	}
// 	fmt.Println()
// }
// func (l *LinkedList) Put(elems ...string) {
// 	for _, elem := range elems {
// 		newNode := &Node{value: elem}
// 		// Вставка первого элемента
// 		if l.cnt == 0 {
// 			l.Head, l.Tail = newNode, newNode
// 		} else {
// 			l.Tail.next = newNode
// 			newNode.prev = l.Tail
// 			l.Tail = newNode
// 		}
// 		l.cnt++
// 	}
// }
// func (l *LinkedList) Take() *Node {
// 	node := l.Head
// 	if l.cnt > 1 {
// 		l.Head = l.Head.next
// 		l.Head.prev = nil
// 		node.next = nil
// 		l.cnt--
// 	} else if l.cnt == 1 {
// 		l.Head, l.Tail = nil, nil
// 		l.cnt--
// 	}
// 	return node
// }

func breadthFirstSearch(queue *LinkedList, graph map[string][]string, target string) *Node {
	checkIn := make(map[string]struct{})

	for queue.cnt > 0 { // Пока очередь не пуста
		man := queue.Take() // Извлекаем человека

		// Проверка человека, если он не проверялся ранее
		if _, ok := checkIn[man.value]; !ok {

			// Проверка, что человек является продавцом манго
			if man.value == target {
				return man
			} else {
				queue.Put(graph[man.value]...)
				checkIn[man.value] = struct{}{}
			}
		}
	}
	return nil
}

func main() {
	// LinkedList
	queue := &LinkedList{}

	// Graph
	graph := map[string][]string{}
	graph["you"] = []string{"alice", "bob", "claire"}
	graph["bob"] = []string{"anuj", "peggy"}
	graph["alice"] = []string{"peggy"}
	graph["claire"] = []string{"thom", "jonny"}
	graph["anuj"] = []string{}
	graph["peggy"] = []string{}
	graph["thom"] = []string{}
	graph["jonny"] = []string{}

	// Добавляем всех соседей в очередь поиска
	queue.Put(graph["you"]...)
	// Поиск
	man := breadthFirstSearch(queue, graph, "thom")
	// Вывод
	fmt.Printf("Mango Seller: %v\n", man.value)

}
