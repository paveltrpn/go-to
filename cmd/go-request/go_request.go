package main

import (
	"fmt"
	"net/http"
	"time"
)

const adress = "https://httpbin.org/bytes/32"

func main() {
	client := http.Client{
		Timeout: time.Second * 2,
	}
	resp, err := client.Get(adress)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("get - %v status is - %v\n", adress, resp.StatusCode)

	for _, field := range resp.Header {
		fmt.Println(field)
	}

	for {
		bs := make([]byte, 1014)
		n, err := resp.Body.Read(bs)
		fmt.Printf("%v", bs[:n])

		if n == 0 || err != nil {
			break
		}
	}
}
