package examples

var a = "G"

func main9() {
	n()
	m()
	n()

	//结果：G O G
}

func n() { print(a) }

func m() {
	a := "O"
	print(a)
}
