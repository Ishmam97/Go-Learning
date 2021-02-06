package main

import (
	"fmt"
	"unicode"
)

type passStrength struct {
	isWeak          bool
	isStrong        bool
	Length          int
	SpecialCharsMap map[string]bool
}

func main() {
	SpecialCharsMap := map[string]bool{
		"!": false,
		"@": false,
		"#": false,
		"$": false,
		"%": false,
		"^": false,
		"&": false,
		"*": false,
		"(": false,
		")": false,
	}
	pS := passStrength{true, true, 0, SpecialCharsMap}

	pwd := "test!@pass123"
	checkpwd(pwd, &pS)
	fmt.Println(pS)
}
func checkpwd(pwd string, pS *passStrength) {
	if len(pwd) < 12 {
		fmt.Println("password length less than 12", len(pwd))
		(*pS).Length = len(pwd)
	}
	var number bool
	var upper bool
	var lower bool

	for _, c := range pwd {
		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsUpper(c):
			upper = true
		case unicode.IsLower(c):
			lower = true
		case c == '!':
			(*pS).SpecialCharsMap["!"] = true
		case c == '@':
			(*pS).SpecialCharsMap["@"] = true
		case c == '#':
			(*pS).SpecialCharsMap["#"] = true
		case c == '$':
			(*pS).SpecialCharsMap["$"] = true
		case c == '%':
			(*pS).SpecialCharsMap["%"] = true
		case c == '^':
			(*pS).SpecialCharsMap["^"] = true
		case c == '&':
			(*pS).SpecialCharsMap["&"] = true
		case c == '*':
			(*pS).SpecialCharsMap["*"] = true
		case c == '(':
			(*pS).SpecialCharsMap["("] = true
		case c == ')':
			(*pS).SpecialCharsMap[")"] = true
		default:
			(*pS).isWeak = true
			(*pS).isStrong = false
		}
		if ((*pS).isWeak == false) || number == true || upper == true || lower == true {
			(*pS).isStrong = true
			(*pS).isWeak = false
		} else {
			(*pS).isStrong = false
			(*pS).isWeak = true
		}
	}

}
