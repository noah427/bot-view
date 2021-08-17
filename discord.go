package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/bwmarrin/discordgo"
)

var Client *discordgo.Session
var Voice *discordgo.VoiceConnection

func loadDiscord() bool {
	var err error
	Client, err = discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		fmt.Println("Invalid token")
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

func disconnectVoice() {
	Voice.Disconnect()
}

func displayFiles() []string {
	files,_ := ioutil.ReadDir("audio")
	var names = []string{}
	for _,file := range files {
		names = append(names, file.Name())
	}
	return names
}

func recieveFile(fileName string) {
	var buffer = make([][]byte, 0)
	file, err := os.Open(fileName)
	if err != nil {
		return
	}

	var opuslen int16

	for {

		err = binary.Read(file, binary.LittleEndian, &opuslen)

		if err == io.EOF || err == io.ErrUnexpectedEOF {
			err := file.Close()
			if err != nil {
				break
			}
			break
		}

		if err != nil {
			break
		}

		InBuf := make([]byte, opuslen)
		err = binary.Read(file, binary.LittleEndian, &InBuf)

		if err != nil {
			break
		}

		buffer = append(buffer, InBuf)
	}

	Voice.Speaking(true)

	for _, buff := range buffer {
		Voice.OpusSend <- buff
	}

	Voice.Speaking(false)

}
