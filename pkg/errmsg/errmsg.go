package errmsg

import "github.com/basliqlabs/qwest-services-auth/translation"

// TODO: migrate messages to translation json files
const (
	NotFound            = "record not found"
	CantScanQueryResult = "can't scan query result"
	SomethingWentWrong  = "something went wrong"
)

type ErrorMessage struct {
	translate translation.Translator
}

func New(t translation.Translator) ErrorMessage {
	return ErrorMessage{translate: t}
}

func (e ErrorMessage) InvalidInput(lang string) string {
	return e.translate.T(lang, "invalid_input", nil)
}
