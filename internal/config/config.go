package config

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

// Config holds all configuration for the bot
type Config struct {
	Token    string   `env:"DISCORD_TOKEN,required"`
	GuildID  string   `env:"GUILD_ID"`         // Optional: for testing commands in specific guild
	OwnerIDs []string `env:"BOT_OWNER_IDS"`    // Bot owner Discord IDs (comma-separated)
	AdminIDs []string `env:"BOT_ADMIN_IDS"`    // Bot admin Discord IDs (comma-separated)
}

// Load returns configuration from environment variables
func Load() (*Config, error) {
	var cfg Config
	if err := envconfig.Process(context.Background(), &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
