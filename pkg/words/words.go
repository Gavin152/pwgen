package words

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

var requestUrl string = "https://random-word-api.herokuapp.com/word?number="

func GetWords(count int) {

	if count <= 0 {
		count = 1
	}

	url := requestUrl + strconv.Itoa(count)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Printf("error creating http request: %s\n", err)
		os.Exit(1)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}
	
	fmt.Printf("Http call returned code: %d\n", res.StatusCode)
	
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("error reading response body: %s\n", err)
		os.Exit(1)
	}

	var words []string
	err = json.Unmarshal(resBody, &words)
	if err != nil {
		fmt.Printf("error while parsing the response body: %s\n", err)
	}

	fmt.Printf("Http call returned %d words\n", len(words))
	for index, word := range words {
		fmt.Printf("%d. %s\n", index + 1, word)
	}
}
