package fetcher

import (
	"bufio"
	"context"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
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

	page_count := len(url_map[category])
	var wg sync.WaitGroup
	jobs := make(chan string, page_count)

	for i := 0 ; i < page_count ; i++{
		//wg.Add(1)
		wg.Go(func(){
			for job := range jobs{	//job is the value received from the channel, so job has type string.
				ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
				Client(ctx, job)
				if ctx.Err() != nil{
					log.Println(ctx.Err())
				}
				// select {
				// case <-jobs:
				// 	Client(job)
				// 	fmt.Println(job)		this consumes the job too, which causes less urls to be accessed(serious bug)
				// case <- ctx.Done():
				// 	fmt.Println(ctx.Err())
				// }
				cancel()
			}
		})
	}

	for w.Scan(){
		//defer wg.Done()
		jobs <- w.Text()
	}

	//defer os.Remove(fullpath)

	close(jobs) //needs to be here before wait
	wg.Wait()

}