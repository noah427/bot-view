package main

import (
	"fmt"
	// "io"
	"github.com/bwmarrin/dgvoice"
	"io/ioutil"
	"os"
	// "time"
	// "strings"
	// "encoding/binary"

	"github.com/bwmarrin/discordgo"
	// "github.com/jonas747/dca"
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
	files, _ := ioutil.ReadDir("audio")
	var names = []string{}
	for _, file := range files {
		names = append(names, file.Name())
	}
	return names
}

func recieveFile(fileName string) {
	Voice.Speaking(true)
	stopped := make(<-chan bool)
	dgvoice.PlayAudioFile(Voice, fmt.Sprintf("./audio/%s", fileName), stopped)
	<-stopped
	Voice.Speaking(false)
}

// func recieveFile(fileName string) {
// 	file,err := os.OpenFile(fmt.Sprintf("./audio/%s",fileName),os.O_RDONLY,0777)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	decoder := dca.NewDecoder(file)

// 	for {
// 		frame, err := decoder.OpusFrame()
// 		if err != nil {
// 			if err != io.EOF {
// 				fmt.Println("unhandled error")
// 				break
// 			}

// 			break
// 		}

// 		select {
// 		case Voice.OpusSend <- frame:
// 		case <-time.After(time.Second):
// 			return
// 		}
// 	}
// }

// func recieveFile(fileName string) {
// 	var buffer = make([][]byte, 0)

// 	if strings.Split(fileName, ".")[1] != "dca" {
// 		encodeSession, err := dca.EncodeFile(fmt.Sprintf("./audio/%s",fileName), dca.StdEncodeOptions)
// 		if err != nil {
// 			fmt.Println(err, "92")
// 		}
// 		defer encodeSession.Cleanup()
// 		output, _ := os.Create(fmt.Sprintf("./audio/%s.dca",strings.Split(fileName, ".")[0]))
// 		io.Copy(output, encodeSession)
// 	}

// 	file, err := os.Open(fmt.Sprintf("./audio/%s.dca",strings.Split(fileName, ".")[0]))
// 	if err != nil {
// 		fmt.Println(err, "101")
// 		return
// 	}

// 	var opuslen int16

// 	for {

// 		err = binary.Read(file, binary.LittleEndian, &opuslen)

// 		if err == io.EOF || err == io.ErrUnexpectedEOF {
// 			err := file.Close()
// 			if err != nil {
// 				break
// 			}
// 			break
// 		}

// 		if err != nil {
// 			break
// 		}

// 		if opuslen < 1 {
// 			return
// 		}
// 		InBuf := make([]byte, opuslen)
// 		err = binary.Read(file, binary.LittleEndian, &InBuf)

// 		if err != nil {
// 			break
// 		}

// 		buffer = append(buffer, InBuf)
// 	}

// 	Voice.Speaking(true)

// 	for _, buff := range buffer {
// 		Voice.OpusSend <- buff
// 	}

// 	Voice.Speaking(false)

// }
