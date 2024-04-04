package utils

import (
	"math/rand"
	"time"

	"github.com/mskelton/tsk/internal/arg_parser"
)

func Pluralize(count int, singular string, plural string) string {
	if count == 1 {
		return singular
	} else {
		return plural
	}
}

func IsBulk(context arg_parser.ParseContext, count int) bool {
	size := 4

	for _, config := range context.Config {
		if bulk, ok := config.(arg_parser.BulkConfig); ok {
			size = bulk.Size
		}
	}

	return count >= size
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateId() string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, 8)

	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(b)
}
