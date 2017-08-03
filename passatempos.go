package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"encoding/json"
)

func main() {
	artefactos_contests := getContestsFromPage(
		"artefactos",
		"http://www.arte-factos.net/passatempos-af/",
		"div.titulo-pagina > a",
		getContestsArtefactos,
	)

	transporteslisboa_contests := getContestsFromPage(
		"transporteslisboa",
		"http://passatempos.transporteslisboa.pt/",
		"h1.entry-title > a",
		getContestsTransporteslisboa,
	)

	contests_list := []contestsData{artefactos_contests, transporteslisboa_contests}

	all_contests := mergeMaps(contests_list)
	println(convertToJson(all_contests))
}

type pageContestsData []map[string]string
type contestsData map[string]pageContestsData
type getDataFromContestsElemFunc func(*goquery.Selection, pageContestsData) pageContestsData

func getContestsFromPage(page_name string, page_url string, contests_element_path string, getDataFromContestsElem getDataFromContestsElemFunc) contestsData {
	log.Println("Scraping", page_name)

	doc, err := goquery.NewDocument(page_url)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Website successfully obtained")
	}

	contests_elem := doc.Find(contests_element_path)
	number_contests := contests_elem.Length()

	contests := make(pageContestsData, number_contests)

	contests = getDataFromContestsElem(contests_elem, contests)

	page_contests := make(contestsData)
	page_contests[page_name] = contests

	return page_contests
}

func getContestsArtefactos(contests_elem *goquery.Selection, contests pageContestsData) pageContestsData {
	contests_elem.Each(func(i int, s *goquery.Selection) {
		contest_map := make(map[string]string)
		contest_map["name"] = s.Find("h4").Text()
		contest_map["url"], _ = s.Attr("href")
		contests[i] = contest_map
	})
	return contests
}

func getContestsTransporteslisboa(contests_elem *goquery.Selection, contests pageContestsData) pageContestsData {
	contests_elem.Each(func(i int, s *goquery.Selection) {
		contest_map := make(map[string]string)
		contest_map["name"] = s.Text()
		contest_map["url"], _ = s.Attr("href")
		contests[i] = contest_map
	})
	return contests
}


func convertToJson(to_convert contestsData) string {
	bytes, err := json.Marshal(to_convert)
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}

// Merges all contestsData maps in contestsData into one contestsData map
func mergeMaps(contests_list []contestsData) contestsData {
	final_map := make(contestsData)

	for _, contests_data := range contests_list {
		for k, v := range contests_data {
			final_map[k] = v
		}
	}

	return final_map
}