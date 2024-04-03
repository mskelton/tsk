package parser

import (
	"strings"
)

func parseTag(text string) (Operator, string) {

	if len(text) < 2 || text[1] == ' ' {
		return "", ""
	}

	var operator Operator
	switch text[0] {
	case '+':
		operator = Include
	case '-':
		operator = Exclude
	default:
		return "", ""
	}

	return operator, strings.TrimSpace(text[1:])
}
