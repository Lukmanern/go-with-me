package main

import (
	"fmt"
	"net/http"
)

func main() {
	escapeList := []string{
		"%00",
		"%0a",
		"%0a%20",
		"%0d",
		"\r\t",
	}
	target := "http://ajakan.xyz/?data="

	for _, escape := range escapeList {
		url := target + escape
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Printf("[ERR] %s\n", url)
			continue
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("[ERR] %s\n", url)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			fmt.Printf("[ERR] %s\n", url)
		} else {
			fmt.Printf("[OK] %s\n", url)
		}
	}

}
