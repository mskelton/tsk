package utils

import (
	"fmt"
	"strings"
	"time"
)

const (
	hoursInDay   = 24
	hoursInWeek  = 24 * 7
	hoursInMonth = 24 * 30
	hoursInYear  = 24 * 365
)

func ShortDuration(t time.Time) string {
	now := time.Now()
	duration := now.Sub(t)

	if duration.Hours() > hoursInYear {
		return strings.Replace(fmt.Sprintf("%.1fy", duration.Hours()/hoursInYear), ".0", "", 1)
	} else if duration.Hours() > hoursInMonth {
		return fmt.Sprintf("%dmo", int(duration.Hours()/hoursInMonth))
	} else if duration.Hours() > hoursInWeek {
		return fmt.Sprintf("%dw", int(duration.Hours()/hoursInWeek))
	} else if duration.Hours() > hoursInDay {
		return fmt.Sprintf("%dd", int(duration.Hours()/hoursInDay))
	} else if duration.Hours() > 1 {
		return fmt.Sprintf("%dh", int(duration.Hours()))
	} else if duration.Minutes() > 1 {
		return fmt.Sprintf("%dm", int(duration.Minutes()))
	} else if duration.Seconds() > 1 {
		return fmt.Sprintf("%ds", int(duration.Seconds()))
	}

	// If the duration is less than 1 second, we just return "-". This is
	// primarily to make the tests more stable.
	return "-"
}
