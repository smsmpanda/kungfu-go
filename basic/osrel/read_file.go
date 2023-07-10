package osrel

import (
	"fmt"
	"os"
	"path/filepath"
)

type Book struct {
	title    string
	price    float64
	quantity int
}

func Gorun1d() {
	file, err := os.Open("resources/test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var col1, col2, col3 []string
	for {
		var v1, v2, v3 string
		_, err := fmt.Fscanln(file, &v1, &v2, &v3)
		// scans until newline
		if err != nil {
			break
		}
		col1 = append(col1, v1)
		col2 = append(col2, v2)
		col3 = append(col3, v3)
	}

	fmt.Println(col1)
	fmt.Println(col2)
	fmt.Println(col3)

	filename := filepath.Base("C:\\Users\\19705\\Downloads\\favicon_io")
	fmt.Println(filename)
}
