package contextutil

import (
	"context"
	"github.com/basliqlabs/qwest-services-auth/translation"
)

func WithTranslator(ctx context.Context, t *translation.Translator) context.Context {
	return context.WithValue(ctx, translatorKey, t)
}

func GetTranslation(ctx context.Context, key string, data map[string]any) string {
	if t, ok := ctx.Value(translatorKey).(*translation.Translator); ok {
		return t.T(GetLanguage(ctx), key, data)
	}

	return key
}

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
