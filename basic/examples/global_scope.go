package main

var a = "G"

func main() {
	n()
	m()
	n()

	//结果：G O O
}

func n() {
	print(a)
}

func m() {
	a = "O"
	print(a)
}
