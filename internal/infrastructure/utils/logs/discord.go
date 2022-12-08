package logs

import (
	"fmt"
	"github.com/gtuk/discordwebhook"
	"go-gin-template/internal/infrastructure/utils/vars"
)

func SendDiscord(_type, message string) {
	var msg discordwebhook.Message
	var (
		title      string
		colort     string
		footerText = vars.AppName
		footerIcon = vars.AppIconNotTransparent
	)
	switch _type {
	case vars.TypeLogs.Error:
		title = "API Core - ERROR"
		colort = "14296837"
		msg = discordwebhook.Message{
			Embeds: &[]discordwebhook.Embed{
				{
					Title:       &title,
					Description: &message,
					Color:       &colort,
					Footer: &discordwebhook.Footer{
						Text:    &footerText,
						IconUrl: &footerIcon,
					},
				},
			},
		}
	case "log":
		title = "API Core - LOG"
		colort = "304603"
		msg = discordwebhook.Message{
			Embeds: &[]discordwebhook.Embed{
				{
					Title:       &title,
					Description: &message,
					Color:       &colort,
					Footer: &discordwebhook.Footer{
						Text:    &footerText,
						IconUrl: &footerIcon,
					},
				},
			},
		}
	case vars.TypeLogs.Warning:
		title = "API Core - ALERTA"
		colort = "14296837"
		msg = discordwebhook.Message{
			Embeds: &[]discordwebhook.Embed{
				{
					Title:       &title,
					Description: &message,
					Color:       &colort,
					Footer: &discordwebhook.Footer{
						Text:    &footerText,
						IconUrl: &footerIcon,
					},
				},
			},
		}
	case "text":
		msg = discordwebhook.Message{
			Content: &message,
		}
	}

	if err := discordwebhook.SendMessage(vars.DiscordHook, msg); err != nil {
		AddLogFile(fmt.Sprintf("2^[DISCORD LOG]0^ %s", "Connection could not be established"), vars.TypeLogs.Error)
		return
	}
}
