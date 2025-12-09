package fetcher

import (
	"bufio"
	"os"
	"path/filepath"
	"web_scraper/internal/storage"
)

func Fetcher(category string){

	base_url := "https://en.wikipedia.org/wiki/"

	url_map := map[string][]string{
		"animals" : {"dogs", "cats", "wolf", "deer", "lion", "tiger"},
		"automobile" : {"audi", "tata", "bmw", "corvette"},
	}

	for _, val := range url_map[category]{
		storage.Store(base_url + val)
	}

	path, err := os.Getwd()
	if err != nil{
		panic(err)
	}

	fullpath := filepath.Join(path, "url.txt")
	file, err := os.Open(fullpath)

	if err != nil{
		panic(err)
	}

	

	w := bufio.NewScanner(file)
	
	for w.Scan(){
		for _, page := range url_map[category]{
			go Client(w.Text(), page)
		}
	}


}