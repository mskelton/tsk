package utils_test

import (
	"testing"
	"time"

	"github.com/mskelton/tsk/internal/utils"
	"github.com/stretchr/testify/assert"
)

const (
	day  = 24 * time.Hour
	week = 7 * day
)

func TestShortDuration(t *testing.T) {
	// Now
	assert.Equal(t, utils.ShortDuration(time.Now()), "-")
	assert.Equal(t, utils.ShortDuration(time.Now().Add(-1)), "-")

	// Seconds
	assert.Equal(t, utils.ShortDuration(time.Now().Add(-2*time.Second)), "2s")
	assert.Equal(t,
		utils.ShortDuration(time.Now().Add(-59*time.Second)),
		"59s",
	)

	// Minutes
	assert.Equal(t, utils.ShortDuration(time.Now().Add(-1*time.Minute)), "1m")
	assert.Equal(t, utils.ShortDuration(time.Now().Add(-3*time.Minute)), "3m")

	// Hours
	assert.Equal(t, utils.ShortDuration(time.Now().Add(-1*time.Hour)), "1h")
	assert.Equal(t, utils.ShortDuration(time.Now().Add(-3*time.Hour)), "3h")

	// Days
	assert.Equal(t, utils.ShortDuration(time.Now().Add(-1*day)), "1d")
	assert.Equal(t, utils.ShortDuration(time.Now().Add(-6*day)), "6d")

	// Weeks
	assert.Equal(t, utils.ShortDuration(time.Now().Add(-1*week)), "1w")
	assert.Equal(t, utils.ShortDuration(time.Now().Add(-3*week)), "3w")

	// Months
	assert.Equal(t, utils.ShortDuration(time.Now().Add(-30*day)), "1mo")
	assert.Equal(t, utils.ShortDuration(time.Now().Add(-60*day)), "2mo")
	assert.Equal(t, utils.ShortDuration(time.Now().Add(-90*day)), "3mo")

	// Years
	assert.Equal(t, utils.ShortDuration(time.Now().Add(-365*day)), "1y")
	assert.Equal(t,
		utils.ShortDuration(time.Now().Add(-400*day)),
		"1.1y",
	)
	assert.Equal(t, utils.ShortDuration(time.Now().Add(-730*day)), "2y")
	assert.Equal(t,
		utils.ShortDuration(time.Now().Add(-830*day)),
		"2.3y",
	)
	assert.Equal(t, utils.ShortDuration(time.Now().Add(-1095*day)), "3y")
}
