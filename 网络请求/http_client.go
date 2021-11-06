package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	/* 	resp, err := http.Get("http://localhost:8081/select/all")
	   	if err != nil {
	   		return
	   	}
	   	fmt.Println(resp) */
	resp, err := http.Get("http://localhost:8081/select/all")
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}
