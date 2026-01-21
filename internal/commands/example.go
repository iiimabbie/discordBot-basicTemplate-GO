package commands

import (
	"time"

	"discord-bot-template/internal/component"
	"discord-bot-template/internal/embed"

	"github.com/bwmarrin/discordgo"
)

func init() {
	// 自動註冊指令
	RegisterCommand(exampleCommand, ExampleHandler)

	// 自動註冊按鈕
	RegisterComponent("example_reload", ExampleReloadHandler)
	RegisterComponent("example_primary", ExampleButtonHandler)
	RegisterComponent("example_secondary", ExampleButtonHandler)
	RegisterComponent("example_success", ExampleButtonHandler)
	RegisterComponent("example_danger", ExampleButtonHandler)
	RegisterComponent("example_emoji", ExampleButtonHandler)
}

var exampleCommand = &discordgo.ApplicationCommand{
	Name:        "example",
	Description: "Show a complete example message with all features",
}

// buildExampleMessage builds the example embed and components
func buildExampleMessage(user *discordgo.User) (*discordgo.MessageEmbed, []discordgo.MessageComponent) {
	// Build embed
	e := embed.New().
		Title("Example Message").
		URL("https://discord.com").
		Description("This is a **complete example** showing all embed features!\n\n" +
			embed.Bold("Bold text") + " | " +
			embed.Italic("Italic text") + " | " +
			embed.InlineCode("code") + "\n" +
			embed.Spoiler("Hidden text") + " | " +
			embed.Mention(user.ID)).
		Color(embed.ColorFuchsia).
		Author(user.Username, "", user.AvatarURL("32")).
		Thumbnail(user.AvatarURL("128")).
		InlineField("Inline Field 1", "Value 1").
		InlineField("Inline Field 2", "Value 2").
		InlineField("Inline Field 3", "Value 3").
		BlockField("Block Field", "This field takes the full width").
		BlockField("Timestamp", embed.RelativeTime(time.Now().Add(-2*time.Hour))+" (2 hours ago)").
		Image("https://cdn-icons-png.flaticon.com/512/5277/5277459.png").
		Footer("Example Footer", "").
		Timestamp().
		Build()

	// Build components
	buttonRow := component.NewActionRow().
		AddButton(component.PrimaryButton("example_primary", "Primary")).
		AddButton(component.SecondaryButton("example_secondary", "Secondary")).
		AddButton(component.SuccessButton("example_success", "Success")).
		AddButton(component.DangerButton("example_danger", "Danger")).
		Build()

	buttonRow2 := component.NewActionRow().
		AddButton(component.NewButton().
			CustomID("example_emoji").
			Label("With Emoji").
			Primary().
			Emoji("🎉").
			Build()).
		AddButton(component.LinkButton("https://discord.com", "Discord Link")).
		Build()

	reloadRow := component.ReloadButtonRow("example_reload")

	return e, []discordgo.MessageComponent{buttonRow, buttonRow2, reloadRow}
}

// ExampleHandler handles the /example command
func ExampleHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	e, components := buildExampleMessage(i.Member.User)

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds:     []*discordgo.MessageEmbed{e},
			Components: components,
		},
	})
}

// ExampleReloadHandler handles the reload button click
func ExampleReloadHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	e, components := buildExampleMessage(i.Member.User)

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Embeds:     []*discordgo.MessageEmbed{e},
			Components: components,
		},
	})
}

// ExampleButtonHandler handles other button clicks
func ExampleButtonHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	buttonID := i.MessageComponentData().CustomID

	var response string
	switch buttonID {
	case "example_primary":
		response = "You clicked the **Primary** button!"
	case "example_secondary":
		response = "You clicked the **Secondary** button!"
	case "example_success":
		response = "You clicked the **Success** button!"
	case "example_danger":
		response = "You clicked the **Danger** button!"
	case "example_emoji":
		response = "You clicked the **Emoji** button! 🎉"
	default:
		response = "Button clicked!"
	}

	e := embed.New().
		Description(response).
		Color(embed.ColorFuchsia).
		Build()

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{e},
			Flags:  discordgo.MessageFlagsEphemeral,
		},
	})
}
