package main

import (
	"github.com/skorobogatov/input"
	"strconv"
)

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
				h := true
				for _, k := range gr.graph[parent]{
					if k.name == nameStr{
						h  = false
						break
					}
				}
				if h {
					gr.graph[parent] = append(gr.graph[parent], Edge{name: nameStr})
				}
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