package pkg

import "strings"

func Red(args ...string) string {
	return "\033[91m" + strings.Join(args, " ") + "\033[0m"
}

func Yellow(args ...string) string {
	return "\033[93m" + strings.Join(args, " ") + "\033[0m"
}

func Purple(args ...string) string {
	return "\033[95m" + strings.Join(args, " ") + "\033[0m"
}

func Blue(args ...string) string {
	return "\033[94m" + strings.Join(args, " ") + "\033[0m"
}
