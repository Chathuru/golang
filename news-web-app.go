package main

import(
	"net/http"
	"io/ioutil"
	"encoding/xml"
	"html/template"
)

type Sitemapindex struct {
	Locatons []string `xml:"url>loc"`
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
	Dates []string `xml:"url>news>publication_date"`
}

type NewsMap struct {
	Locatons string
	Keywords string
	Dates string
}

type htmlPage struct {
	Title string
	Body map[string]NewsMap
}

func index_hander(w http.ResponseWriter, r *http.Request) {
	resp, _ := http.Get("https://www.washingtonpost.com/news-business-sitemap.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	var s Sitemapindex
	xml.Unmarshal(bytes, &s)

	news := make(map[string]NewsMap)
	for idx, _ := range s.Keywords {
		news[s.Titles[idx]] = NewsMap{s.Locatons[idx], s.Keywords[idx], s.Dates[idx]}
	}

	p := htmlPage{Title: "Washingtonpost Business News", Body: news}
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, p)
}

func main(){
	http.HandleFunc("/", index_hander)
	http.ListenAndServe(":8000", nil)
}