package main

import (
	//"fmt"
	"os"
	"web_scraper/internal/fetcher"
)

func main(){

	category := os.Args[1]
	fetcher.Fetcher(category)
	
}