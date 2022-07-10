package utils

import (
	"github.com/bwmarrin/discordgo"
	"os"
)

func HasPermissionsAdmin(member *discordgo.Member) bool {
	var AdminRole = os.Getenv("ADMIN_ROLE")

	return ExistsInArray(member.Roles, AdminRole) || member.User.ID == os.Getenv("DEV_ID")
}
