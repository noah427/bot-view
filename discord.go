package main

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

var Client *discordgo.Session
var Voice *discordgo.VoiceConnection

func loadDiscord() bool {
	var err error
	Client, err = discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		fmt.Println("invalid token")
		return false
	}

	Client.StateEnabled = true

	err = Client.Open()
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func joinVoice(guild, channel string) {
	var err error
	Voice, err = Client.ChannelVoiceJoin(guild, channel, false, false)
	if err != nil {
		fmt.Println(err)
	}
}

func disconnectVoice(){
	Voice.Disconnect()
}
