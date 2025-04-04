package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/JonHunt1995/DegreeProgressTracker.git/internal/config"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func getCurrentNickname(m *discordgo.MessageCreate) string {
	if m.Member == nil {
		return ""
	}

	return m.Member.Nick
}

func updateNickname(s *discordgo.Session, guildID, userID, newNickname string) error {
	err := s.GuildMemberNickname(guildID, userID, newNickname)
	if err != nil {
		log.Printf("Error on updating nickname: %v", err)
		return err
	}
	log.Printf("Success in Changing Nickname to %v", newNickname)
	return nil
}

func main() {
	// Load .env File
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error in loading .env file: %v", err)
	}
	log.Printf("Success in loading .env file")
	// Load Config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Unable to log config due to %v", err)
	}
	//Initialize Bot
	botToken := os.Getenv("DISCORD_BOT_TOKEN")
	if len(botToken) < 1 {
		log.Fatalf("Error with bot token")
	}
	log.Printf("Success with bot token")

	if len(cfg.DiscordID) == 0 {
		log.Printf("Error with loading user id")
	}
	db, err := discordgo.New("Bot " + botToken)
	if err != nil {
		log.Fatalf("Error with initializing bot: %v", err)
	}
	// Handle Message
	messageHandler := func(s *discordgo.Session, m *discordgo.MessageCreate) {
		// Check To See If Command Called Correctly
		if len(m.Content) < 2 || !strings.HasPrefix(m.Content, "!") {
			return
		}
		if m.GuildID == "" {
			log.Printf("Attempted command without being in server")
			s.ChannelMessageSend(m.ChannelID, "The `!updatecredits` command only works in servers")
			return
		}
		if m.Author.ID != cfg.DiscordID {
			log.Printf("Caller of command not authorized to use command")
			return
		}

		command := strings.ToLower(strings.Fields(strings.TrimPrefix(m.Content, "!"))[0])
		if command != "updatecredits" {
			return
		}
		log.Printf("Using command %v in server %v", command, m.GuildID)
		// Execute Command
		numCredits, err := config.GetGoogleSheet(cfg)
		if err != nil {
			log.Printf("Error in loading credits: %v", err)
			s.ChannelMessageSend(m.ChannelID, "Error in loading credits")
			return
		}
		// Obtain Current Nickname
		currentNickname := getCurrentNickname(m)
		// Changing Nickname
		newNickname := fmt.Sprintf("GenTech34 [BSCS %v/117]", numCredits)
		if len(newNickname) > 32 {
			newNickname = newNickname[:32]
			s.ChannelMessageSend(m.ChannelID, "nickname too long: truncated to first 32 letters")
		}
		log.Printf("Attempting to change nickname for user %v on guild %v from %v to %v",
			m.Author.Username,
			m.GuildID,
			currentNickname,
			newNickname)

		//err = updateNickname(s, m.GuildID, m.Author.ID, newNickname)
		log.Printf("TEST: Bot's own User ID (s.State.User.ID): %s", s.State.User.ID)
		log.Printf("TEST: Attempting to change BOT's own nickname...")
		err = updateNickname(s, m.GuildID, s.State.User.ID, "BotNickTest")
		if err != nil {
			log.Printf("Error in attempting to change nickname: %v", err)
			s.ChannelMessageSend(m.ChannelID, "Error in attempting to change nickname")
			return
		}
		successMessage := fmt.Sprintf("Successfully changed nickname from %v to %v!",
			currentNickname,
			newNickname)
		s.ChannelMessageSend(m.ChannelID, successMessage)
	}
	log.Printf("Bot started up successfully")
	// Declare Intents
	db.Identify.Intents = (discordgo.IntentsGuilds |
		discordgo.IntentsGuildMessages |
		discordgo.IntentMessageContent)
	log.Printf("Intents declared...")
	db.AddHandler(messageHandler)
	log.Printf("Message handler registered!")
	// Open Websocket Connection to Discord
	err = db.Open()
	if err != nil {
		log.Fatalf("Error with opening session: %v", err)
	}
	// Keep Connection Running Until Kill Signal, then Close Gracefully
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	log.Println("Shutting down...")
	db.Close()
	log.Println("Discord session closed.")
}
