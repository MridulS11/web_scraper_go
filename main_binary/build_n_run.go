package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"time"
	"syscall"
)

func main(){
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	category := os.Args[1]
	path, _ := os.Getwd()
	fullpath := filepath.Join(path, "main")
	fmt.Println(fullpath)
	build := exec.Command("go", "build", "cmd/main.go")
	build.Run()
	binary := exec.Command(fullpath, category)
	binary.Stdout = os.Stdout
	binary.Stderr = os.Stderr
	fmt.Println("Running Your Scraper...")
	binary.Run()
	if binary.Err != nil{
		panic("Panicked")
	}
	fmt.Println("Output Files will be deleted in 30 seconds, press ctrl + c to interrupt the deletion")
	select{
	case <-time.After(30 * time.Second):
		fmt.Println("\nDeleting Files...")
		end := exec.Command(filepath.Join(path, "reset"), "-deep")
		end.Stdout = os.Stdout
		end.Stderr = os.Stderr
		end.Run()
	case <- sigs:
		fmt.Println("\nFiles Saved")
	}
}