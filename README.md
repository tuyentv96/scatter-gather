#scatter-gather

A simple implementation of scatter gather pattern in Go.

## Usage

```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	scattergather "github.com/tuyentv96/scatter-gather"
)

func main() {
	urls := []string{
		"https://jsonplaceholder.typicode.com/todos/1",
		"https://jsonplaceholder.typicode.com/todos/2",
		"https://jsonplaceholder.typicode.com/todos/3",
		"https://jsonplaceholder.typicode.com/todos/4",
		"https://jsonplaceholder.typicode.com/todos/5",
		"https://jsonplaceholder.typicode.com/todos/6",
		"https://jsonplaceholder.typicode.com/todos/7",
		"https://jsonplaceholder.typicode.com/todos/8",
		"https://jsonplaceholder.typicode.com/todos/9",
	}

	result, err := scattergather.ScattergatherWithInputParams(urls, 4, func(params []string) ([]string, error) {
		rs := make([]string, 0, len(params))
		for _, url := range params {
			fmt.Printf("Fetching url: %s\n", url)
			resp, err := http.Get(url)
			if err != nil {
				return nil, err
			}

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			rs = append(rs, string(body))
			fmt.Printf("Finished url: %s\n", url)
		}

		return rs, nil
	})
	if err != nil {
		fmt.Printf("Err: %+v\n", err)
	}

	fmt.Printf("result size: %d\n", len(result))
}
```
