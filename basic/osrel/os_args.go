package osrel

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var NewLine = flag.Bool("n1", false, "print newline") // echo -n flag, of type *bool

const (
	Space   = " "
	Newline = "\n"
)

func Goruff_os_args_n() {
	//声明一个数字数组，长度为5，并随机给5个不同的值
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	for i, v := range a {
		fmt.Println(i, v)
	}

	pt := &a
	_pt := *pt
	fmt.Print(pt)
	fmt.Print(_pt)
	fmt.Print(_pt == a)

	//声明一个map，将key和value分别设置为1和2
	m := make(map[string]int)
	m["key"] = 1
	m["value"] = 2
	fmt.Println(m)
	fmt.Println(m["key"])
	fmt.Println(m["value"])

	//声明一个切片
	who := "Alice "
	if len(os.Args) > 1 {
		who += strings.Join(os.Args[1:], " ")
	}
	fmt.Println("Good Morning", who)

	var NewLine = flag.Bool("n2", false, "print newline")
	println(NewLine)

	flag.PrintDefaults()
	flag.Parse() // Scans the arg list and sets up flags
	var s string = ""
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += " "
			if *NewLine { // -n is parsed, flag becomes true
				s += Newline
			}
		}
		s += flag.Arg(i)
	}
	os.Stdout.WriteString(s)

}
