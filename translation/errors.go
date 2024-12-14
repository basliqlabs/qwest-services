package translation

import (
	"fmt"
)

type Error struct {
	Lang    string
	File    string
	Message string
}

func (e Error) Error() string {
	return fmt.Sprintf("translation error in %s/%s: %s", e.Lang, e.File, e.Message)
}
