package main

import(
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/xml"
	"html/template"
)

type Moveindexs struct {
	Movetitles []string `xml:"channel>item>title"`
	Movedescriptions []string `xml:"channel>item>description"`
	MovepubDates []string `xml:"channel>item>pubDate"`
}

type Movelist struct {
	Movetitles string
	Movedescriptions string
	MovepubDates string
}

type htmlPage struct {
	Title string
	New_movelist map[int]Movelist
}

func index_hander(w http.ResponseWriter, r *http.Request) {
	var s Moveindexs
	New_movelist := make(map[int]Movelist)

	resp, _ := http.Get("https://yts.mx/rss/0/720p/all/0/all")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	xml.Unmarshal(bytes, &s)

	for idx, _ := range s.Movetitles {
		New_movelist[idx] = Movelist{s.Movetitles[idx], s.Movedescriptions[idx], s.MovepubDates[idx]}
	}

	p := htmlPage{Title: "", New_movelist: New_movelist}
	t, _ := template.ParseFiles("movie-db.html")
	fmt.Println(t.Execute(w, p))
}

func main(){
	http.HandleFunc("/", index_hander)
	http.ListenAndServe(":8000", nil)
}