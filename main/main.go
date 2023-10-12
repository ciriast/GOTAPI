package main

import(
	"fmt"
	"net/http"
)

func main() {
	url := "https://api.gameofthronesquotes.xyz/v1/random"
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(req)
}


