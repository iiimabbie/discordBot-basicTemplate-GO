package commands

import (
	"fmt"
	"strings"
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
	RegisterComponent("example_open_modal", ExampleOpenModalHandler)

	// 自動註冊 Select Menu
	RegisterComponent("example_select", ExampleSelectHandler)
	RegisterComponent("example_user_select", ExampleUserSelectHandler)

	// 自動註冊 Modal
	RegisterModal("example_modal", ExampleModalSubmitHandler)
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
		Description("This is a **complete example** showing all embed and component features!\n\n" +
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

	// Row 1: Buttons
	buttonRow := component.NewActionRow().
		AddButton(component.PrimaryButton("example_primary", "Primary")).
		AddButton(component.SecondaryButton("example_secondary", "Secondary")).
		AddButton(component.SuccessButton("example_success", "Success")).
		AddButton(component.DangerButton("example_danger", "Danger")).
		Build()

	// Row 2: More buttons
	buttonRow2 := component.NewActionRow().
		AddButton(component.NewButton().
			CustomID("example_emoji").
			Label("With Emoji").
			Primary().
			Emoji("🎉").
			Build()).
		AddButton(component.NewButton().
			CustomID("example_open_modal").
			Label("Open Form").
			Secondary().
			Emoji("📝").
			Build()).
		AddButton(component.LinkButton("https://discord.com", "Discord Link")).
		Build()

	// Row 3: String Select Menu
	stringSelect := component.NewSelect().
		CustomID("example_select").
		Placeholder("Choose your favorite color...").
		AddOptionWithEmoji("Red", "red", "A warm, passionate color", "🔴").
		AddOptionWithEmoji("Green", "green", "The color of nature", "🟢").
		AddOptionWithEmoji("Blue", "blue", "A calm, peaceful color", "🔵").
		AddOptionWithEmoji("Purple", "purple", "A royal, creative color", "🟣").
		Build()
	selectRow := component.SelectRow(stringSelect)

	// Row 4: User Select Menu
	userSelect := component.NewUserSelect("example_user_select").
		Placeholder("Select a user...").
		Build()
	userSelectRow := component.SelectRow(userSelect)

	// Row 5: Reload button
	reloadRow := component.ReloadButtonRow("example_reload")

	return e, []discordgo.MessageComponent{buttonRow, buttonRow2, selectRow, userSelectRow, reloadRow}
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

// ExampleButtonHandler handles button clicks
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

// ExampleSelectHandler handles the string select menu
func ExampleSelectHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.MessageComponentData()
	selected := data.Values[0]

	colorMap := map[string]int{
		"red":    embed.ColorRed,
		"green":  embed.ColorGreen,
		"blue":   embed.ColorBlue,
		"purple": embed.ColorPurple,
	}

	colorName := strings.ToUpper(selected[:1]) + selected[1:]
	color := colorMap[selected]

	e := embed.New().
		Description(fmt.Sprintf("You selected **%s**! 🎨", colorName)).
		Color(color).
		Build()

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{e},
			Flags:  discordgo.MessageFlagsEphemeral,
		},
	})
}

// ExampleUserSelectHandler handles the user select menu
func ExampleUserSelectHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.MessageComponentData()
	userID := data.Values[0]

	e := embed.New().
		Description(fmt.Sprintf("You selected %s! 👤", embed.Mention(userID))).
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

// ExampleOpenModalHandler opens the example modal
func ExampleOpenModalHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	titleInput := component.NewTextInput().
		CustomID("example_modal_title").
		Label("Title").
		Placeholder("Enter a title...").
		Short().
		Required().
		MaxLength(100).
		Build()

	messageInput := component.NewTextInput().
		CustomID("example_modal_message").
		Label("Message").
		Placeholder("Enter your message here...").
		Paragraph().
		Required().
		MinLength(10).
		MaxLength(500).
		Build()

	modal := component.NewModal().
		CustomID("example_modal").
		Title("📝 Example Form").
		AddTextInput(titleInput).
		AddTextInput(messageInput).
		Build()

	s.InteractionRespond(i.Interaction, modal)
}

// ExampleModalSubmitHandler handles the modal submission
func ExampleModalSubmitHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ModalSubmitData()

	title := component.GetModalValue(data, "example_modal_title")
	message := component.GetModalValue(data, "example_modal_message")

	e := embed.New().
		Title("✅ Form Submitted!").
		Color(embed.ColorSuccess).
		BlockField("Title", title).
		BlockField("Message", message).
		Footer(fmt.Sprintf("Submitted by %s", i.Member.User.Username), i.Member.User.AvatarURL("32")).
		Timestamp().
		Build()

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{e},
			Flags:  discordgo.MessageFlagsEphemeral,
		},
	})
}
