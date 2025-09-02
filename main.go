package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "http://www.thepaper.cn"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("fetch url, error:", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		fmt.Println("status code", resp.StatusCode)
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read body, error:", err)
	}
	fmt.Println(string(body))
}
