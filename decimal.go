package breaketernity

import (
	"math"
	"strconv"
	"strings"
)

type Decimal struct {
	sign  float64
	layer float64
	mag   float64
}

type DecimalSource interface {
	~*Decimal | float64 | float32 | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | string
}

// D is a function that creates a new Decimal instance from a given source.
// The source can be another *Decimal, a string, or any numeric type.
// The function returns a pointer to the newly created Decimal instance.
func D[S DecimalSource](source S) *Decimal {
	return decimalFromSource(source)
}

func DFC(sign float64, layer float64, mag float64) *Decimal {
	return dFC(sign, layer, mag)
}

func decimalFromSource[DS DecimalSource](value DS) *Decimal {
	switch v := any(value).(type) {
	case *Decimal:
		return decimalFromDecimal(v)
	case float64:
		return decimalFromFloat64(v)
	case float32:
		return decimalFromFloat64(float64(v))
	case int:
		return decimalFromFloat64(float64(v))
	case int8:
		return decimalFromFloat64(float64(v))
	case int16:
		return decimalFromFloat64(float64(v))
	case int32:
		return decimalFromFloat64(float64(v))
	case int64:
		return decimalFromFloat64(float64(v))
	case uint:
		return decimalFromFloat64(float64(v))
	case uint8:
		return decimalFromFloat64(float64(v))
	case uint16:
		return decimalFromFloat64(float64(v))
	case uint32:
		return decimalFromFloat64(float64(v))
	case uint64:
		return decimalFromFloat64(float64(v))
	case string:
		return decimalFromString(v, false)
	}
	return nil
}

func decimalFromDecimal(d *Decimal) *Decimal {
	return &Decimal{sign: d.sign, layer: d.layer, mag: d.mag}
}

func decimalFromFloat64(f float64) *Decimal {
	return &Decimal{sign: sign(f), layer: 0, mag: math.Abs(f)}
}

