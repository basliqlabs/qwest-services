package contextutil

import (
	"context"
	"github.com/basliqlabs/qwest-services-auth/pkg/translation"
)

func WithLanguage(ctx context.Context, lang string) context.Context {
	if lang == "" {
		if t, ok := ctx.Value(translatorKey).(*translation.Translator); ok {
			lang = t.GetDefaultLang()
		}
	}
	return context.WithValue(ctx, languageKey, lang)
}

func GetLanguage(ctx context.Context) string {
	lang := ""

	if lang, ok := ctx.Value(languageKey).(string); ok {
		return lang
	}

	if t, ok := ctx.Value(translatorKey).(*translation.Translator); ok {
		lang = t.GetDefaultLang()
	}

	return lang
}

func GetCoreLang(ctx context.Context) string {
	lang := ""

	if t, ok := ctx.Value(translatorKey).(*translation.Translator); ok {
		lang = t.GetCoreLang()
	}

	if lang == "" {
		if t, ok := ctx.Value(translatorKey).(*translation.Translator); ok {
			lang = t.GetDefaultLang()
		}
	}

	return lang
}
