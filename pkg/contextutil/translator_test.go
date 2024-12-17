package contextutil

import (
	"context"
	"os"
	"testing"

	"github.com/basliqlabs/qwest-services-auth/translation"
)

func setupTestFiles() func() {
	if err := os.MkdirAll("translation/en", 0755); err != nil {
		panic(err)
	}
	if err := os.MkdirAll("translation/fa", 0755); err != nil {
		panic(err)
	}

	enContent := `{
		"welcome": "You are very welcome."
	}`
	if err := os.WriteFile("translation/en/messages.json", []byte(enContent), 0644); err != nil {
		panic(err)
	}

	faContent := `{
		"welcome": "خوش آمدید."
	}`
	if err := os.WriteFile("translation/fa/messages.json", []byte(faContent), 0644); err != nil {
		panic(err)
	}

	return func() {
		os.RemoveAll("translation")
	}
}

func setupTranslator() (*translation.Translator, func()) {
	cleanup := setupTestFiles()
	translator := translation.New(translation.Config{
		Default: "en",
		Core:    "en",
	})
	return translator, cleanup
}

func TestWithTranslator(t *testing.T) {
	translator, cleanup := setupTranslator()
	defer cleanup()
	ctx := context.Background()

	// Test setting translator
	ctxWithTranslator := WithTranslator(ctx, translator)
	if got := ctxWithTranslator.Value(translatorKey); got != translator {
		t.Errorf("WithTranslator() = %v, want %v", got, translator)
	}

	// Test overwriting translator
	newTranslator, cleanup := setupTranslator()
	defer cleanup()
	ctxWithNewTranslator := WithTranslator(ctxWithTranslator, newTranslator)
	if got := ctxWithNewTranslator.Value(translatorKey); got != newTranslator {
		t.Errorf("WithTranslator() overwrite = %v, want %v", got, newTranslator)
	}
}

func TestGetTranslation(t *testing.T) {
	translator, cleanup := setupTranslator()
	defer cleanup()
	ctx := context.Background()
	ctxWithTranslator := WithTranslator(ctx, translator)

	tests := []struct {
		name     string
		ctx      context.Context
		key      string
		data     map[string]any
		expected string
	}{
		{
			name:     "with valid translator and key",
			ctx:      ctxWithTranslator,
			key:      "welcome",
			data:     nil,
			expected: "You are very welcome.",
		},
		{
			name:     "with missing translator",
			ctx:      context.Background(),
			key:      "welcome",
			data:     nil,
			expected: "welcome",
		},
		{
			name:     "with invalid key",
			ctx:      ctxWithTranslator,
			key:      "nonexistent_key",
			data:     nil,
			expected: "nonexistent_key",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetTranslation(tt.ctx, tt.key, tt.data)
			if got != tt.expected {
				t.Errorf("GetTranslation() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestWithLanguage(t *testing.T) {
	translator, cleanup := setupTranslator()
	defer cleanup()
	ctx := context.Background()
	ctxWithTranslator := WithTranslator(ctx, translator)

	tests := []struct {
		name     string
		ctx      context.Context
		lang     string
		expected string
	}{
		{
			name:     "set valid language",
			ctx:      ctxWithTranslator,
			lang:     "es",
			expected: "es",
		},
		{
			name:     "empty language should use default",
			ctx:      ctxWithTranslator,
			lang:     "",
			expected: "en",
		},
		{
			name:     "set language without translator",
			ctx:      context.Background(),
			lang:     "fr",
			expected: "fr",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctxWithLang := WithLanguage(tt.ctx, tt.lang)
			got := GetLanguage(ctxWithLang)
			if got != tt.expected {
				t.Errorf("GetLanguage() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestGetCoreLang(t *testing.T) {
	translator, cleanup := setupTranslator()
	defer cleanup()
	ctx := context.Background()
	ctxWithTranslator := WithTranslator(ctx, translator)

	tests := []struct {
		name     string
		ctx      context.Context
		expected string
	}{
		{
			name:     "get core language with translator",
			ctx:      ctxWithTranslator,
			expected: "en",
		},
		{
			name:     "get core language without translator",
			ctx:      context.Background(),
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetCoreLang(tt.ctx)
			if got != tt.expected {
				t.Errorf("GetCoreLang() = %v, want %v", got, tt.expected)
			}
		})
	}
}
