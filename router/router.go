package router

import crawler "SUSTechHelperGoBackend/crawler"
import "fmt"
import "net/http"

func GPAQuery(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form["username"][0]
	password := r.Form["password"][0]

	fmt.Fprintln(w, "", crawler.GetAllGrade(username, password))
}

func CourseQuery(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form["username"][0]
	password := r.Form["password"][0]

	fmt.Fprintln(w, "", crawler.GetAllCourse(username, password))
}
