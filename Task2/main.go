package main

import (
	"fmt"
	"github.com/skorobogatov/input"
	"strconv"
)

const (
	WHITE = 0
	GRAY = 1
	BLACK = 2
)

//структура для вершин
type Vertex struct {
	weight    int  //вес, цвет вершины
	blue, red bool //какого цвета вершина синего или красного
}

//структура для ребер
type Edge struct {
	name      string //имя вершины
	blue, red bool   //какого цвета вершина синего или красного
}

//структура для графа
type Graph struct {
	global []string
	vertex    map[string]Vertex //ассоциативный массив для вершин
	graph     map[string][]Edge //ассоциативный  массив инцидентных вершин
}

type vertexRed struct {
	children []string
	max int
}


//ввод графа НЕ ТРОГАТЬ ДАЖЕ ПОД ПРИДЛОГОМ СМЕРТИ ОНО РАБОТАЕТ И ВСЕ ДАЖЕ НЕ ОТКРЫВАЙ ЭТОТ КОД!!!!!!!!!!
func (gr *Graph) input() {

	//инициализация самой вершины
	gr.vertex = make(map[string]Vertex)
	gr.graph = make(map[string][]Edge)
	iter := 0
	name := make([]byte, 0, 32)
	weightStr := make([]byte, 0, 4)
	parent := ""
	gr.global = make([]string, 0, 2)

Loop:
	for {
		input.Scanf("\n")
		str := input.Gets()
		iter = 0

		for iter < len(str) {

			for str[iter] == ' ' {
				iter++
			}

			for iter < len(str) && str[iter] != ' ' && str[iter] != '(' && str[iter] != ';' && str[iter] != '\n' {
				name = append(name, str[iter])
				iter++
			}

			nameStr := string(name)
			_, f := gr.vertex[nameStr]
			if parent == "" && !f{
				gr.global = append(gr.global, nameStr)
			}

			if iter < len(str) && str[iter] == '(' {
				//1)считываем вес
				iter++
				for str[iter] != ')' {
					weightStr = append(weightStr, str[iter])
					iter++
				}
				iter++

				//2)переводим в integer и добавляем вершину в массив вершин
				w, _ := strconv.Atoi(string(weightStr))
				//fmt.Println(w)
				gr.vertex[nameStr] = Vertex{weight: w}
			}

			if parent != "" {
				gr.graph[parent] = append(gr.graph[parent], Edge{name: nameStr})
			}

			if iter < len(str) && str[iter] == ';' {
				parent = ""
				iter++
			} else {
				for iter < len(str) && str[iter] == ' ' {
					iter++
				}

				if iter < len(str) && str[iter] == '<' {
					iter++
					parent = nameStr
				} else {
					break Loop
				}
			}
			name = name[:0]
			weightStr = weightStr[:0]
		}
	}
}

//DFS-ы закрашивающие граф в синий
func (gr *Graph) dfsFillTheBlue() {

	var stackBlue StackString //стэк главных синих вершин

	//DFS на определение циклов
	func() {
		colorVertex := make(map[string]int, len(gr.vertex)) //мапа для цветов
		//for name := range gr.vertex{
		//	colorVertex[name] = WHITE
		//}

		stackBlue = initStack(10)        //инициализируем стэк главных синих вершин

		for _, name := range gr.global { //бежим по всем вершинам графа
			//fmt.Println(v, color)
			if colorVertex[name] == WHITE { //если белая начинаем обход в глубину
				var visitVertex func(string)

				//посещение вершин
				visitVertex = func(v string){
					colorVertex[v] = GRAY
					for _, u := range gr.graph[v] {
						//fmt.Println(u.name, colorVertex[u.name])
						switch colorVertex[u.name] {
						case WHITE:
							visitVertex(u.name)
						case GRAY:
							stackBlue.push(u.name)
						}
					}
					colorVertex[v] = BLACK
				}

				visitVertex(name)
			}
		}
	}()

	//fmt.Println(stackBlue)
	//DFS на само закрашивание вершин
	for !stackBlue.isEmpty(){
		name := stackBlue.pop()
		v := gr.vertex[name]
		if !v.blue{
			colorVertex := make(map[string]int) //мапа для цветов
			var visitVertex func(string)

			//посещение вершин
			visitVertex = func(v string) {
				colorVertex[v] = GRAY
				gr.vertex[v] = Vertex{weight: gr.vertex[v].weight,
					blue:true}
				for i, u := range gr.graph[v] {
					if WHITE == colorVertex[u.name] {
						gr.graph[v][i].blue = true
						visitVertex(u.name)
					}

					if GRAY == colorVertex[u.name]{
						gr.graph[v][i].blue = true
					}
				}
				colorVertex[v] = BLACK
			}

			visitVertex(name)
		}
	}
}

