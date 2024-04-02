package utils

import "time"

func GenerateId() string {
	return time.Now().Format("20060102150405")
}
