package examples

import (
	"fmt"
	"math"
	"strings"
)

//1. 在循环中使用+拼接字符串并不是最高效的方法，更好的办法应该是使用函数string.join
//2. 使用字节缓冲bytes.buffer

func main14() {
	var n int16 = 34
	var m int32
	// compiler error: cannot use n (type int16) as type int32 in assignment
	//m = n
	m = int32(n)

	fmt.Printf("32 bit int is: %d\n", m)
	fmt.Printf("16 bit int is: %d\n", n)

	r, e := Uint8FromInt(256)
	fmt.Println(r)
	fmt.Println(e)

	var isTrue bool = strings.HasPrefix("jiangrui", "j")
	fmt.Printf("name begine of char 'j' start == %t \n", isTrue)

	var isFalse bool = strings.HasSuffix("jiangrui", "i")
	fmt.Printf("name begine of char 'j' end == %t \n", isFalse)

	var str1 string = "This is an example of a string"
	fmt.Printf("T/F? Does the string \"%s\" have prefix %s?\n", str1, "Th")
	fmt.Printf("%t\n", strings.HasPrefix(str1, "Th"))

	fmt.Println(strings.Contains("jiangrui", "drui"))

	fmt.Println(strings.Index(str1, "a"))
	fmt.Println(strings.LastIndex(str1, "t"))

	var str string = "Hi, I'm Marc, Hi."

	fmt.Printf("The position of \"Marc\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Marc"))

	fmt.Printf("The position of the first instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Hi"))
	fmt.Printf("The position of the last instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.LastIndex(str, "Hi"))

	fmt.Printf("The position of \"Burger\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Burger"))

	oldName := "JiangJunPing"
	newName := "JiangRui"
	resultName := strings.Replace("My old name is JiangJunPing JiangJunPing JiangJunPing JiangJunPing ,JiangJunPing,My now name is JiangRui.", oldName, newName, -1)
	resultName1 := strings.Replace("My old name is JiangJunPing JiangJunPing JiangJunPing JiangJunPing ,JiangJunPing,My now name is JiangRui.", oldName, newName, 2)
	fmt.Println(resultName)
	fmt.Println(resultName1)

	fmt.Println(strings.Count(oldName, "J"))

	var str2 string = "Hello, how is it going, Hugo?"
	var manyG = "gggggggggg"

	fmt.Printf("Number of H's in %s is: ", str2)
	fmt.Printf("%d\n", strings.Count(str, "H"))

	fmt.Printf("Number of double g's in %s is: ", manyG)
	fmt.Printf("%d\n", strings.Count(manyG, "gg"))

	fmt.Println(strings.ToLower("Jiangrui"))
	fmt.Println(strings.ToUpper("Jiangrui"))

	var example_str string = "Sasssssss "
	fmt.Println(example_str)
	fmt.Println(strings.TrimSpace(example_str))
	fmt.Println(strings.Trim(example_str, "S"))

}

func Uint8FromInt(n int) (uint8, error) {
	if 0 <= n && n <= math.MaxUint8 {
		return uint8(n), nil
	}
	return 0, fmt.Errorf("%d is out of uint8 range", n)
}
