package examples

var a1 string

func main7() {
	a = "G"
	print(a)
	f1()
	//结果：G O G
}

func f1() {
	a := "O"
	print(a)
	f2()
}

func f2() {
	print(a)
}
