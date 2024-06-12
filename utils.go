package breaketernitygo

import (
	"fmt"
	"math"
	"strconv"
)

func decimalPlaces(value float64, places int) float64 {
	len := float64(places) + 1
	numDigits := math.Ceil(math.Log10(math.Abs(value)))
	rounded := math.Round(value*math.Pow(10, len-numDigits)) * math.Pow(10, numDigits-len)

	result, _ := strconv.ParseFloat(fmt.Sprintf("%.*f", int(places), rounded), 64)
	return result
}

func numberToExponentialString(value float64, digits int) string {
	if math.IsNaN(value) {
		return "NaN"
	}
	if math.IsInf(value, 1) {
		return "Infinity"
	}
	if math.IsInf(value, -1) {
		return "-Infinity"
	}
	format := fmt.Sprintf("%%.%de", digits) // Creates a format string like "%.2e"
	return fmt.Sprintf(format, value)
}

func numberToFixedString(value float64, digits int) string {
	if math.IsNaN(value) {
		return "NaN"
	}
	if math.IsInf(value, 1) {
		return "Infinity"
	}
	if math.IsInf(value, -1) {
		return "-Infinity"
	}
	format := fmt.Sprintf("%%.%df", digits) // Creates a format string like "%.2f"
	return fmt.Sprintf(format, value)
}

func sign(x float64) float64 {
	if x == 0 {
		return 0
	}
	return math.Copysign(1, x)
}

func slogCritical(base float64, height float64) float64 {
	if base > 10 {
		return height - 1
	}
	return criticalSection(base, height, CRITICAL_SLOG_VALUES, false)
}

func tetrateCritical(base float64, height float64) float64 {
	return criticalSection(base, height, CRITICAL_TETR_VALUES, false)
}

func criticalSection(base float64, height float64, grid [][]float64, _ bool) float64 {
	height *= 10
	if height < 0 {
		height = 0
	}
	if height > 10 {
		height = 10
	}
	if base < 2 {
		base = 2
	}
	if base > 10 {
		base = 10
	}
	var lower, upper float64
	for i, criticalHeader := range CRITICAL_HEADERS {
		if criticalHeader == base {
			lower = grid[i][int(math.Floor(height))]
			upper = grid[i][int(math.Ceil(height))]
			break
		} else if criticalHeader < base && CRITICAL_HEADERS[i+1] > base {
			baseFrac := (base - criticalHeader) / (CRITICAL_HEADERS[i+1] - criticalHeader)
			lower = grid[i][int(math.Floor(height))]*(1-baseFrac) + grid[i+1][int(math.Floor(height))]*baseFrac
			upper = grid[i][int(math.Ceil(height))]*(1-baseFrac) + grid[i+1][int(math.Ceil(height))]*baseFrac
			break
		}
	}
	frac := height - math.Floor(height)
	if lower <= 0 || upper <= 0 {
		return lower*(1-frac) + upper*frac
	} else {
		return math.Pow(base, (math.Log(lower)/math.Log(base))*(1-frac)+(math.Log(upper)/math.Log(base))*frac)
	}
}
