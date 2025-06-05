package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func FetchApi(url string, object any) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if res.StatusCode > 299 {
		fmt.Println("Error:", res.Status)
		res.Body.Close()
		os.Exit(1)
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)

	if err := decoder.Decode(object); err != nil {
		fmt.Println("Failed to decode JSON response:", err)
		res.Body.Close()
		os.Exit(1)
	}
}
