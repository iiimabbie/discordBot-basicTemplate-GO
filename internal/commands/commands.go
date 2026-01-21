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

// ============================================
// Component Handlers (Buttons, Select Menus)
// ============================================

// componentHandlers stores handlers for message components (buttons, select menus)
var componentHandlers = map[string]Handler{
	"ping_reload": PingReloadHandler,
	// Add your component handlers here:
	// "button_custom_id": ButtonClickHandler,
}

// RegisterComponentHandler registers a handler for a component custom ID
func RegisterComponentHandler(customID string, handler Handler) {
	componentHandlers[customID] = handler
}

// GetComponentHandlers returns all component handlers
func GetComponentHandlers() map[string]Handler {
	return componentHandlers
}
