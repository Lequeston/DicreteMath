package main

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
			}

			visitVertex(name)
		}
	}
}

