  
package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"encoding/base64"

    "strings"
	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}


	dg.AddHandler(messageCreate)


	dg.Identify.Intents = discordgo.IntentsGuildMessages


	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

    myString := "https://linkvertise.com/"

    // Option 1: (Recommended)

    if strings.ContainsAny(myString, m.Content) {
        fmt.Println("Message Contains Prefix")
		
		b := strings.Split(m.Content, "=")
		v := strings.Split(b[1],"%")

		fmt.Println(v[0])


		data, err := base64.URLEncoding.DecodeString(v[0])
		if err != nil {
			fmt.Println("error:", err)
		}
		fmt.Println(string(data))



		_, _ = s.ChannelMessageSend(m.ChannelID, string(data))
    } else {
        fmt.Println("Message Does not contain Prefix")
    }






	

}