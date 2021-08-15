package main

import (
	"os"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var Client *discordgo.Session

func loadDiscord(){
	var err error
	Client, err = discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil{
		fmt.Println("invalid token")
	}

	Client.StateEnabled = true

	err = Client.Open()
	if err != nil {
		fmt.Println(err)
	}
}
