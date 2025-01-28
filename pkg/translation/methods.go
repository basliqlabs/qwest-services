package translation

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

var translationDirectory = filepath.Join("pkg", "translation")

func (t *Translator) loadTranslations() {
	dirs, err := os.ReadDir(translationDirectory)
	if err != nil {
		panic(fmt.Sprintf("failed to read translation directory: %v", err))
	}

	fileKeys := make(map[string]map[string]bool)

	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}

		lang := dir.Name()
		t.translations[lang] = make(map[string]string)

		files, err := os.ReadDir(filepath.Join(translationDirectory, lang))
		if err != nil {
			panic(fmt.Sprintf("failed to read language directory %s: %v", lang, err))
		}

		for _, file := range files {
			if filepath.Ext(file.Name()) != ".json" {
				continue
			}

			content, err := os.ReadFile(filepath.Join(translationDirectory, lang, file.Name()))
			if err != nil {
				panic(fmt.Sprintf("failed to read translation file %s/%s: %v", lang, file.Name(), err))
			}

			var fileTranslations map[string]string
			if err := json.Unmarshal(content, &fileTranslations); err != nil {
				panic(fmt.Sprintf("failed to parse translation file %s/%s: %v", lang, file.Name(), err))
			}

			if _, exists := fileKeys[file.Name()]; !exists {
				fileKeys[file.Name()] = make(map[string]bool)
			}

			for key, value := range fileTranslations {
				t.translations[lang][key] = value
				fileKeys[file.Name()][key] = true
			}
		}
	}

	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}

		lang := dir.Name()
		files, _ := os.ReadDir(filepath.Join(translationDirectory, lang))

		for _, file := range files {
			if filepath.Ext(file.Name()) != ".json" {
				continue
			}

			content, _ := os.ReadFile(filepath.Join(translationDirectory, lang, file.Name()))
			var fileTranslations map[string]string
			json.Unmarshal(content, &fileTranslations)

			expectedKeys := fileKeys[file.Name()]

			for key := range expectedKeys {
				if _, exists := fileTranslations[key]; !exists {
					panic(Error{
						Lang:    lang,
						File:    file.Name(),
						Message: fmt.Sprintf("missing key %q that exists in other languages", key),
					})
				}
			}

			for key := range fileTranslations {
				if !expectedKeys[key] {
					panic(Error{
						Lang:    lang,
						File:    file.Name(),
						Message: fmt.Sprintf("extra key %q that doesn't exist in other languages", key),
					})
				}
			}
		}
	}
}

func (t *Translator) getMessage(lang, key string) (string, bool) {
	if messages, ok := t.translations[lang]; ok {
		if msg, ok := messages[key]; ok {
			return msg, true
		}
	}

	return "", false
}

func (t *Translator) renderTemplate(tmpl string, data map[string]any) string {
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

func (t *Translator) T(lang, key string, data map[string]any) string {
	if msg, ok := t.getMessage(lang, key); ok {
		return t.renderTemplate(msg, data)
	}

	if msg, ok := t.getMessage(t.defaultLang, key); ok {
		return t.renderTemplate(msg, data)
	}

	return key
}

func (t *Translator) GetDefaultLang() string {
	return t.defaultLang
}

func (t *Translator) GetCoreLang() string {
	return t.coreLang
}
