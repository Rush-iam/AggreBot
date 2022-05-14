package messages

import (
	"AggreBot/internal/bot_ui/command"
	"AggreBot/internal/bot_ui/errors"
	"AggreBot/internal/bot_ui/markup"
	"AggreBot/internal/pkg/api"
	"fmt"
	"strings"
)

func cmdListReply(userFilter string, sources []*api.Source) string {
	replyLines := make([]string, 0, len(sources)+2)

	if userFilter != "" {
		replyLines = append(replyLines, cmdFilterReply(userFilter))
	}

	if len(sources) == 0 {
		replyLines = append(replyLines, "🗒 No sources. Try to add!")
	} else {
		replyLines = append(replyLines, "📝 Your sources:")
		for i, source := range sources {
			replyLines = append(
				replyLines,
				fmt.Sprintf("%c %d. %s",
					markup.BoolToEmoji(source.IsActive), i+1, source.Name),
			)
		}
	}

	return strings.Join(replyLines, "\n")
}

func (m *Manager) cmdList(c *command.Command) string {
	userFilter, err1 := m.backend.GetUserFilter(c.UserId)
	sources, err2 := m.backend.GetUserSources(c.UserId)
	if err1 != nil || err2 != nil {
		return errors.ErrInternalError
	}

	return cmdListReply(*userFilter, sources)
}
