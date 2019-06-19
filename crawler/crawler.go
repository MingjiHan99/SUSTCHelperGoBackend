package crawler

import "fmt"
import "sync"
import "net/http"
import "github.com/PuerkitoBio/goquery"
import "strings"
import "net/http/cookiejar"
import "log"
//import "io/ioutil"
import "github.com/Jeffail/gabs"


var semester = []string{"2018-2019-1","2010-2011-1", "2010-2011-2", "2011-2012-1", "2011-2012-2", "2012-2013-1", "2012-2013-2", "2013-2014-1", "2013-2014-2", "2014-2015-1", "2014-2015-2", "2014-2015-3", "2015-2016-1", "2015-2016-2", "2015-2016-3", "2016-2017-1", "2016-2017-2", "2016-2017-3", "2017-2018-1", "2017-2018-2", "2017-2018-3", "2018-2019-1","2018-2019-2","2018-2019-3"}

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

func GetAllGrade(username string, password string) string{
	
	client := Login(username, password)

	jsonObj := gabs.New()
	if client == nil{
		jsonObj.Set(false,"state")
		return jsonObj.String()
	}

	
	var wg sync.WaitGroup
	var lock sync.Mutex
	for _, period := range semester {
		go Currency(period,jsonObj,client,&wg,&lock)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println("Final result")
	jsonObj.Set(true,"state")
	return jsonObj.String()
	
	
	
}


func Currency(period string, jsonObj *gabs.Container,client *http.Client,wg *sync.WaitGroup,lock *sync.Mutex){
		
		var r http.Request
		//fmt.Println("Gooooooooooooo")
		r.ParseForm()
		r.Form.Add("kksj", period)
		r.Form.Add("kcxz", "")
		r.Form.Add("kcmc", "")
		r.Form.Add("xsfs", "all")
		
		req, _ := http.NewRequest("POST", "http://jwxt.sustech.edu.cn/jsxsd/kscj/cjcx_list", strings.NewReader(r.Form.Encode()))

		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		resp, err := client.Do(req)
		if err != nil{

		}
		defer resp.Body.Close()

		doc, err := goquery.NewDocumentFromReader(resp.Body)

		if err != nil{

		}

		table := doc.Find("#dataList")
		courseName := ""
		lock.Lock()
		table.Find("tbody>tr").Each(func(i int, selection *goquery.Selection){
			if i != 0{
				selection.Find("td").Each(func(i int, selection *goquery.Selection){
					if i == 3 {
						courseName = selection.Text()
						
						jsonObj.Array(period,courseName) //course name
					}
					if 4 <= i && i <= 6{
						jsonObj.ArrayAppend(selection.Text(),period,courseName)
						
					}
				})
				//fmt.Println(jsonObj.String())	
			}
			
		})
		lock.Unlock()
		wg.Done()
}

