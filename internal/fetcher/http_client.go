package fetcher

import "net/http"

func Client(url, name string){

	resp, err := http.Get(url)
	if err != nil{
		panic(err)
	}

	defer resp.Body.Close()



}