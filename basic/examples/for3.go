package examples

import "fmt"

func main6() {
	var i int = 4
	for {
		i = i - 1
		fmt.Printf("The variable i is now: %d\n", i)
		if i < 0 {
			break
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 10; j++ {
			if j > 5 {
				break
			}
			print(j)
		}
		print("  ")
	}
}
