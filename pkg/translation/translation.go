package translation

var globalTranslator *Translator

type Config struct {
	Default string `koanf:"default"`
	Core    string `koanf:"core"`
}

type Translator struct {
	translations map[string]map[string]string
	defaultLang  string
	coreLang     string
	currentLang  string
}

func Init(cfg Config) {
	t := &Translator{
		translations: make(map[string]map[string]string),
		defaultLang:  cfg.Default,
		coreLang:     cfg.Core,
		currentLang:  cfg.Default,
	}

	t.loadTranslations()
	globalTranslator = t
}

// T returns a translation for a key in the current language
func T(lang string, key string, data map[string]any) string {
	if globalTranslator == nil {
		panic("translator not initialized")
	}
	return globalTranslator.T(lang, key, data)
}

// SetLanguage sets the current language
func SetLanguage(lang string) {
	if globalTranslator == nil {
		panic("translator not initialized")
	}
	if lang == "" {
		lang = globalTranslator.defaultLang
	}
	globalTranslator.currentLang = lang
}

// GetCurrentLang returns the current language
func GetCurrentLang() string {
	if globalTranslator == nil {
		panic("translator not initialized")
	}
	return globalTranslator.currentLang
}

// GetDefaultLang returns the default language
func GetDefaultLang() string {
	if globalTranslator == nil {
		panic("translator not initialized")
	}
	return globalTranslator.defaultLang
}

// GetCoreLang returns the core language
func GetCoreLang() string {
	if globalTranslator == nil {
		panic("translator not initialized")
	}
	return globalTranslator.coreLang
}
