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

	mainPath := filepath.Join(file, "cmd/main")
	fmt.Println(mainPath)
	urlPath := filepath.Join(file, "url.txt")
	scrapedPath := filepath.Join(file, "scraped_text.txt")

	if cleanType == "-deep"{
		err := os.Remove(mainPath)
		if err != nil{
			fmt.Println(err)
		}
	}
	os.Remove(urlPath)
	os.Remove(scrapedPath)
	if cleanType == "-deep"{
		fmt.Println("Directory Deep Cleaned")
	} else{
		fmt.Println("Directory Cleaned")
	}
	

}