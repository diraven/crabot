package publicroles

import (
	"github.com/bwmarrin/discordgo"
	"github.com/diraven/sugo"
	"strings"
)

var addCmd = &sugo.Command{
	Trigger:             "register",
	Description:         "Makes existing role public.",
	PermissionsRequired: discordgo.PermissionManageRoles,
	HasParams:           true,
	Execute: func(req *sugo.Request) error {
		var err error

		// Make sure query is not empty.
		if req.Query == "" {
			return sugo.NewBadCommandUsageError(req)
		}

		// Get guild.
		guild, err := req.GetGuild()
		if err != nil {
			return err
		}

		// Get all guild roles.
		roles, err := req.Sugo.Session.GuildRoles(guild.ID)
		if err != nil {
			return sugo.WrapError(req, err)
		}

		// Process request.
		var request string

		if len(req.Message.MentionRoles) > 0 {
			// If there is at least one role mention - we use that mention.
			request = req.Message.MentionRoles[0]
		} else {
			// Otherwise we just take full request.
			request = req.Query
		}

		// Make a storage for role we matched.
		var matchedRole *discordgo.Role

		// Try to match role.
		for _, role := range roles {
			if strings.ToLower(role.Name) == strings.ToLower(request) || role.ID == request {
				if matchedRole != nil {
					return sugo.NewError(req, "too many roles found, try again with a different search")
				}
				matchedRole = role
			}
		}

		// If we did not find any match:
		if matchedRole == nil {
			// Notify user about fail.
			return sugo.NewError(req, "no roles found")
		}

		// Otherwise add new role to the public roles list.
		err = storage.add(matchedRole.ID)
		if err != nil {
			return sugo.WrapError(req, err)
		}

		// Save our changes.
		err = storage.save()
		if err != nil {
			return sugo.WrapError(req, err)
		}

		// And notify user about success.
		_, err = req.Respond("", sugo.NewSuccessEmbed(req, "role `"+matchedRole.Name+"` is public now"), false)
		return err
	},
}