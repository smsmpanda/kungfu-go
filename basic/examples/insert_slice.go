package examples

import "fmt"

func Going() {

	s1 := []string{"L", "o", "v", "e"}
	s2 := []string{"I", "Y", "o", "u"}
	res := InsertStringSlice(s1, s2, 2)
	fmt.Println(res)
}

func InsertStringSlice(slice, insertion []string, index int) []string {
	result := make([]string, len(slice)+len(insertion))
	at := copy(result, slice[:index])
	at += copy(result[at:], insertion)
	copy(result[at:], slice[index:])
	return result
}
