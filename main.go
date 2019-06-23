package main

import "net/http"
import router "SUSTechHelperGoBackend/router"
import "log"

func main() {

	http.HandleFunc("/gpa", router.GPAQuery)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
