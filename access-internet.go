package main

import(
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/xml"
)

type Sitemapindex struct {
	Locations []Location `xml:"sitemap"`
}

type Location struct {
	Loc string `xml:"loc"`
}

func (l Location) String () string {
  return fmt.Sprintf(l.Loc)
}

func main(){
	resp, _ := http.Get("https://www.washingtonpost.com/sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	//resp_body := string(bytes)
	//fmt.Println(resp_body)
	resp.Body.Close()

	var s Sitemapindex
	xml.Unmarshal(bytes, &s)
	//fmt.Println(s.Locations)

	for _, val := range s.Locations {
		fmt.Printf("%s",val)
	}
}