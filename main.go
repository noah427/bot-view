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
	ln, _ := net.Listen("tcp", "127.0.0.1:3000")
	go http.Serve(ln, http.FileServer(http.FS(fs)))

	ui, _ := lorca.New("", "", 480, 320)
	defer ui.Close()

	ui.Bind("guilds", func() []*discordgo.UserGuild {
		guilds, err := Client.UserGuilds(100, "", "")
		if err != nil {
			fmt.Println(err)
		}
		return guilds
	})

	loadDiscord()

	ui.Load("http://127.0.0.1:3000/public")

	<-ui.Done()
	Client.Close()
}
