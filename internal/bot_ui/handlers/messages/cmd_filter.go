package messages

import (
	"AggreBot/internal/bot_ui/command"
	"AggreBot/internal/bot_ui/errors"
	"fmt"
	"regexp"
)

func cmdFilterReply(userFilter string) string {
	if userFilter == "" {
		return "🔍 No Filter set"
	} else {
		return fmt.Sprintf("🔍 RegExp Filter: '%s'", userFilter)
	}
}

func (m *Manager) cmdFilter(c *command.Command) string {
	var userFilter string
	if len(c.Args) > 0 {
		userFilter = c.Args[0]
	}

	if len([]rune(userFilter)) > 256 {
		return errors.ErrFilterTooLong
	}

	_, err := regexp.Compile(userFilter)
	if err != nil {
		return errors.ErrFilterRegExp
	}

	err = m.backend.UpdateUserFilter(c.UserId, userFilter)
	if err != nil {
		return errors.ErrInternalError
	}

	return cmdFilterReply(userFilter)
}
