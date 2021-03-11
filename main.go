package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {

	fmt.Println("Creating Armplay!")

	//Load .env File
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Failure loading .env: %v", err)
	}

	botToken := os.Getenv("TOKEN")

	discord, err := discordgo.New("Bot " + botToken)
	if err != nil {
		log.Fatal("Incorrect Credentials: %v", err)
	}

}

var 
func main() {

	fmt.Println("Hello World!")

}
