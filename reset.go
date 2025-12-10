package webscraper

import (
	"fmt"
	"os"
	"path/filepath"
)

func reset(){

	file, _ := os.Getwd()
	mainPath := filepath.Join(file, "main")
	urlPath := filepath.Join(file, "url.txt")

	os.Remove(mainPath)
	os.Remove(urlPath)
	
}

func main(){
	reset()
	fmt.Println("Directory Cleaned")
}