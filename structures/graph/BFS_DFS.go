package main

import (
	"fmt"
)

/*
	Breadth-First Search - алгорит поиска в ширину

	Поиск в ширину позволяет найти кратчайшее расстояние между двумя объ­ектами.

	Алгоритм поиска, который работает для графа
	Отвечает на вопросы двух типов.
		тип 1: существует ли путь от узла А к узлу В?
		тип 2: как выглядит кратчайший путь от узла А к узлу В?
*/

/*
	Depth First Search - поиск в глубину

*/

func breadthFirstSearch(queue []string, graph map[string][]string, target string) string {
	checkIn := make(map[string]struct{})

	for len(queue) > 0 { // Пока очередь не пуста
		man := queue[0] // Извлекаем человека
		queue = queue[1:]

		// Проверка человека, если он не проверялся ранее
		if _, ok := checkIn[man]; !ok {

			// Проверка, что человек является продавцом манго
			if man == target {
				return man
			} else {
				queue = append(queue, graph[man]...)
				checkIn[man] = struct{}{}
			}
		}
	}
	return ""
}

func main() {
	// Graph List
	graph := map[string][]string{}
	graph["you"] = []string{"alice", "bob", "claire"}
	graph["bob"] = []string{"anuj", "peggy"}
	graph["alice"] = []string{"peggy"}
	graph["claire"] = []string{"thom", "jonny"}
	graph["anuj"] = []string{}
	graph["peggy"] = []string{}
	graph["thom"] = []string{}
	graph["jonny"] = []string{}

	var queue []string
	queue = append(queue, graph["you"]...)

	man := breadthFirstSearch(queue, graph, "jonny")
	fmt.Printf("Mango Seller: %v\n", man)

}
