package main

import (
	"embed"
	"fmt"
	"net"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/zserge/lorca"
)

//go:embed public
var fs embed.FS

func main() {
	godotenv.Load()
	ln, _ := net.Listen("tcp", "127.0.0.1:4242")
	go http.Serve(ln, http.FileServer(http.FS(fs)))

	ui, _ := lorca.New("", "", 480, 320)
	defer ui.Close()

	ui.Bind("guildsGO", func() []*discordgo.UserGuild {
		guilds, err := Client.UserGuilds(100, "", "")
		if err != nil {
			fmt.Println(err)
		}
		return guilds
	})

	ui.Bind("channelsGO", func(id string) []*discordgo.Channel {
		channels, err := Client.GuildChannels(id)
		if err != nil {
			fmt.Println(err)
		}
		
		return channels
	})

	ui.Bind("messagesGO", func(id string) []*discordgo.Message {
		messages, err := Client.ChannelMessages(id, 25, "", "", "")
		if err != nil {
			fmt.Println(err)
		}

		for i, j := 0, len(messages)-1; i < j; i, j = i+1, j-1 {
			messages[i], messages[j] = messages[j], messages[i]
		}

		return messages
	})

	ui.Bind("sendGO", func(id, content string) bool {
		_, err := Client.ChannelMessageSend(id, content)
		if err != nil {
			fmt.Println(err)
			return false
		}
		return true
	})

	ui.Bind("joinVoiceGO", joinVoice)
	ui.Bind("disconnectVoiceGO",disconnectVoice)

	if !loadDiscord() {
		ui.Load("http://127.0.0.1:4242/public/login")
	} else {
		ui.Load("http://127.0.0.1:4242/public")
	}

	<-ui.Done()
	Client.Close()
}
