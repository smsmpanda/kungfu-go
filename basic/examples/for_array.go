package examples

import "fmt"

func Going111111() {
	var arr1 [5]int
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i * 2
	}

	for i := 0; i < len(arr1); i++ {
		fmt.Printf("Array at index %d is %d\n", i, arr1[i])
	}

	fmt.Println("----------------------------------------")
	for k, v := range arr1 {
		fmt.Printf("Array at index %d is %d\n", k, v)
	}
	fmt.Println("----------------------------------------")
	a := [...]string{"a", "b", "c", "d", "e"}
	for i, v := range a {
		fmt.Println("Array item", i, "is", a[i], v)
	}
	fmt.Println("----------------------------------------")
	var ar [3]int
	//f(ar)
	fp(&ar)
	f(ar)
	fmt.Println("----------------------------------------")
	var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}
	fmt.Println(arrKeyValue[0])
	fmt.Println(arrKeyValue[1])
	fmt.Println(arrKeyValue[2])
	fmt.Println(arrKeyValue[3])
	fmt.Println(arrKeyValue[4])
}

func f(a [3]int) {
	a[0] = 1314
	fmt.Println(a)
}

func fp(a *[3]int) {
	a[0] = 520
	fmt.Println(a)

}
