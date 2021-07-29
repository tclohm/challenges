import (
    "unicode"
    "unicode/utf8"
)

func isDigit(symbol string) bool {
    r, _ := utf8.DecodeRuneInString(symbol)
    return unicode.IsDigit(r)
}
