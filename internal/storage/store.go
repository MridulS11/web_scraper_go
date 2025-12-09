package storage

import (
	"bufio"
	"os"
	"path/filepath"
	//"runtime"
)

func Store(url string){

	d := []byte(url)
	path, err := os.Getwd()	// for the path of the module

	// _, path, _, ok := runtime.Caller(0)

	// if !ok{
	// 	panic(err)
	// }

	if err != nil{
		panic(err)
	}

	fullpath := filepath.Join(path, "url.txt")

	file, err := os.OpenFile(fullpath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)

	if err != nil{
		panic(err)
	}

	w := bufio.NewWriter(file)
	w.WriteString(string(d) + "\n")
	w.Flush()

	defer file.Close()

}