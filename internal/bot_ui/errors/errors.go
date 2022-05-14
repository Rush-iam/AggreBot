package errors

const (
	ErrHelp          = "👇 Use commands from Menu"
	ErrInternalError = "⚠ Oops. Internal Error. Please try again later."

	ErrNoAnySources     = "🤷 You didn't add any Sources"
	ErrNoSourceIndex    = "👉 Hey! You forgot index of Source"
	ErrWrongSourceIndex = "👉 Hey! Index of Source should be a number"
	ErrNoSuchSource     = "👉 Hey! There is no Source with such index"

	ErrAddNoUrl         = "👉 Hey! You forgot source URL"
	ErrAddUrlTooLong    = "🤯 Oh! Your URL is too looong"
	ErrAddRssParseError = "🤒 I had troubles parsing RSS/Atom from that URL, sorry"

	ErrFilterTooLong = "🤯 Oh! Your Filter is too looong"
	ErrFilterRegExp  = "🤒 I had troubles compiling that RegExp Filter, sorry"

	ErrRenameNoName = "👉 Hey! You forgot to provide new name"
)
