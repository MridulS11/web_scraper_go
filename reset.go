package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main(){
	
	file, _ := os.Getwd()

	cleanType := ""

	if len(os.Args) > 1{
		cleanType = os.Args[1]
	}

	mainPath := filepath.Join(file, "main")
	urlPath := filepath.Join(file, "url.txt")
	scrapedPath := filepath.Join(file, "scraped_text.txt")

	if cleanType == "-deep"{
		os.Remove(mainPath)
	}

	os.Remove(urlPath)
	os.Remove(scrapedPath)
	fmt.Println("Directory Cleaned")

}