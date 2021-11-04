package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Response struct {
	Updates []Update `json:"result"`
}

type Update struct {
	ID      int  `json:"update_id"`
	Message Message `json:"message"`
}

type Message struct {
	ID   int `json:"message_id"`
	Text string `json:"text"`
}

var botToken = ""

func main() {
	flag.StringVar(&botToken, "telegram-token", "", "The telegram bot token without the bot prefix")
	flag.Parse()

	var response Response
	// This should be persistent through reboots
	var lastSeenUpdate int
	client := http.Client{
		Timeout: 15 * time.Minute,
	}

	httpResp, err := client.Get(fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates", botToken))
	if err != nil {
		panic(err)
	}


	if err := json.NewDecoder(httpResp.Body).Decode(&response); err != nil {
		panic(err)
	}

	if len(response.Updates) > 0 {
		lastSeenUpdate = response.Updates[len(response.Updates) - 1].ID
	}

	log.Println(lastSeenUpdate)


	encoder:= json.NewEncoder(os.Stdout)
	encoder.SetIndent("","  ")
	encoder.Encode(&response.Updates)

}

