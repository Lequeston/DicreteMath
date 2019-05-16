package main

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
