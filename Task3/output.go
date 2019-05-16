package main

import "fmt"

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
