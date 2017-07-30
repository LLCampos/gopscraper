package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"encoding/json"
)

func main() {
	println(getContestsFromPage(
		"artefactos",
		"http://www.arte-factos.net/passatempos-af/",
		"div.titulo-pagina > a",
		getContestsArtefactos,
	))

	println(getContestsFromPage(
		"transporteslisboa",
		"http://passatempos.transporteslisboa.pt/",
		"h1.entry-title > a",
		getContestsTransporteslisboa,
	))
}


type getDataFromContestsElemFunc func(*goquery.Selection, []map[string]string) []map[string]string

func getContestsFromPage(page_name string, page_url string, contests_element_path string, getDataFromContestsElem getDataFromContestsElemFunc) string {
	log.Println("Scraping", page_name)

	doc, err := goquery.NewDocument(page_url)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Website successfully obtained")
	}

	contests_elem := doc.Find(contests_element_path)
	number_contests := contests_elem.Length()

	contests := make([]map[string]string, number_contests)

	contests = getDataFromContestsElem(contests_elem, contests)

	contests_map := make(map[string][]map[string]string)
	contests_map[page_name] = contests

	return convertToJson(contests_map)
}

func getContestsArtefactos(contests_elem *goquery.Selection, contests []map[string]string) []map[string]string {
	contests_elem.Each(func(i int, s *goquery.Selection) {
		contest_map := make(map[string]string)
		contest_map["name"] = s.Find("h4").Text()
		contest_map["url"], _ = s.Attr("href")
		contests[i] = contest_map
	})
	return contests
}

func getContestsTransporteslisboa(contests_elem *goquery.Selection, contests []map[string]string) []map[string]string {
	contests_elem.Each(func(i int, s *goquery.Selection) {
		contest_map := make(map[string]string)
		contest_map["name"] = s.Text()
		contest_map["url"], _ = s.Attr("href")
		contests[i] = contest_map
	})
	return contests
}


func convertToJson(to_convert map[string][]map[string]string) string {
	bytes, err := json.Marshal(to_convert)
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}
