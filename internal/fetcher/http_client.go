package fetcher

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func Client(url string){

	//resp, err := http.NewRequest("GET", url, nil)


	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0 Safari/537.36")

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()


	fmt.Println("Received:", resp.Status)

	path, err := os.Getwd()
	if err != nil{
		panic(err)
	}
	fullpath := filepath.Join(path, "scraped_text.txt")
	f, err := os.OpenFile(fullpath, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0777)	//create file(and other privileges)
	if err != nil{
		panic(err)
	}

	pathString := filepath.Base(url)

	// bs := bufio.NewScanner(resp.Body)
	// bs.Buffer(make([]byte, 0, 1024*1024), 1024*1024)	//increased space
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil{
		panic(err)
	}
	br := bufio.NewWriter(f)

	var mu sync.Mutex
	br.WriteString(pathString)

	doc.Find("#mw-content-text p").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		
		mu.Lock()
		if len(text) > 0{
			if _, err := br.WriteString(text); err != nil{
				log.Println("Error Writing Text:", err)
				return
			}
			if _, err := br.WriteString("\n"); err != nil{
				log.Println("Error Writing New Line:", err)
				return
			}
		}
		mu.Unlock()
	})

	// for bs.Scan(){
	// 	br.WriteString(bs.Text()+"\n")
	// }

	// if err := bs.Err(); err != nil{
	// 	panic(err)
	// }

	br.Flush()

}