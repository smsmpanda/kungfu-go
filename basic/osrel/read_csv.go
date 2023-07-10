package osrel

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func Gorun123f() {
	bks := make([]Book, 1)
	file, err := os.Open("resources/product.txt")
	if err != nil {
		log.Fatalf("Error %s opening file products.txt: ", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {

		// read one line from the file:
		line, err := reader.ReadString('\n')
		readErr := err

		line = string(line[:len(line)-2])

		strSl := strings.Split(line, ";")
		book := new(Book)
		book.title = strSl[0]
		book.price, err = strconv.ParseFloat(strSl[1], 32)

		if err != nil {
			fmt.Printf("Error in file: %v", err)
		}
		//fmt.Printf("The quan was:-%s-", strSl[2])
		book.quantity, err = strconv.Atoi(strSl[2])

		if err != nil {
			fmt.Printf("Error in file: %v", err)
		}
		if bks[0].title == "" {
			bks[0] = *book
		} else {
			bks = append(bks, *book)
		}
		if readErr == io.EOF {
			break
		}
	}

	fmt.Println("We have read the following books from the file: ")
	for _, bk := range bks {
		fmt.Println(bk)
	}
}
