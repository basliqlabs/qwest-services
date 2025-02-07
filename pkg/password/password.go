package password

import "fmt"

const MinLength = 8
const MaxLength = 32

var Regex = fmt.Sprintf(`^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{%d,%d}$`, MinLength, MaxLength)
