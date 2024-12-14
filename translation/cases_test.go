package translation

type testCase struct {
	description string
	cfg         Config
	lang        string
	key         string
	data        map[string]interface{}
	expected    string
}

var translationTestCases = []testCase{
	{
		description: "should translate basic text correctly",
		cfg:         Config{Default: "en"},
		lang:        "es",
		key:         "farewell",
		data:        nil,
		expected:    "¡Adiós!",
	},
	{
		description: "should translate template with variables",
		cfg:         Config{Default: "en"},
		lang:        "es",
		key:         "greeting",
		data:        map[string]interface{}{"Name": "Juan"},
		expected:    "¡Hola, Juan!",
	},
	{
		description: "should fallback to default language when translation missing",
		cfg:         Config{Default: "en"},
		lang:        "fr",
		key:         "welcome",
		data:        nil,
		expected:    "Welcome to our app",
	},
	{
		description: "should return key itself when no translation found",
		cfg:         Config{Default: "en"},
		lang:        "es",
		key:         "missing_key",
		data:        nil,
		expected:    "missing_key",
	},
}
