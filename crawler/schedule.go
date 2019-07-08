package crawler
import "github.com/Jeffail/gabs"
import "github.com/PuerkitoBio/goquery"
import "net/http"

func GetSchedule(username string, password string ) string {

	client := Login(username, password)

	jsonObj := gabs.New()
	if client == nil {
		jsonObj.Set(false, "state")
		return jsonObj.String()
	}

	req, _ := http.NewRequest("GET","http://jwxt.sustech.edu.cn/jsxsd/xskb/xskb_list.do",nil)

	
	resp, err := client.Do(req)
	
	if err != nil {
		jsonObj.Set(false, "state")
		return jsonObj.String()
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	
	if err != nil {
		jsonObj.Set(false, "state")
		return jsonObj.String()
	}

	

	table := doc.Find("#kbtable")
	var name string
	table.Find("tbody>tr").Each(func(i int,selection *goquery.Selection){
		 if i != 0{
			

			selection.Find("th").Each(func(i int,selection *goquery.Selection){
				if i == 0{
				  name = selection.Text()
				  jsonObj.Array(name,"class")

				} else{
					jsonObj.ArrayAppend(selection.Text(),name,"class")
				}
				
				
			})
		 }
		 
		 
	})

	return jsonObj.String()
}