func decimalFromString(s string, linearhyper4 bool) *Decimal {
	// TODO: cache
	if IGNORE_COMMAS {
		s = strings.Replace(s, ",", "", -1)
	} else if COMMAS_ARE_DECIMAL_POINTS {
		s = strings.Replace(s, ",", ".", -1)
	}

	pentationParts := strings.Split(s, "^^^")
	if len(pentationParts) == 2 {
		base, _ := strconv.ParseFloat(pentationParts[0], 64)
		height, _ := strconv.ParseFloat(pentationParts[1], 64)
		heightParts := strings.Split(pentationParts[1], ";")
		payload := float64(1)
		if len(heightParts) == 2 {
			payload, _ = strconv.ParseFloat(heightParts[1], 64)
			if math.IsInf(payload, 0) {
				payload = 1
			}
		}
		if !math.IsInf(base, 0) && !math.IsInf(height, 0) {
			result := Pentate(D(base), height, D(payload), linearhyper4)
			return &Decimal{sign: result.sign, layer: result.layer, mag: result.mag}
		}
	}

	tetrationParts := strings.Split(s, "^^")
	if len(tetrationParts) == 2 {
		base, _ := strconv.ParseFloat(tetrationParts[0], 64)
		height, _ := strconv.ParseFloat(tetrationParts[1], 64)
		heightParts := strings.Split(tetrationParts[1], ";")
		payload := float64(1)
		if len(heightParts) == 2 {
			payload, _ = strconv.ParseFloat(heightParts[1], 64)
			if math.IsInf(payload, 0) {
				payload = 1
			}
		}
		if !math.IsInf(base, 0) && !math.IsInf(height, 0) {
			result := Tetrate(D(base), height, D(payload), linearhyper4)
			return &Decimal{sign: result.sign, layer: result.layer, mag: result.mag}
		}
	}

	powParts := strings.Split(s, "^")
	if len(powParts) == 2 {
		base, _ := strconv.ParseFloat(powParts[0], 64)
		exponent, _ := strconv.ParseFloat(powParts[1], 64)
		if !math.IsInf(base, 0) && !math.IsInf(exponent, 0) {
			result := Pow(D(base), D(exponent))
			return &Decimal{sign: result.sign, layer: result.layer, mag: result.mag}
		}
	}

	s = strings.ToLower(strings.TrimSpace(s))

	ptParts := strings.Split(s, "pt")
	if len(ptParts) == 2 {
		base := 10.0
		negative := false
		if ptParts[0][0] == '-' {
			negative = true
			ptParts[0] = ptParts[0][1:]
		}
		height, _ := strconv.ParseFloat(ptParts[0], 64) // Error ignored
		ptParts[1] = strings.Replace(ptParts[1], "(", "", -1)
		ptParts[1] = strings.Replace(ptParts[1], ")", "", -1)
		payload, _ := strconv.ParseFloat(ptParts[1], 64) // Error ignored
		if math.IsInf(payload, 0) {
			payload = 1
		}
		if !math.IsInf(base, 0) && !math.IsInf(height, 0) {
			result := Tetrate(D(base), height, D(payload), linearhyper4)
			if negative {
				result.sign *= -1
			}
			return &Decimal{sign: result.sign, layer: result.layer, mag: result.mag}
		}
	}

	ptParts = strings.Split(s, "p")
	if len(ptParts) == 2 {
		base := 10.0
		negative := false
		if ptParts[0][0] == '-' {
			negative = true
			ptParts[0] = ptParts[0][1:]
		}
		height, _ := strconv.ParseFloat(ptParts[0], 64) // Error ignored
		ptParts[1] = strings.Replace(ptParts[1], "(", "", -1)
		ptParts[1] = strings.Replace(ptParts[1], ")", "", -1)
		payload, _ := strconv.ParseFloat(ptParts[1], 64) // Error ignored
		if math.IsInf(payload, 0) {
			payload = 1
		}
		if !math.IsInf(base, 0) && !math.IsInf(height, 0) {
			result := Tetrate(D(base), height, D(payload), linearhyper4)
			if negative {
				result.sign *= -1
			}
			return &Decimal{sign: result.sign, layer: result.layer, mag: result.mag}
		}
	}

	fParts := strings.Split(s, "f")
	if len(fParts) == 2 {
		base := 10.0
		negative := false
		if fParts[0][0] == '-' {
			negative = true
			fParts[0] = fParts[0][1:]
		}
		fParts[0] = strings.Replace(fParts[0], "(", "", -1)
		fParts[0] = strings.Replace(fParts[0], ")", "", -1)
		payload, _ := strconv.ParseFloat(fParts[0], 64) // Error ignored
		fParts[1] = strings.Replace(fParts[1], "(", "", -1)
		fParts[1] = strings.Replace(fParts[1], ")", "", -1)
		height, _ := strconv.ParseFloat(fParts[1], 64) // Error ignored
		if math.IsInf(payload, 0) {
			payload = 1
		}
		if !math.IsInf(base, 0) && !math.IsInf(height, 0) {
			result := Tetrate(D(base), height, D(payload), linearhyper4)
			if negative {
				result.sign *= -1
			}
			return &Decimal{sign: result.sign, layer: result.layer, mag: result.mag}
		}
	}

	eParts := strings.Split(s, "e")
	eCount := len(eParts) - 1

	// Handle numbers that are exactly floats (0 or 1 "e"s).
	if eCount == 0 {
		numberAttempt, _ := strconv.ParseFloat(s, 64)
		if !math.IsInf(numberAttempt, 0) {
			if !math.IsInf(numberAttempt, 0) && !math.IsNaN(numberAttempt) {
				return decimalFromFloat64(numberAttempt)
			}
		}
	} else if eCount == 1 {
		numberAttempt, _ := strconv.ParseFloat(s, 64)
		if !math.IsInf(numberAttempt, 0) && numberAttempt != 0 {
			return decimalFromFloat64(numberAttempt)
		}
	}

	// TODO: (e^N)X format

	if eCount < 1 {
		return &Decimal{sign: 0, layer: 0, mag: 0}
	}

	mantissa, _ := strconv.ParseFloat(eParts[0], 64)
	if mantissa == 0 {
		return &Decimal{sign: 0, layer: 0, mag: 0}
	}

	exponent, _ := strconv.ParseFloat(eParts[len(eParts)-1], 64)
	if eCount >= 2 {
		me, _ := strconv.ParseFloat(eParts[len(eParts)-2], 64)
		if !math.IsInf(me, 0) {
			exponent *= sign(me)
			exponent += fMagLog10(me)
		}
	}

	result := &Decimal{sign: sign(mantissa), layer: float64(eCount), mag: 0}
	if math.IsInf(mantissa, 0) {
		if eParts[0] == "-" {
			return &Decimal{sign: -1, layer: 0, mag: 0}
		} else {
			return &Decimal{sign: 1, layer: 0, mag: 0}
		}
	} else if eCount == 1 {
		return &Decimal{sign: sign(mantissa), layer: 1, mag: exponent + math.Log10(math.Abs(mantissa))}
	} else {
		if eCount == 2 {
			result2 := Multiply(dFC(1, 2, exponent), D(mantissa))
			return result2
		} else {
			result.mag = exponent
		}
	}

	result = result.Normalize()
	return result
}

