package gopscraper

import (
	"encoding/json"
	"log"
	"strings"
)

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

// "http://activa.sapo.pt/passatempos/" => "http://activa.sapo.pt"
func getURLroot(url string) string {
	splitted_url := strings.SplitAfter(url, "/")
	root_url := strings.Join(splitted_url[0:3],"")
	return strings.TrimSuffix(root_url, "/")
}
