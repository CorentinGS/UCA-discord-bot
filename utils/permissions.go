package utils

import (
	"os"

	"github.com/bwmarrin/discordgo"
)

func HasPermissionsAdmin(member *discordgo.Member) bool {
	AdminRole := os.Getenv("ADMIN_ROLE")

	return ExistsInArray(member.Roles, AdminRole) || member.User.ID == os.Getenv("DEV_ID")
}
