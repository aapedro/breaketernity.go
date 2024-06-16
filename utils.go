package breaketernity

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

func excessSlog(d *Decimal, base *Decimal, linear bool) (*Decimal, int) {
	nBase := base.ToFloat64()
	if nBase == 1 || nBase <= 0 {
		return dFC_NN(math.NaN(), math.NaN(), math.NaN()), 0
	}
	if nBase > 1.44466786100976613366 {
		return d.Slog(base, 100, linear), 0
	}
	negLnBase := base.Ln().Neg()
	lower := negLnBase.LambertW(true).Divide(negLnBase)
	upper := dInf
	if nBase > 1 {
		upper = negLnBase.LambertW(false).Divide(negLnBase)
	}
	if nBase > 1.444667861009766133 {
		lower = decimalFromFloat64(math.E)
		upper = decimalFromFloat64(math.E)
	}
	if d.Lt(lower) {
		return d.Slog(base, 100, linear), 0
	}
	if d.Eq(lower) {
		return dFC_NN(1, math.Inf(1), math.Inf(1)), 0
	}
	if d.Eq(upper) {
		return dFC_NN(1, math.Inf(-1), math.Inf(-1)), 0
	}
	if d.Gt(upper) {
		slogZero := upper.Multiply(D(2))
		slogOne := base.Pow(slogZero)
		estimate := 0.
		if d.Gte(slogZero) && d.Lt(slogOne) {
			estimate = 0.
		} else if d.Gte(slogOne) {
			payload := slogOne
			estimate = 1.
			for payload.Lt(d) {
				payload = base.Pow(payload)
				estimate += 1
				if payload.layer > 3 {
					layersLeft := math.Floor(d.layer - payload.layer)
					payload = base.IteratedExp(layersLeft, payload, linear)
				}
			}
			if payload.Gt(d) {
				newPayload := payload.Log(base)
				payload.sign = newPayload.sign
				payload.layer = newPayload.layer
				payload.mag = newPayload.mag
				estimate -= 1
			}
		} else if d.Lt(slogZero) {
			payload := slogZero
			estimate = 0.
			for payload.Gt(d) {
				payload = payload.Log(base)
				estimate -= 1
			}
		}
		fracHeight := 0.
		tested := 0.
		stepSize := 0.5
		towerTop := slogZero
		guess := dZero

		for stepSize > 1e-16 {
			tested = fracHeight + stepSize
			towerTop = slogZero.Pow(D(1 - tested)).Multiply(slogOne.Pow(D(tested)))
			guess = IteratedExp(base, estimate, towerTop, false)
			if guess.Eq(d) {
				fracHeight *= stepSize
			}
			stepSize /= 2
		}
		if guess.NeqTolerance(d, 1e-7) {
			return dFC_NN(math.NaN(), math.NaN(), math.NaN()), 0
		}
		return decimalFromFloat64(estimate + fracHeight), 2
	}
	if d.Lt(upper) && d.Gt(lower) {
		slogZero := lower.Multiply(upper).Sqrt()
		slogOne := base.Pow(slogZero)
		estimate := 0.
		if d.Lte(slogZero) && d.Gt(slogOne) {
			estimate = 0.
		} else if d.Lte(slogOne) {
			payload := slogOne
			estimate = 1
			for payload.Gt(d) {
				payload = base.Pow(payload)
				estimate += 1
			}
			if payload.Lt(d) {
				newPayload := payload.Log(base)
				payload.sign = newPayload.sign
				payload.layer = newPayload.layer
				payload.mag = newPayload.mag
				estimate -= 1
			}
		} else if d.Gt(slogZero) {
			payload := slogZero
			estimate = 0.
			for payload.Lt(d) {
				payload = payload.Log(base)
				estimate -= 1
			}
		}

		fracHeight := 0.
		tested := 0.
		stepSize := 0.5
		towerTop := slogZero
		guess := dZero
		for stepSize > 1e-16 {
			tested = fracHeight + stepSize
			towerTop = slogZero.Pow(D(1 - tested)).Multiply(slogOne.Pow(D(tested)))
			guess = IteratedExp(base, estimate, towerTop, false)
			if guess.Eq(d) {
				return decimalFromFloat64(estimate + tested), 1
			}
			stepSize /= 2
		}
		if guess.NeqTolerance(d, 1e-7) {
			return dFC_NN(math.NaN(), math.NaN(), math.NaN()), 0
		}
		return decimalFromFloat64(estimate + fracHeight), 1
	}

	panic("Unhandled behavior in excessSlog")
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