//DFS-ы закрашивающие граф в красный
func (gr *Graph) dfsFillTheRed() {
	redVertex := make(map[string]vertexRed, len(gr.vertex))
	colorVertex := make(map[string]int, len(gr.vertex))
	res := make([]string, 0, 2)
	max := 0
	func(){

		for _, name := range gr.global{
			if colorVertex[name] == WHITE {
				var visitVertex func(string) (int, string)

				visitVertex = func(v string) (int, string) {
					colorVertex[v] = GRAY
					arr := make([]string, 0, 8)
					max := 0

					for _, u := range gr.graph[v] {
						if !gr.vertex[u.name].blue {
							if colorVertex[u.name] == WHITE {
								val, n := visitVertex(u.name)
								if val > max {
									max = val
									arr = arr[:0]
									arr = append(arr, n)
								} else if val == max {
									arr = append(arr, n)
								}
							} else if colorVertex[u.name] == GRAY {
								if redVertex[u.name].max > max {
									max = redVertex[u.name].max
									arr = arr[:0]
									arr = append(arr, u.name)
								} else if redVertex[u.name].max == max {
									arr = append(arr, u.name)
								}
							}
						}
					}

					//fmt.Println(v, arr, max)
					redVertex[v] = vertexRed{max: max, children: arr}
					return max + gr.vertex[v].weight, v
				}

				val, _ := visitVertex(name)

				if val > max {
					max = val
					res = res[:0]
					res = append(res, name)
				} else if val == max {
					res = append(res, name)
				}

				//fmt.Println(redVertex)
				//redVertex[name] = vertexRed{max: max, children: arr}
			}
		}
	}()

	for _, n := range res{
		if colorVertex[n] == GRAY {
			parent := ""

			var colorRed func(string)

			colorRed = func(v string) {
				gr.vertex[v] = Vertex{weight: gr.vertex[v].weight, red: true}
				if parent != "" {
					for i, k := range gr.graph[parent] {
						if k.name == v {
							gr.graph[parent][i].red = true
							break
						}
					}
				}
				parent = v
				for _, name := range redVertex[v].children {
					colorRed(name)
				}
			}

			colorRed(n)
			colorVertex[n] = BLACK
		}
	}
}

//стэк
type StackString []string

//инициализация стэка
func initStack(size int) StackString {
	return make(StackString, 0, size)
}

//добавление элемента в стэк
func (st *StackString) push(value string) {
	*st = append(*st, value)
}

//удаление элемента из стэка
func (st *StackString) pop() (value string) {
	value = (*st)[len(*st)-1]
	*st = (*st)[:len(*st)-1]
	return
}

//пуст ли стэк?
func (st *StackString) isEmpty() bool {
	return len(*st) == 0
}

func (gr *Graph) output() {
	fmt.Println("digraph {")
	for name, value := range gr.vertex{
		fmt.Print("  " + name + " [label = \"" + name + "(")
		fmt.Print(value.weight)
		fmt.Print(")\"")
		switch {
		case value.red:
			fmt.Print(", color = red")
		case value.blue:
			fmt.Print(", color = blue")
		}
		fmt.Println("]")
	}

	for vertex, arr := range gr.graph{
		for _, str := range arr{
			fmt.Print("  " + vertex + " -> " + str.name)
			switch {
			case str.red:
				fmt.Print(" [color = red]")
			case str.blue:
				fmt.Print(" [color = blue]")
			}
			fmt.Println()
		}

	}
	fmt.Println("}")
}

func main() {
	var graph Graph
	graph.input()
	graph.dfsFillTheBlue()
	//fmt.Println(graph.global)
	graph.dfsFillTheRed()
	graph.output()
}
