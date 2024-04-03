package parser

import "strings"

func scopeFromStr(str string) (Scope, bool) {
	switch Scope(str) {
	case Priority:
		return Scope(str), true
	default:
		return "", false
	}
}

func parseScope(arg string) (Scope, string) {
	parts := strings.Split(arg, ":")

	// Scoped arguments are only valid if they have exactly 2 parts. This
	// includes the case where the second part is empty.
	if len(parts) != 2 {
		return "", ""
	}

	// Try to parse the scope
	if scope, ok := scopeFromStr(parts[0]); ok {
		return scope, strings.TrimSpace(parts[1])
	}

	return "", ""
}
