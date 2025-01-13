package methods

import (
	"fmt"
	"net/http"
)

func Post(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Type of ResponseWriter: %T\n", w)
	fmt.Printf("%+v\n", r)

	fmt.Println(r.PostFormValue("ID"))
}