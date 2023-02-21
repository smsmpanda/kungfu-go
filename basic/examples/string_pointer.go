package examples

import "fmt"

func main11() {
	s := "good bye"
	var sptr *string = &s
	fmt.Printf("%p address store content is %s", sptr, *sptr)
}
