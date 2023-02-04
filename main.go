package main

import "log"

func main() {
	config, err := LoadConfig("./env")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	search(config.CustomSearchApiKey, config.CustomSearchCx, "The Witcher 3 Wild Hunt")
}
