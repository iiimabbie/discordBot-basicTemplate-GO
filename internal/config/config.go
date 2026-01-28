package config

import (
	"os"
	"strings"
)

// Config holds all configuration for the bot
type Config struct {
	Token   string
	GuildID string // Optional: for testing commands in specific guild
	OwnerIDs []string // Bot owner Discord IDs
	AdminIDs []string // Bot admin Discord IDs
}

// Load returns configuration from environment variables
func Load() *Config {
	return &Config{
		Token:   getEnv("DISCORD_TOKEN", ""),
		GuildID: getEnv("GUILD_ID", ""), // Leave empty to register global commands
		OwnerIDs: parseCSV(getEnv("BOT_OWNER_IDS", "")),
		AdminIDs: parseCSV(getEnv("BOT_ADMIN_IDS", "")),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// parseCSV parses a comma-separated string into a slice
func parseCSV(s string) []string {
	if s == "" {
		return nil
	}
	parts := strings.Split(s, ",")
	result := make([]string, 0, len(parts))
	for _, p := range parts {
		if trimmed := strings.TrimSpace(p); trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}
