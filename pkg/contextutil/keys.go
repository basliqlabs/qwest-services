package contextutil

type contextKey string

const (
	translatorKey contextKey = "translator"
	languageKey   contextKey = "language"
)
