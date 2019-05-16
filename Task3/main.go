package main

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

func main() {
	var graph Graph
	graph.input()
	graph.dfsFillTheBlue()
	//fmt.Println(graph.global)
	graph.dfsFillTheRed()
	graph.output()
}
