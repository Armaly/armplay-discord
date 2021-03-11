package main

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo")

func init() {
	fmt.Println("Creating Armplay!")
	discord, err := discordgo.New("Bot " + "authentication token")
	if err != nil {
		log.Fatal("Incorrect Credentials: %v", err)
	}

}
func main() {

	fmt.Println("Hello World!")

	

}
