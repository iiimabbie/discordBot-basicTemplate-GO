package commands

import (
	"github.com/bwmarrin/discordgo"
)

// Handler is a function that handles an interaction
type Handler func(s *discordgo.Session, i *discordgo.InteractionCreate)

// Command represents a slash command with its definition and handler
type Command struct {
	Definition *discordgo.ApplicationCommand
	Handler    Handler
}

// AllCommands returns all available commands
func AllCommands() []*Command {
	return []*Command{
		{Definition: PingCommand, Handler: PingHandler},
		// Add your commands here:
		// {Definition: YourCommand, Handler: YourHandler},
	}
}

// GetDefinitions returns all command definitions (for registration)
func GetDefinitions() []*discordgo.ApplicationCommand {
	commands := AllCommands()
	definitions := make([]*discordgo.ApplicationCommand, len(commands))
	for i, cmd := range commands {
		definitions[i] = cmd.Definition
	}
	return definitions
}

// GetHandlers returns a map of command names to handlers
func GetHandlers() map[string]Handler {
	commands := AllCommands()
	handlers := make(map[string]Handler)
	for _, cmd := range commands {
		handlers[cmd.Definition.Name] = cmd.Handler
	}
	return handlers
}
