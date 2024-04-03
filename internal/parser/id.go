package parser

import (
	"strconv"
	"strings"
)

func parseRange(text string, ids *[]int) error {
	parts := strings.Split(text, "-")

	if len(parts) == 1 {
		id, err := strconv.Atoi(parts[0])
		if err != nil {
			return err
		}

		*ids = append(*ids, id)
	} else if len(parts) == 2 {
		start, err := strconv.Atoi(parts[0])
		if err != nil {
			return err
		}

		end, err := strconv.Atoi(parts[1])
		if err != nil {
			return err
		}

		for i := start; i <= end; i++ {
			*ids = append(*ids, i)
		}
	}

	return nil
}

func parseIds(text string) ([]int, bool) {
	ids := []int{}
	parts := strings.Split(text, ",")

	for _, part := range parts {
		if err := parseRange(part, &ids); err != nil {
			return []int{}, false
		}
	}

	return ids, true
}
