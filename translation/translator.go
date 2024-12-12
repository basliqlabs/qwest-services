package translation

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

type Config struct {
	Default string `koanf:"default"`
}

type Translator struct {
	translations map[string]map[string]string // lang -> key -> message
	defaultLang  string
}

func New(cfg Config) *Translator {
	t := &Translator{
		translations: make(map[string]map[string]string),
		defaultLang:  cfg.Default,
	}

	t.loadTranslations()

	return t
}

func (t *Translator) loadTranslations() {
	dirs, err := os.ReadDir("translation")
	if err != nil {
		panic(fmt.Sprintf("failed to read translation directory: %v", err))
	}

	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}

		lang := dir.Name()
		t.translations[lang] = make(map[string]string)

		files, err := os.ReadDir(filepath.Join("translation", lang))
		if err != nil {
			panic(fmt.Sprintf("failed to read language directory %s: %v", lang, err))
		}

		for _, file := range files {
			if filepath.Ext(file.Name()) != ".json" {
				continue
			}

			content, err := os.ReadFile(filepath.Join("translation", lang, file.Name()))
			if err != nil {
				panic(fmt.Sprintf("failed to read translation file %s/%s: %v", lang, file.Name(), err))
			}

			var fileTranslations map[string]string
			if err := json.Unmarshal(content, &fileTranslations); err != nil {
				panic(fmt.Sprintf("failed to parse translation file %s/%s: %v", lang, file.Name(), err))
			}

			for key, value := range fileTranslations {
				t.translations[lang][key] = value
			}
		}
	}

	return
}

func (t *Translator) T(lang, key string, data map[string]interface{}) string {
	if msg, ok := t.getMessage(lang, key); ok {
		return t.renderTemplate(msg, data)
	}

	if msg, ok := t.getMessage(t.defaultLang, key); ok {
		return t.renderTemplate(msg, data)
	}

	return key
}

func (t *Translator) getMessage(lang, key string) (string, bool) {
	if messages, ok := t.translations[lang]; ok {
		if msg, ok := messages[key]; ok {
			return msg, true
		}
	}

	return "", false
}

func (t *Translator) renderTemplate(tmpl string, data map[string]interface{}) string {
	if data == nil {
		return tmpl
	}

	textTemplate, err := template.New("translation").Parse(tmpl)
	if err != nil {
		return tmpl
	}

	var buf bytes.Buffer
	if err := textTemplate.Execute(&buf, data); err != nil {
		return tmpl
	}

	return buf.String()
}
