# Discord Bot Template (Go)

內部使用的 Discord Bot 模板，提供漂亮的 Embed 訊息工具和常用功能。

## Features

- **Embed Builder**: Fluent API 建立漂亮的嵌入訊息
- **Button Builder**: 按鈕元件支援
- **Color Palette**: 40+ 預設顏色
- **Message Styles**: Success, Error, Warning, Info 模板
- **Public/Private Messages**: 支援私人訊息 (ephemeral)
- **Text Formatting**: 粗體、斜體、程式碼區塊、spoiler 等
- **Discord Timestamps**: 相對時間、日期格式化

## Project Structure

```
discord-bot-template/
├── cmd/
│   └── bot/
│       └── main.go          # Entry point
├── internal/                # 內部套件（僅限本專案使用）
│   ├── bot/
│   │   └── bot.go           # Bot 核心邏輯
│   ├── commands/
│   │   ├── commands.go      # 指令註冊中心
│   │   └── ping.go          # /ping 指令
│   ├── component/
│   │   └── button.go        # Button Builder
│   ├── config/
│   │   └── config.go        # 設定管理
│   └── embed/
│       ├── builder.go       # Embed Builder (Fluent API)
│       └── colors.go        # 顏色常數
├── Dockerfile
├── docker-compose.yml
└── go.mod
```

## Quick Start

### Docker

```bash
# 設定環境變數
cp .env.example .env
# 編輯 .env 填入 DISCORD_TOKEN

# 啟動
docker compose up -d

# 查看日誌
docker compose logs -f

# 停止
docker compose down
```

### Development (Hot Reload)

```bash
docker compose --profile dev up discord-bot-dev
```

## 新增指令

1. 在 `internal/commands/` 建立新檔案：

```go
// internal/commands/hello.go
package commands

import (
    "discord-bot-template/internal/embed"
    "github.com/bwmarrin/discordgo"
)

var HelloCommand = &discordgo.ApplicationCommand{
    Name:        "hello",
    Description: "Say hello to the bot",
}

func HelloHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
    e := embed.New().
        Title("Hello!").
        Description("Hi there! Nice to meet you!").
        Color(embed.ColorSuccess).
        Build()

    s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
        Type: discordgo.InteractionResponseChannelMessageWithSource,
        Data: &discordgo.InteractionResponseData{
            Embeds: []*discordgo.MessageEmbed{e},
        },
    })
}
```

2. 在 `internal/commands/commands.go` 註冊：

```go
func AllCommands() []*Command {
    return []*Command{
        {Definition: PingCommand, Handler: PingHandler},
        {Definition: HelloCommand, Handler: HelloHandler},
    }
}
```

## Embed 使用方式

### 快速模板

```go
embed.Success("成功", "操作完成！")
embed.Error("錯誤", "發生問題")
embed.Warning("警告", "請注意")
embed.Info("資訊", "說明內容")
```

### 完整 Builder

```go
embed.New().
    Title("標題").
    Description("描述內容").
    Color(embed.ColorBlurple).
    Author(user.Username, "", user.AvatarURL("32")).
    Thumbnail("https://example.com/image.png").
    InlineField("欄位1", "值1").   // 並排欄位
    InlineField("欄位2", "值2").
    BlockField("完整寬度", "內容"). // 獨立一行
    Footer("Footer 文字", "").
    Timestamp().
    Build()
```

### 顏色

```go
// 狀態色
embed.ColorSuccess  // 綠色
embed.ColorError    // 紅色
embed.ColorWarning  // 黃色
embed.ColorInfo     // Blurple

// 品牌色
embed.ColorBlurple  // Discord 藍紫色
embed.ColorFuchsia  // 粉紅

// 其他
embed.ColorAqua, embed.ColorPurple, embed.ColorGold
embed.ColorOrange, embed.ColorBlue, embed.ColorTeal
// ... 40+ 顏色
```

### 文字格式化

```go
embed.Bold("粗體")           // **粗體**
embed.Italic("斜體")         // *斜體*
embed.InlineCode("程式碼")   // `程式碼`
embed.CodeBlock("go", code)  // ```go ... ```
embed.Spoiler("劇透")        // ||劇透||
embed.Mention("user_id")     // <@user_id>
embed.MentionChannel("id")   // <#channel_id>
embed.RelativeTime(t)        // "2 小時前"
```

## 公開 vs 私人訊息

```go
// 公開訊息
s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
    Type: discordgo.InteractionResponseChannelMessageWithSource,
    Data: &discordgo.InteractionResponseData{
        Embeds: []*discordgo.MessageEmbed{myEmbed},
    },
})

// 私人訊息（僅用戶可見）
s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
    Type: discordgo.InteractionResponseChannelMessageWithSource,
    Data: &discordgo.InteractionResponseData{
        Embeds: []*discordgo.MessageEmbed{myEmbed},
        Flags:  discordgo.MessageFlagsEphemeral,  // 關鍵！
    },
})
```

## 按鈕使用方式

### 發送帶按鈕的訊息

```go
import "discord-bot-template/internal/component"

// 建立按鈕
row := component.NewActionRow().
    AddButton(component.PrimaryButton("btn_confirm", "確認")).
    AddButton(component.DangerButton("btn_cancel", "取消")).
    Build()

// 發送訊息
s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
    Type: discordgo.InteractionResponseChannelMessageWithSource,
    Data: &discordgo.InteractionResponseData{
        Content:    "請選擇：",
        Components: []discordgo.MessageComponent{row},
    },
})
```

### 處理按鈕點擊

在 `internal/commands/commands.go` 註冊 handler：

```go
var componentHandlers = map[string]Handler{
    "btn_confirm": ConfirmHandler,
    "btn_cancel":  CancelHandler,
}

func ConfirmHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
    s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
        Type: discordgo.InteractionResponseUpdateMessage,
        Data: &discordgo.InteractionResponseData{
            Content:    "已確認！",
            Components: []discordgo.MessageComponent{}, // 移除按鈕
        },
    })
}
```

### 按鈕樣式

```go
component.PrimaryButton("id", "藍色")    // 主要
component.SecondaryButton("id", "灰色") // 次要
component.SuccessButton("id", "綠色")   // 成功
component.DangerButton("id", "紅色")    // 危險
component.LinkButton("https://...", "連結") // 外部連結
```

### 進階 Builder

```go
component.NewButton().
    CustomID("my_button").
    Label("按我").
    Primary().
    Emoji("👍").
    Build()
```

## 環境變數

| 變數 | 必填 | 說明 |
|------|------|------|
| `DISCORD_TOKEN` | Yes | Discord Bot Token |
| `GUILD_ID` | No | 測試用伺服器 ID（指令即時更新） |
