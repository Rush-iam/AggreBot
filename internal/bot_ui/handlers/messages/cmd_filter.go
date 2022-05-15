package messages

import (
	"AggreBot/internal/bot_ui/command"
	"AggreBot/internal/bot_ui/errors"
	"fmt"
	"regexp"
	"strings"
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
	args := strings.Fields(c.Text)

	if len(args) > 1 {
		userFilter = args[1]
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
