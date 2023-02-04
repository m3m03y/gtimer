package search

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	customsearch "google.golang.org/api/customsearch/v1"
	"google.golang.org/api/googleapi/transport"
)

type Game struct {
	title string
	link  string
}

func SearchLinks(apiKey string, cx string, query string) (linksLibrary []Game) {
	client := &http.Client{Transport: &transport.APIKey{Key: apiKey}}

	svc, err := customsearch.New(client)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := svc.Cse.List().Cx(cx).Q(query).Do()
	if err != nil {
		log.Fatal(err)
	}

	for _, result := range resp.Items {
		game := Game{result.Title, result.Link}
		linksLibrary = append(linksLibrary, game)
	}
	return linksLibrary
}

func GetGameData(url string) (resp *http.Response) {
	// resp, getErr := http.Get(url)
	// if getErr != nil {
	// 	log.Fatal(getErr)
	// }

	// return resp

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36")

	resp, err = client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%s\n", body)
	return resp
}
