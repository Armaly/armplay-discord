package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var session *discordgo.Session

func init() {

	fmt.Println("Creating Armplay!")

	//Load .env File
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Failure loading .env: %v", err)
	}

	botToken := os.Getenv("TOKEN")

	session, err := discordgo.New("Bot " + botToken)
	if err != nil {
		log.Fatal("Incorrect Credentials: %v", err)
	}

}

var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name: "test-command",
			Description: "test description",
		},
		{
			Name: "options",
			Description: "Command options",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type: discordgo.ApplicationCommandOptionString,
					Name: "option-1",
					Description: "testing-inner",
					Required: true,
				},
			},
		},

	}
)

func main() {

	fmt.Println("Hello World!")

	session.AddHandler(func(session *discordgo.Session, r* discordgo.Ready){
		fmt.Println("Armplay is up.")
	})

	err := session.Open()
	if err != nil  {
		log.Fatalf("Armplay can't open: %v", err)
	}

	for _, v := range commands {
		_, err := session.ApplicationCommandCreate(session.State.User.ID, *GuildID, v)
		if err != nil  {
			log.Fatalf("Can't create '%v' command: %v", v.Name, err)
		}
	}

	defer session.Close()

	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Armplay is shutting down")
}
