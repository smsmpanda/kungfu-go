package network

import (
	"fmt"
	"log"
	"net/http"
)

type Had struct{}

func (had *Had) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<title>👩‍👩‍👧Go lang</title><body>%s %s</body>", "Web handle", "go")
}

func Gorunrouter() {

	had := &Had{}

	http.HandleFunc("/", hellowServer)
	http.Handle("/had", had)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

func hellowServer(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<title>👩‍👩‍👧Go lang</title><body>%s %s</body>", "Web", "go")
}
