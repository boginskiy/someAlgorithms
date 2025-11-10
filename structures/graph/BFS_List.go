package main

import "fmt"

type Graph struct {
	ajList map[int][]int //Список смежности вершин
}

func NewGraph(verticesCount int) *Graph {
	return &Graph{
		ajList: make(map[int][]int),
	}
}

// Метод добавления ребра
func (g *Graph) AddEdge(u, v int) {
	g.ajList[u] = append(g.ajList[u], v)
	g.ajList[v] = append(g.ajList[v], u)
}

// Алгоритм Breadth-First Search: начиная с начальной вершины startVertex,
// выводит порядок посещения вершин и возвращает массив расстояний
// от стартовой вершины до каждой другой вершины.
func bfs(graph *Graph, startVertex int) ([]bool, []int, []int) {
	visited := make([]bool, len(graph.ajList))    // Массив посещённых вершин
	distance := make([]int, len(graph.ajList))    // Массив расстояний от startVertex
	vertexes := make([]int, 0, len(graph.ajList)) // Узлы которые обошли

	// Изначально расстояние бесконечно (-1)
	for i := range distance {
		distance[i] = -1
	}

	// Очередь
	queue := LinkedList2{}
	queue.Put2(startVertex)
	visited[startVertex] = true
	distance[startVertex] = 0 // Расстояние от себя самого равно нулю

	//
	for queue.Len > 0 {
		vertexElement := queue.Take2()
		valueVertex := vertexElement.value

		//
		vertexes = append(vertexes, valueVertex)

		// Обход соседей
		neighbors := graph.ajList[valueVertex]
		for _, neighbor := range neighbors {
			if !visited[neighbor] {

				visited[neighbor] = true
				distance[neighbor] = distance[valueVertex] + 1
				queue.Put2(neighbor)
			}
		}
	}
	return visited, distance, vertexes
}

func main() {
	graph := NewGraph(8)
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(1, 4)
	graph.AddEdge(2, 5)
	graph.AddEdge(2, 6)
	graph.AddEdge(3, 7)
	graph.AddEdge(4, 7)
	visited, distance, vertexes := bfs(graph, 2)

	fmt.Println(visited, distance, vertexes)
}
