package commands

import (
	"fmt"
	"github.com/m4xirq/Zertex/XDGv2/api/server"
	"github.com/m4xirq/Zertex/XDGv2/cmd/apiserver/discord/discord"
	"github.com/m4xirq/Zertex/XDGv2/cmd/apiserver/discord/utils"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func UserListCommand(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if !UserHasPermission(s,m, "iMH5RjESHkMJr8EwbcNjYtcsz9fIXEGlp0fXDc5u") {
		utils.SendError(s,m, "Sorry, this is only available for Staff members.")
		return
	}

	if program, err := server.GetProgram("iMH5RjESHkMJr8EwbcNjYtcsz9fIXEGlp0fXDc5u"); err == nil {
		emb := utils.NewEmbed().SetTitle(fmt.Sprintf("List of active users for %s", program.Name)).
			SetThumbnail(discord.BotUser.AvatarURL("250x250")).
			SetColor(utils.RandomColor()).
			SetFooter(utils.FooterTimestamp())

		var usernames []string

		for _, client := range server.APIServer.GetClients() {
			usernames = append(usernames, fmt.Sprintf("`%s`", client.Name))
		}

		emb.Description = strings.Join(usernames, ", ")
		s.ChannelMessageSendEmbed(m.ChannelID, emb.MessageEmbed)
	} else {
		fmt.Println(err.Error())
	}
}