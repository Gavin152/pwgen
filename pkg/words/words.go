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

func GetWords(count int, seperator string) string {

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

	outString := ""
	
	//fmt.Print(words)
	for index, word := range words {
	//	fmt.Printf("Adding %s to %s", word, outString)
		outString += word
		if index + 1 < len(words) {
			outString += string(seperator)
		}
	//	fmt.Printf("Resulting in %s", outString)
	}

	return outString
}
