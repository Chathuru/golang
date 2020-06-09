package main

import(
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/xml"
)

type Moveindexs struct {
	Movetitles []string `xml:"channel>item>title"`
	Movedescriptions []string `xml:"channel>item>description"`
	MovepubDates []string `xml:"channel>item>pubDate"`
}

func index_hander(w http.ResponseWriter, r *http.Request) {
	var s Moveindexs
	out := "<html><head><style>div.gallery {border: 1px solid #ccc;}div.gallery:hover {border: 1px solid #777;}div.gallery img {width: 100%;height: auto;}* {box-sizing: border-box;}.responsive {padding: 0 6px;float: left;width: 24.99999%;}</style></head><body>"

	resp, _ := http.Get("https://yts.mx/rss/0/720p/all/0/all")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	xml.Unmarshal(bytes, &s)

	for idx, _ := range s.Movetitles {
		out += "<div class='responsive'><div class='gallery'><h2>" + s.Movetitles[idx] + "</h2>" + s.Movedescriptions[idx] + "<p>" + s.MovepubDates[idx] + "</p></div></div>"
	}
	out += "</body></html>"
	fmt.Fprintf(w, out)
}

func main(){
	http.HandleFunc("/", index_hander)
	http.ListenAndServe(":8000", nil)
}