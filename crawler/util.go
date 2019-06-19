package crawler

import "fmt"
import "net/http"
import "github.com/PuerkitoBio/goquery"
import "strings"
import "net/http/cookiejar"
import "log"

func GetExecutionCode() string {
	resp, err := http.Get("https://cas.sustech.edu.cn/cas/login?service=http%3A%2F%2Fjwxt.sustech.edu.cn%2Fjsxsd%2F")

	if err != nil {
		fmt.Println("Error!")
		return ""
	}
	defer resp.Body.Close()
	doc, err2 := goquery.NewDocumentFromReader(resp.Body)

	if err2 != nil {
		fmt.Println("Error in parsing")
		return ""
	}
	executionCode := ""
	doc.Find("section>input").Each(func(i int, selection *goquery.Selection) {
		if i == 0 {
			executionCode, _ = selection.Attr("value")
		}
	})
	return executionCode
}

func Login(username string, password string) *http.Client {

	var r http.Request
	r.ParseForm()
	r.Form.Add("username", username)
	r.Form.Add("password", password)
	r.Form.Add("execution", GetExecutionCode())
	r.Form.Add("_eventId", "submit")
	r.Form.Add("geolocation", "")

	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{Jar: jar}
	req, err := http.NewRequest("POST", "https://cas.sustech.edu.cn/cas/login?service=http%3A%2F%2Fjwxt.sustech.edu.cn%2Fjsxsd%2F",
		strings.NewReader(r.Form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	
	resp, err := client.Do(req)

	if err != nil || resp.StatusCode != 200 {
		return nil
	}
	
	defer resp.Body.Close()
	return client
}
