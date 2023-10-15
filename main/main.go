package main

import(
	"fmt"
	"net/http"
	"io"
)

func main() {
	url := "https://api.gameofthronesquotes.xyz/v1/random"
	response, err := http.Get(url)

	if err != nil {
		fmt.Println("Error making http request: %s\n", err)
	}
	
	body, err := io.ReadAll(response.Body)
	defer response.Body.Close()

	fmt.Println(string(body))
}
