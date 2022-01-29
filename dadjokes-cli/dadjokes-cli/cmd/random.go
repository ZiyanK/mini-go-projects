/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// randomCmd represent the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Get a random dad joke.",
	Long:  `Dad jokes cli is a tool that gives you random dad jokes in your terminal`,
	Run: func(cmd *cobra.Command, args []string) {
		getRandomJoke()
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
}

type Joke struct {
	id     string `json:id`
	Joke   string `json:joke`
	Status int    `json:int`
}

func getRandomJoke() {
	url := "https://icanhazdadjoke.com/"
	responseBytes := getJokeData(url)
	joke := Joke{}
	if err := json.Unmarshal(responseBytes, &joke); err != nil {
		log.Printf("Could not unmarshall reponse - %v", err)
	}
	fmt.Println(string(joke.Joke))
}

func getJokeData(url string) []byte {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Printf("Could not request a dad joke - %v", err)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "For a cli tool")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("Could make the request - %v", err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Could not read the response body - %v", err)
	}

	return responseBytes
}
