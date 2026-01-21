package commands

import (
	"time"

	"discord-bot-template/internal/component"
	"discord-bot-template/internal/embed"

	"github.com/bwmarrin/discordgo"
)

func init() {
	// 自動註冊指令
	RegisterCommand(pingCommand, PingHandler)

	// 自動註冊按鈕
	RegisterComponent("ping_reload", PingReloadHandler)
}

var pingCommand = &discordgo.ApplicationCommand{
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

	// Create reload button
	row := component.ReloadButtonRow("ping_reload")

	// Edit the deferred response with actual content
	_, err = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Embeds:     &[]*discordgo.MessageEmbed{pingEmbed},
		Components: &[]discordgo.MessageComponent{row},
	})
	if err != nil {
		// Log error if needed
		return
	}
}

// PingReloadHandler handles the reload button click
func PingReloadHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	start := time.Now()

	// Defer update to show loading state
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredMessageUpdate,
	})
	if err != nil {
		return
	}

	// Calculate new latencies
	roundtrip := time.Since(start)
	wsLatency := s.HeartbeatLatency()

	// Create updated embed
	pingEmbed := embed.Ping(roundtrip, wsLatency)

	// Keep the reload button
	row := component.ReloadButtonRow("ping_reload")

	// Update the message
	_, err = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Embeds:     &[]*discordgo.MessageEmbed{pingEmbed},
		Components: &[]discordgo.MessageComponent{row},
	})
	if err != nil {
		return
	}
}