func dFC_NN(sign float64, layer float64, mag float64) *Decimal {
	return &Decimal{sign: sign, layer: layer, mag: mag}
}
func dFC(sign float64, layer float64, mag float64) *Decimal {
	d := Decimal{sign: sign, layer: layer, mag: mag}
	d.Normalize()
	return &d
}

//	func dME_NN(mantissa, exponent float64) *Decimal {
//		return &Decimal{sign: sign(mantissa), layer: 1, mag: exponent + math.Log10(math.Abs(mantissa))}
//	}
func dME(mantissa, exponent float64) *Decimal {
	d := Decimal{sign: sign(mantissa), layer: 1, mag: exponent + math.Log10(math.Abs(mantissa))}
	d.Normalize()
	return &d
}

func (d *Decimal) GetMantissa() float64 {
	if d.sign == 0 {
		return 0
	} else if d.layer == 0 {
		exp := math.Floor(math.Log10(d.mag))
		var man float64
		if d.mag == 5e-324 {
			man = 5
		} else {
			man = d.mag / math.Pow(10, exp)
		}
		return d.sign * man
	} else if d.layer == 1 {
		residue := d.mag - math.Floor(d.mag)
		return d.sign * math.Pow(10, residue)
	} else {
		return d.sign
	}
}
func (d *Decimal) SetMantissa(value float64) {
	if d.layer <= 2 {
		newD := dME(value, d.GetExponent())
		d.sign = newD.sign
		d.layer = newD.layer
		d.mag = newD.mag
	} else {
		// lol
		d.sign = sign(value)
		if value == 0 {
			d.layer = 0
			d.mag = 0
		}
	}
}

func (d *Decimal) GetExponent() float64 {
	if d.sign == 0 {
		return 0
	} else if d.layer == 0 {
		return math.Floor(math.Log10(d.mag))
	} else if d.layer == 1 {
		return math.Floor(d.mag)
	} else if d.layer == 2 {
		return math.Floor(sign(d.mag) * math.Pow(10, math.Abs(d.mag)))
	} else {
		return d.mag * math.Inf(1)
	}
}
func (d *Decimal) SetExponent(value float64) {
	newD := dME(d.GetMantissa(), value)
	d.sign = newD.sign
	d.layer = newD.layer
	d.mag = newD.mag
}

func (d *Decimal) GetSign() float64 {
	return d.sign
}
func (d *Decimal) SetSign(value float64) {
	if value == 0 {
		d.sign = 0
		d.layer = 0
		d.mag = 0
	} else {
		d.sign = value
	}
}

func (d *Decimal) MantissaWithNDecimalPlaces(places int) float64 {
	m := d.GetMantissa()
	if math.IsNaN(m) {
		return math.NaN()
	}

	if m == 0 {
		return 0
	}

	return decimalPlaces(m, places)
}

func (d *Decimal) MagnitudeWithNDecimalPlaces(places int) float64 {
	if math.IsNaN(d.mag) {
		return math.NaN()
	}

	if d.mag == 0 {
		return 0
	}

	return decimalPlaces(d.mag, places)
}
