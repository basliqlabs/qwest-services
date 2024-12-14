package translation

import (
	"os"
	"testing"
)

func setupTestFiles() func() {
	if err := os.MkdirAll("translation/en", 0755); err != nil {
		panic(err)
	}
	if err := os.MkdirAll("translation/es", 0755); err != nil {
		panic(err)
	}

	enContent := `{
		"greeting": "Hello, {{.Name}}!",
		"farewell": "Goodbye!",
		"welcome": "Welcome to our app"
	}`
	if err := os.WriteFile("translation/en/messages.json", []byte(enContent), 0644); err != nil {
		panic(err)
	}

	esContent := `{
		"greeting": "¡Hola, {{.Name}}!",
		"farewell": "¡Adiós!",
		"welcome": "Bienvenido a nuestra aplicación"
	}`
	if err := os.WriteFile("translation/es/messages.json", []byte(esContent), 0644); err != nil {
		panic(err)
	}

	return func() {
		os.RemoveAll("translation")
	}
}

func TestTranslator(t *testing.T) {
	cleanup := setupTestFiles()
	defer cleanup()

	for _, tc := range translationTestCases {
		t.Run(tc.description, func(t *testing.T) {
			translator := New(tc.cfg)
			if actual := translator.T(tc.lang, tc.key, tc.data); actual != tc.expected {
				t.Errorf("\nExpected: %v\nActual: %v\n", tc.expected, actual)
			}
		})
	}
}

func TestTranslatorValidation(t *testing.T) {
	if err := os.MkdirAll("translation/en", 0755); err != nil {
		panic(err)
	}
	if err := os.MkdirAll("translation/es", 0755); err != nil {
		panic(err)
	}
	defer os.RemoveAll("translation")

	enContent := `{
		"greeting": "Hello!",
		"extra_key": "Extra"
	}`
	if err := os.WriteFile("translation/en/messages.json", []byte(enContent), 0644); err != nil {
		panic(err)
	}

	esContent := `{
		"greeting": "¡Hola!"
	}`
	if err := os.WriteFile("translation/es/messages.json", []byte(esContent), 0644); err != nil {
		panic(err)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for mismatched keys, but got none")
		}
	}()

	New(Config{Default: "en"})
}
