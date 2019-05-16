package main

import (
	"fmt"
	"github.com/skorobogatov/input"
)

//-------------------------------------------------------------

type Time struct { // вспомогательная структура для Digraph
	t1, t2 int //время захода и выхода
}

type Digraph struct { //структура орграфа
	vertex, edge    int      //кол-во вершин и ребер
	array           [][]int  //матрицы смежности
	timesVertex     []Time   //время захода и выхода для вершин
	numberComponent int      //кол-во компонент в графе
	vertexComponent []int    //номер компоненты для каждой вершины
	timesTarjan     []Time   //время захода и выхода для Тарьяна
	low             []int    //мин время захода для Тарьяна
	stack           []int    //стэк для Тарьяна
	matrix          [][]bool //матрица смежности для комонент
	res             []int
	result          []int
}

//ввод матрицы смежности
func (gr *Digraph) input() {
	input.Scanf("%d %d", &gr.vertex, &gr.edge)
	gr.array = make([][]int, gr.vertex)

	var a, b int
	for j := 0; j < gr.edge; j++ {
		input.Scanf("%d %d", &a, &b)
		gr.array[a] = append(gr.array[a], b)
	}
}

//вывод матицы смежности
func (gr *Digraph) output() {
	for p, a := range gr.array {
		fmt.Println(p, "=", a)
	}
}

//проход в глубину со временем
func (gr *Digraph) visitVertexTime(time, v int) int {
	gr.timesVertex[v].t1 = time
	time++
	for _, u := range gr.array[v] {
		if gr.timesVertex[u].t1 == 0 {
			time = gr.visitVertexTime(time, u)
		}
	}
	gr.timesVertex[v].t2 = time
	time++
	return time
}

//высчитывает время захода и выхода вершины
func (gr *Digraph) calculateTimes() {
	time := 1
	gr.timesVertex = make([]Time, gr.vertex)

	for i := 0; i < gr.vertex; i++ {
		if gr.timesVertex[i].t1 == 0 {
			time = gr.visitVertexTime(time, i)
		}
	}
}

//обход вглубину для Тарьяна
func (gr *Digraph) visitVertexTarjan(v, time int) int {
	gr.timesTarjan[v].t1, gr.low[v] = time, time
	time++
	gr.stack = append(gr.stack, v)
	for _, u := range gr.array[v] {
		if gr.timesTarjan[u].t1 == 0 {
			time = gr.visitVertexTarjan(u, time)
		}
		if gr.vertexComponent[u] == 0 && gr.low[v] > gr.low[u] {
			gr.low[v] = gr.low[u]
		}
	}
	len := len(gr.stack) - 1
	if gr.timesTarjan[v].t1 == gr.low[v] {
		u := gr.stack[len]
		//fmt.Println(gr.stack)
		for u != v {
			gr.vertexComponent[u] = gr.numberComponent
			len--
			u = gr.stack[len]
		}
		gr.vertexComponent[u] = gr.numberComponent
		gr.stack = gr.stack[:len]
		//fmt.Println(gr.stack)
		gr.numberComponent++
	}
	return time
}

//вызов алгоритма Тарьяна
func (gr *Digraph) tarjan() {
	gr.timesTarjan = make([]Time, gr.vertex)
	gr.vertexComponent = make([]int, gr.vertex)
	gr.low = make([]int, gr.vertex)
	gr.stack = make([]int, 0)

	time := 1
	gr.numberComponent = 1

	for i := 0; i < gr.vertex; i++ {
		if gr.timesTarjan[i].t1 == 0 {
			time = gr.visitVertexTarjan(i, time)
		}
	}
}

//-------------------------------------------------------------

func (gr *Digraph) initMatrix() {
	gr.res = make([]int, 0)
	gr.numberComponent--
	gr.matrix = make([][]bool, gr.numberComponent)
	for i := range gr.matrix {
		gr.matrix[i] = make([]bool, gr.numberComponent)
	}
	for i, num1 := range gr.vertexComponent {
		for _, ver := range gr.array[i] {
			num2 := gr.vertexComponent[ver]
			if num1 != num2 {
				gr.matrix[num1-1][num2-1] = true
			}
		}
	}
	for i := range gr.matrix {
		p := false
		for j := 0; j < gr.numberComponent; j++ {
			p = p || gr.matrix[j][i]
		}
		if !p {
			gr.res = append(gr.res, i+1)
		}
	}
	gr.result = make([]int, len(gr.result))
	for _, num := range gr.res {
		for i, comp := range gr.vertexComponent {
			if num == comp {
				gr.result = append(gr.result, i)
				break
			}
		}
	}
}

func main() {
	var gr Digraph
	gr.input()
	gr.tarjan()
	//fmt.Println(gr.vertexComponent)
	gr.initMatrix()
	//fmt.Println(gr.res)
	if gr.numberComponent == 0 {
		fmt.Println(0)
	} else {
		for _, p := range gr.result {
			fmt.Print(p, " ")
		}
	}
}
