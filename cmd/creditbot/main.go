package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error in loading .env file: %v", err)
	}
	log.Printf("Success in loading .env file")
	botToken := os.Getenv("DISCORD_BOT_TOKEN")
	if len(botToken) < 1 {
		log.Fatalf("Error with bot token")
	}
	log.Printf("Success with bot token")
	db, err := discordgo.New("Bot " + botToken)
	if err != nil {
		log.Fatalf("Error with initializing bot: %v", err)
	}
	log.Printf("Bot started up successfully")
	// Open Websocket Connection to Discord
	err = db.Open()
	if err != nil {
		log.Fatalf("Error with opening session: %v", err)
	}
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	log.Println("Shutting down...")
	db.Close()
	log.Println("Discord session closed.")
}
