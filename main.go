package main

import (
	"context"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/corentings/UCA-discord-bot/commands"
	"github.com/corentings/UCA-discord-bot/database"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
)

var (
	Token   string
	GuildID string
)

func loadVar() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	Token = os.Getenv("TOKEN")
	GuildID = os.Getenv("GUILD_ID")

}

func init() {
	loadVar()
}

func main() {
	// Try to connect to the database
	if err := database.Connect(); err != nil {
		log.Panic("Can't connect database:", err.Error())
	}
	fmt.Println("Connected to database")

	defer func() {
		fmt.Println("Disconnect from database")
		err := database.Mg.Client.Disconnect(context.TODO())
		if err != nil {
			return
		}
	}()

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	commandHandlers := commands.GetCommandHandlers()

	dg.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	myCommands := commands.GetCommands()

	registeredCommands := make([]*discordgo.ApplicationCommand, len(myCommands))
	for i, v := range myCommands {
		cmd, err := dg.ApplicationCommandCreate(dg.State.User.ID, GuildID, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	defer dg.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	if true {
		log.Println("Removing commands...")
		for _, v := range registeredCommands {
			err := dg.ApplicationCommandDelete(dg.State.User.ID, GuildID, v.ID)
			if err != nil {
				log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
			}
		}
	}

	log.Println("Gracefully shutting down.")
}
