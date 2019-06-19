package router
import "../crawler"
import "fmt"
import "net/http"

func GPAQuery(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form["username"][0]
	password := r.Form["password"][0]

	fmt.Fprintln(w, "", crawler.GetAllGrade(username,password))
}