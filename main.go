package main

import "github.com/bwmarrin/discordgo"

func main() {

	discordClient, err := discordgo.New("Bot" + "Authentication Token")
	if err != nil{
		panic(err)
	}

}
