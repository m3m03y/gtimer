package main

import (
	"fmt"
	"log"

	"github.com/m3m03y/gtimer/pkg/config"
	"github.com/m3m03y/gtimer/pkg/search"
)

func getLinks() {
	config, err := config.LoadConfig("./env")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	links := search.SearchLinks(config.CustomSearchApiKey, config.CustomSearchCx, "The Witcher 3 Wild Hunt")
	fmt.Print(links)
}

func main() {

	fmt.Print(search.GetGameData("https://howlongtobeat.com/game/10270"))
}
