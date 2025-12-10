package main

import (
	"os"
	"os/exec"
	"path/filepath"
)

func main(){
	category := os.Args[1]
	path, _ := os.Getwd()
	fullpath := filepath.Join(path, "main "+category)
	binary := exec.Command(fullpath)
	if binary.Err != nil{
		panic("Panicked")
	}
}