package main

//DFS-ы закрашивающие граф в красный
func (gr *Graph) dfsFillTheRed() {
	redVertex := make(map[string]vertexRed, len(gr.vertex))
	colorVertex := make(map[string]int, len(gr.vertex))
	res := make([]string, 0, 2)
	max := 0
	func(){

		for _, name := range gr.global{
			if colorVertex[name] == WHITE && !gr.vertex[name].blue{
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
					redVertex[v] = vertexRed{max: max + gr.vertex[v].weight, children: arr}
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

	//fmt.Println(redVertex)
	for _, n := range res {

		var colorRed func(string, string)

		colorRed = func(v, parent string) {
			gr.vertex[v] = Vertex{weight: gr.vertex[v].weight, red: true}
			if parent != "" {
				for i, k := range gr.graph[parent] {
					if k.name == v {
						gr.graph[parent][i].red = true
						break
					}
				}
			}

			for _, name := range redVertex[v].children {
				colorRed(name, v)
			}
		}

		colorRed(n, "")
	}
}
