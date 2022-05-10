package commands

import (
	"fmt"
	"strings"
)

func cmdList(c Command) *string {
	var reply string
	userFilter, ok := fetchUserFilter(c.userId)
	sources, ok2 := fetchUserSources(c.userId)
	if !ok || !ok2 {
		reply = "⚠ Oops. Internal Error. Please try again later."
		return &reply
	}

	var replyLines []string
	if userFilter != "" {
		replyLines = append(
			replyLines,
			fmt.Sprintf("🔍 RegExp Filter: '%s'", userFilter),
		)
	}
	if len(sources) == 0 {
		replyLines = append(replyLines, "🗒 No sources. Try to add!")
	} else {
		replyLines = append(replyLines, "📝 Your sources:")
		for i, source := range sources {
			replyLines = append(
				replyLines,
				fmt.Sprintf("%c %d. %s",
					boolToEmoji(source.IsActive), i+1, source.Name),
			)
		}
	}
	reply = strings.Join(replyLines, "\n")
	return &reply
}
