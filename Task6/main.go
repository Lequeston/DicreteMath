package main

import (
	"fmt"
	"github.com/skorobogatov/input"
)

const INFINITY = 21000000

func compare(a *Vertex, b *Vertex) bool{
	return a.minSum < b.minSum
}

//структура самой очереди
type PriorityQueue struct {
	heap []*Vertex //массив для очереди
	compare func(*Vertex, *Vertex)bool //функция сравнения
}

//инициализация очереди
func initPriorityQueue(n int, compare func(*Vertex, *Vertex)bool) (res PriorityQueue){
	res.heap = make([]*Vertex, 0, n)
	res.compare = compare
	return
}

//проверка очереди на пустоту
func (q *PriorityQueue) queueEmpty()bool{
	return len(q.heap) == 0
}

//добавление элемента в очередь
func (q *PriorityQueue) insert(elem *Vertex){
	i := len(q.heap)
	//fmt.Println(i)
	q.heap = append(q.heap, elem)
	for i > 0 && q.compare(q.heap[i], q.heap[(i - 1) / 2]){
		q.heap[i], q.heap[(i - 1) / 2] = q.heap[(i - 1) / 2], q.heap[i]
		q.heap[i].index = i
		i = (i - 1) / 2
	}
	q.heap[i].index = i
	//fmt.Println(i)
}

//удаление пары с максимальным ключом
func (q *PriorityQueue) extractMax() (res *Vertex){
	if len(q.heap) == 0 {
		panic("очередь пуста")
	}
	res = q.heap[0]
	if len(q.heap) > 0{
		q.heap[0] = q.heap[len(q.heap) - 1]
		q.heap[0].index = 0
		q.heap = q.heap[:len(q.heap) - 1]
		//хипифай
		func(i int){
			for{
				l := 2 * i + 1
				r := l + 1
				j := i
				if l < len(q.heap) && q.compare(q.heap[l], q.heap[i]){
					i = l
				}
				if r < len(q.heap) && q.compare(q.heap[r], q.heap[i]){
					i = r
				}
				if i == j{
					break
				}
				q.heap[i], q.heap[j] = q.heap[j], q.heap[i]
				q.heap[i].index, q.heap[j].index = i, j
			}
		}(0)
	}
	return
}

func (q *PriorityQueue) IncreaseKey(index int, value int) {
	i := index
	q.heap[i].minSum = value
	for i > 0 && q.compare(q.heap[i], q.heap[(i - 1) / 2]){
		q.heap[i], q.heap[(i - 1) / 2] = q.heap[(i - 1) / 2], q.heap[i]
		q.heap[i].index = i
		i = (i - 1) / 2
	}
	q.heap[i].index = i
}

var k int

//структура для вершины
type Vertex struct {
	weight, minSum, index int //вес вершины и минимальная сумма
	x, y int
}

//структура для графа
type Graph struct {
	n int
	graph [][]Vertex //двухмерный граф
	queue PriorityQueue
}

//ввод графа
func (gr *Graph) input(){
	gr.queue = initPriorityQueue(gr.n, compare)
	input.Scanf(" %d\n", &gr.n)
	gr.graph = make([][]Vertex, gr.n)
	for i := 0; i < gr.n; i++{
		gr.graph[i] = make([]Vertex, gr.n)
		for j := 0; j < gr.n; j++{
			var a int
			v := &gr.graph[i][j]
			input.Scanf(" %d", &a)
			v.weight = a
			if i == 0 && j == 0{
				v.minSum = a
			} else {
				v.minSum = INFINITY
			}
			v.x = i
			v.y = j
			gr.queue.insert(v)
		}
	}
}

//вывод графа
func (gr *Graph) output(){
	for i := 0; i < gr.n; i++{
		for j := 0; j < gr.n; j++{
			fmt.Print(gr.graph[i][j].minSum, " ")
		}
		fmt.Println()
	}
}

func (gr *Graph) relax(x, y, x2, y2 int) (changed bool){
	v := &gr.graph[x][y]
	u := &gr.graph[x2][y2]
	w := u.weight
	changed = v.minSum + w <= u.minSum
	if changed{
		u.minSum = v.minSum + w
	}
	return
}

func (gr *Graph) dijkstra(){
	for !gr.queue.queueEmpty(){
		v := gr.queue.extractMax()
		v.index = -1

		var help func(int, int, int, int)

		help = func(x1, y1, x2, y2 int) {
			if x2 < gr.n && y2 < gr.n && x2 >= 0 && y2 >= 0 && gr.graph[x2][y2].index != -1 && gr.relax(x1, y1, x2, y2){
				gr.queue.IncreaseKey(gr.graph[x2][y2].index, gr.graph[x2][y2].minSum)
			}
		}

		if v.x == gr.n - 1 && v.y == gr.n - 1{
			break
		} else {
			help(v.x, v.y, v.x-1, v.y)
			help(v.x, v.y, v.x+1, v.y)
			help(v.x, v.y, v.x, v.y-1)
			help(v.x, v.y, v.x, v.y+1)
		}
	}
}

func main(){
	var graph Graph
	k = 0
	graph.input()
	graph.dijkstra()
	fmt.Println(graph.graph[graph.n - 1][graph.n - 1].minSum)
}
