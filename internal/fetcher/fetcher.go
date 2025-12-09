package fetcher

import (
	"bufio"
	"os"
	"path/filepath"
	"sync"
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

	var wg sync.WaitGroup
	page_count := len(url_map[category])
	jobs := make(chan string, page_count)

	for i := 0 ; i < page_count ; i++{
		go func(){
			for job := range jobs{
				Client(job)
				wg.Done()
			}
		}()
	}

	for w.Scan(){
		for _, page := range url_map[category]{
			jobs <- w.Text() + page
		}
	}

	wg.Wait()
	close(jobs)

}