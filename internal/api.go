package api

import (
	"fmt"
	"io"
	"net/http"
)

func FetchLocation() {
	resp, err := http.Get("https://pokeapi.co/api/v2/location?offset=0&limit=20")
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Response: %s\n", body)
}
