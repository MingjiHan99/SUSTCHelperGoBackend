package router

import crawler "SUSTechHelperGoBackend/crawler"
import user "SUSTechHelperGoBackend/user"
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

func LoginQuery(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	code := r.Form["code"][0]

	fmt.Fprintln(w, "", user.Login(code))
}

func SubscribeQuery(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	formID := r.Form["formId"][0]
	username := r.Form["username"][0]
	password := r.Form["password"][0]
	session := r.Form["session"][0]
	course := r.Form["course"][0]

	fmt.Fprintln(w, "", user.Subsribe(username, password, formID, course, session))
}
