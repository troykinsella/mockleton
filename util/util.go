package util

import "strings"

func KeyValuesToMap(envStr []string) map[string]string {
	env := make(map[string]string)
	for _, i := range envStr {
		key, val := SplitKeyValue(i)
		env[key] = val
	}
	return env
}

func SplitKeyValue(keyValue string) (string, string) {
	parts := strings.Split(keyValue, "=")
	key := parts[0]
	val := strings.Join(parts[1:], "=")
	return key, val
}
