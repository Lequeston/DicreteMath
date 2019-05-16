package  main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const InputSize = 1024

func input() {
	read := bufio.NewReader(os.Stdin)
	res := make([]byte, 0, InputSize)
	in := make([]byte, InputSize)
	for {
		num, err := read.Read(in)
		if err == io.EOF {
			break
		} else if num < InputSize {
			res = append(res, in[:num]...)
			break
		} else {
			res = append(res, in...)
		}
	}
	fmt.Println(string(res))
}

func main(){

}