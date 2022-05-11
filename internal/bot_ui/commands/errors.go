package commands

const (
	ErrHelp          = "👇 Use commands from Menu"
	errInternalError = "⚠ Oops. Internal Error. Please try again later."

	errNoAnySources     = "🤷 You didn't add any Sources"
	errNoSourceIndex    = "👉 Hey! You forgot index of Source"
	errWrongSourceIndex = "👉 Hey! Index of Source should be a number"
	errNoSuchSource     = "👉 Hey! There is no Source with such index"

	errAddNoUrl         = "👉 Hey! You forgot source URL"
	errAddUrlTooLong    = "🤯 Oh! Your URL is too looong"
	errAddRssParseError = "🤒 I had troubles parsing RSS from that URL, sorry"

	errFilterTooLong = "🤯 Oh! Your Filter is too looong"
	errFilterRegExp  = "🤒 I had troubles compiling that RegExp Filter, sorry"

	errRenameNoName = "👉 Hey! You forgot to provide new name"
)
