package breaketernitygo

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func signPrefix(sign float64) string {
	if sign == -1 {
		return "-"
	}
	return ""
}

func (d *Decimal) ToFloat64() float64 {
	if d.mag == math.Inf(1) && d.layer == math.Inf(1) && d.sign == 1 {
		return math.Inf(1)
	}
	if d.mag == math.Inf(1) && d.layer == math.Inf(1) && d.sign == 0 {
		return math.Inf(-1)
	}
	if math.IsInf(d.layer, 0) {
		return math.NaN()
	}
	if d.layer == 0 {
		return d.sign * d.mag
	} else if d.layer == 1 {
		return d.sign * math.Pow(10, d.layer)
	} else {
		if d.mag > 0 {
			if d.sign > 0 {
				return math.Inf(1)
			} else {
				return math.Inf(-1)
			}
		} else {
			return 0
		}
	}
}

func (d *Decimal) ToString() string {

	if math.IsNaN(d.layer) || math.IsNaN(d.sign) || math.IsNaN(d.mag) {
		return "NaN"
	}
	if math.IsInf(d.mag, 1) || math.IsInf(d.layer, 1) {
		if d.sign == 1 {
			return "Infinity"
		}
		return "-Infinity"
	}

	if d.layer == 0 {
		if (d.mag < 1e21 && d.mag > 1e-7) || d.mag == 0 {
			return fmt.Sprintf("%g", d.sign*d.mag)
		}
		return fmt.Sprintf("%fe%d", d.GetMantissa(), int(d.GetExponent()))
	} else if d.layer == 1 {
		return fmt.Sprintf("%fe%d", d.GetMantissa(), int(d.GetExponent()))
	} else {
		// layer 2+
		if d.layer <= MAX_ES_IN_A_ROW {
			return fmt.Sprintf("%s%s%d", signPrefix(d.sign), strings.Repeat("e", int(d.layer)), int(d.mag))
		} else {
			return fmt.Sprintf("%s(e^%d)%d", signPrefix(d.sign), int(d.layer), int(d.mag))
		}
	}
}

func (d *Decimal) ToExponential(places int) string {
	if d.layer == 0 {
		return numberToExponentialString(d.sign*d.mag, places)
	}
	return d.ToStringWithNDecimalPlaces(places)
}

func (d *Decimal) ToFixed(places int) string {
	if d.layer == 0 {
		return numberToFixedString(d.sign*d.mag, places)
	}
	return d.ToStringWithNDecimalPlaces(places)
}

func (d *Decimal) ToPrecision(places int) string {
	e := d.GetExponent()
	if e <= -7 {
		return d.ToExponential(places - 1)
	}

	if float64(places) > e {
		return d.ToFixed(places - int(e) - 1)
	}

	return d.ToExponential(places - 1)
}

func (d *Decimal) ToStringWithNDecimalPlaces(places int) string {
	m := d.GetMantissa()
	e := d.GetExponent()
	mString := strconv.FormatFloat(decimalPlaces(m, places), 'g', -1, 64)
	eString := strconv.FormatFloat(decimalPlaces(e, places), 'g', -1, 64)
	magString := strconv.FormatFloat(decimalPlaces(d.mag, places), 'g', -1, 64)
	layerString := strconv.FormatFloat(decimalPlaces(d.layer, places), 'g', -1, 64)
	if d.layer == 0 {
		if (d.mag < 1e21 && d.mag > 1e-7) || d.mag == 0 {
			return numberToFixedString(d.sign*d.mag, places)
		}
		return mString + "e" + eString
	} else if d.layer == 1 {
		return mString + "e" + eString
	} else {
		if d.layer <= MAX_ES_IN_A_ROW {
			return signPrefix(d.sign) + strings.Repeat("e", int(d.layer)) + magString
		} else {
			return signPrefix(d.sign) + "(e^" + layerString + ")" + magString
		}
	}
}

func (d *Decimal) Normalize() *Decimal {
	// Any 0 is totally 0
	if d.sign == 0 || (d.mag == 0 && d.layer == 0) || (d.mag == math.Inf(-1) && d.layer > 0 && !math.IsInf(d.layer, 0)) {
		d.sign = 0
		d.mag = 0
		d.layer = 0
		return d
	}

	// Extract sign from negative mag at layer 0
	if d.layer == 0 && d.mag < 0 {
		d.mag = -d.mag
		d.sign = -d.sign
	}

	// Handle infinities
	if d.mag == math.Inf(1) || d.layer == math.Inf(1) || d.mag == math.Inf(-1) || d.layer == math.Inf(-1) {
		d.mag = math.Inf(1)
		d.layer = math.Inf(1)
		return d
	}

	// Handle shifting from layer 0 to negative layers
	if d.layer == 0 && d.mag < FIRST_NEG_LAYER {
		d.layer += 1
		d.mag = math.Log10(d.mag)
		return d
	}

	absMag := math.Abs(d.mag)
	signMag := sign(d.mag)

	if absMag >= EXP_LIMIT {
		d.layer += 1
		d.mag = signMag * math.Log10(absMag)
		return d
	} else {
		for absMag < LAYER_DOWN && d.layer > 0 {
			d.layer -= 1
			if d.layer == 0 {
				d.mag = math.Pow(10, d.mag)
			} else {
				d.mag = signMag * math.Pow(10, absMag)
				absMag = math.Abs(d.mag)
				signMag = sign(d.mag)
			}
		}
		if d.layer == 0 {
			if d.mag < 0 {
				// Extract sign from negative mag at layer 0
				d.mag = -d.mag
				d.sign = -d.sign
			} else if d.mag == 0 {
				// Excessive rounding can give us all zeroes
				d.sign = 0
			}
		}
	}

	if math.IsNaN(d.sign) || math.IsNaN(d.layer) || math.IsNaN(d.mag) {
		d.sign = math.NaN()
		d.layer = math.NaN()
		d.mag = math.NaN()
	}

	return d
}
