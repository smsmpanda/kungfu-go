package network

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var urls = []string{
	"https://www.baidu.com",
	"https://www.sina.com.cn",
	"https://www.qq.com",
}

func Gorunsq() {
	// for _, url := range urls {
	// 	resp, err := http.Head(url)
	// 	if err != nil {
	// 		fmt.Println("Error:", url, err)
	// 	}
	// 	fmt.Println(url, ": ", resp.Status)
	// }
	res, err := http.Get("http://www.baidu.com")
	checkError1(err)
	data, err := ioutil.ReadAll(res.Body)
	checkError1(err)
	fmt.Printf("Got: %q", string(data))
}

func checkError1(err error) {
	if err != nil {
		log.Fatalf("Get : %v", err)
	}
}
