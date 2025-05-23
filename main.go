package main

import (
	"github.com/fatih/color"
	"github.com/hugolgst/rich-go/client"
	"time"
)

var clientID string

func main() {
	color.Blue("Started West Ham RPC")

	clientID = "1256283784245153946"

	err := client.Login(clientID)
	if err != nil {
		panic(err)
	}

	color.Green("Successfully logged in as %s", clientID)

	lyrics := [8]string{"I'm forever blowing bubbles,", "Pretty bubbles in the air,", "They fly so high, nearly reach the sky,", "Then like my dreams they fade and die.", "Fortune's always hiding,", "I've looked everywhere,", "I'm forever blowing bubbles,", "Pretty bubbles in the air."}

	magenta := color.New(color.FgCyan)

	for {
		for i := 0; i < len(lyrics); i++ {
			_, err := magenta.Printf("\rCounter: %d", i)
			if err != nil {
				return
			}
			err = client.SetActivity(client.Activity{
				State:      lyrics[i],
				LargeImage: "hammers",
				LargeText:  "⚒️ Come On You Irons ⚒️",
			})

			if err != nil {
				panic(err)
			}
			time.Sleep(11 * time.Second)
		}
	}
}
