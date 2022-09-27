package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)


func main() {
	response, err := http.Get("https://raider.io/api/v1/characters/profile?region=eu&realm=Stormscale&name=Kixtotem&fields=mythic_plus_scores")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

  fmt.Println(responseObject)

}