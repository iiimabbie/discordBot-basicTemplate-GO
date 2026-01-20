package commands

import (
	"time"

	"discord-bot-template/internal/embed"

	"github.com/bwmarrin/discordgo"
)

// PingCommand definition
var PingCommand = &discordgo.ApplicationCommand{
	Name:        "ping",
	Description: "Check bot latency and response time",
}

// PingHandler handles the /ping command
func PingHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Record start time for roundtrip calculation
	start := time.Now()

	// Send initial "thinking" response
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})
	if err != nil {
		return
	}

	// Calculate latencies
	roundtrip := time.Since(start)
	wsLatency := s.HeartbeatLatency()

	// Create response embed
	pingEmbed := embed.Ping(roundtrip, wsLatency)

	// Edit the deferred response with actual content
	_, err = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Embeds: &[]*discordgo.MessageEmbed{pingEmbed},
	})
	if err != nil {
		// Log error if needed
		return
	}
}
