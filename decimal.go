package breaketernitygo

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
	decimalSource()
}

// D is a function that creates a new Decimal instance from a given source.
// The source can be another *Decimal, a string, or any numeric type.
// The function returns a pointer to the newly created Decimal instance.
func D(source interface{}) *Decimal {
	return decimalFromSource(source)
}

func decimalFromSource(value interface{}) *Decimal {
	switch v := value.(type) {
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
		base, _ := strconv.ParseFloat(pentationParts[0], 10)
		height, _ := strconv.ParseFloat(pentationParts[1], 10)
		heightParts := strings.Split(pentationParts[1], ";")
		payload := float64(1)
		if len(heightParts) == 2 {
			payload, _ = strconv.ParseFloat(heightParts[1], 10)
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
		base, _ := strconv.ParseFloat(tetrationParts[0], 10)
		height, _ := strconv.ParseFloat(tetrationParts[1], 10)
		heightParts := strings.Split(tetrationParts[1], ";")
		payload := float64(1)
		if len(heightParts) == 2 {
			payload, _ = strconv.ParseFloat(heightParts[1], 10)
			if math.IsInf(payload, 0) {
				payload = 1
			}
		}
		if !math.IsInf(base, 0) && !math.IsInf(height, 0) {
			result := Tetrate(D(base), height, D(payload), linearhyper4)
			return &Decimal{sign: result.sign, layer: result.layer, mag: result.mag}
		}
	}
	return &Decimal{sign: 1, layer: 0, mag: 0}
}

func dFC_NN(sign float64, mag, layer float64) *Decimal {
	return &Decimal{sign: sign, layer: layer, mag: mag}
}
func dFC(sign float64, mag, layer float64) *Decimal {
	d := Decimal{sign: sign, layer: layer, mag: mag}
	d.Normalize()
	return &d
}

func dME_NN(mantissa, exponent float64) *Decimal {
	return &Decimal{sign: sign(mantissa), layer: 1, mag: exponent + math.Log10(math.Abs(mantissa))}
}
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
		d = dME(value, d.GetExponent())
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
	d = dME(d.GetMantissa(), value)
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
