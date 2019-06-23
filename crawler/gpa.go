package crawler

import "sync"
import "net/http"
import "github.com/PuerkitoBio/goquery"
import "strings"
import "github.com/Jeffail/gabs"

var semester = []string{"2010-2011-1", "2010-2011-2", "2011-2012-1", "2011-2012-2", "2012-2013-1", "2012-2013-2", "2013-2014-1", "2013-2014-2", "2014-2015-1", "2014-2015-2", "2014-2015-3", "2015-2016-1", "2015-2016-2", "2015-2016-3", "2016-2017-1", "2016-2017-2", "2016-2017-3", "2017-2018-1", "2017-2018-2", "2017-2018-3", "2018-2019-1", "2018-2019-2", "2018-2019-3"}

func GetAllGrade(username string, password string) string {

	client := Login(username, password)

	jsonObj := gabs.New()
	if client == nil {
		jsonObj.Set(false, "state")
		return jsonObj.String()
	}

	var wg sync.WaitGroup
	var lock sync.Mutex
	for _, period := range semester {
		go Currency(period, jsonObj, client, &wg, &lock)
		wg.Add(1)
	}
	wg.Wait()
	jsonObj.Set(true, "state")
	return jsonObj.String()
}

func Currency(period string, jsonObj *gabs.Container, client *http.Client, wg *sync.WaitGroup, lock *sync.Mutex) {

	var r http.Request
	r.ParseForm()
	r.Form.Add("kksj", period)
	r.Form.Add("kcxz", "")
	r.Form.Add("kcmc", "")
	r.Form.Add("xsfs", "all")

	req, _ := http.NewRequest("POST", "http://jwxt.sustech.edu.cn/jsxsd/kscj/cjcx_list", strings.NewReader(r.Form.Encode()))

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {

	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	table := doc.Find("#dataList")
	courseName := ""
	lock.Lock()
	table.Find("tbody>tr").Each(func(i int, selection *goquery.Selection) {
		if i != 0 {
			selection.Find("td").Each(func(i int, selection *goquery.Selection) {
				if i == 3 {
					courseName = selection.Text()
					jsonObj.Array(period, courseName) //course name
				}
				if 4 <= i && i <= 6 {
					jsonObj.ArrayAppend(selection.Text(), period, courseName)
				}
			})
		}

	})
	lock.Unlock()
	wg.Done()
}
