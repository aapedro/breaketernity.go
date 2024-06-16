package breaketernitygo

import (
	"math"
)

// Cmp returns 1 if d > other, -1 if d < other, and 0 if d == other.
func Cmp[DS DecimalSource](d DS, other DS) int {
	return D(d).Cmp(D(other))
}

// Cmp returns 1 if d > other, -1 if d < other, and 0 if d == other.
func (d *Decimal) Cmp(other *Decimal) int {
	if d.sign > other.sign {
		return 1
	}
	if d.sign < other.sign {
		return -1
	}
	return int(d.sign) * d.CmpAbs(other)
}

// CmpAbs returns 1 if |d| > |other|, -1 if |d| < |other| and 0 if |d| == |other|.
func CmpAbs[DS DecimalSource](d DS, other DS) int {
	return D(d).CmpAbs(D(other))
}

// CmpAbs returns 1 if |d| > |other|, -1 if |d| < |other| and 0 if |d| == |other|.
func (d *Decimal) CmpAbs(other *Decimal) int {
	var layera float64
	if d.mag > 0 {
		layera = d.layer
	} else {
		layera = -d.layer
	}

	var layerb float64
	if other.mag > 0 {
		layerb = other.layer
	} else {
		layerb = -other.layer
	}

	if layera > layerb {
		return 1
	}

	if layera < layerb {
		return -1
	}

	if d.mag > other.mag {
		return 1
	}

	if d.mag < other.mag {
		return -1
	}

	return 0
}

// Eq compares two Decimal values and returns true if they are equal, false otherwise.
func Eq[DS DecimalSource](d DS, other DS) bool {
	return D(d).Eq(D(other))
}

// Eq compares two Decimal values and returns true if they are equal, false otherwise.
func (d *Decimal) Eq(other *Decimal) bool {
	return d.sign == other.sign && d.mag == other.mag && d.layer == other.layer
}

// Neq compares two Decimal values and returns true if they are not equal, false otherwise.
func Neq[DS DecimalSource](d DS, other DS) bool {
	return D(d).Neq(D(other))
}

// Neq compares two Decimal values and returns true if they are not equal, false otherwise.
func (d *Decimal) Neq(other *Decimal) bool {
	return !d.Eq(other)
}

// Lt compares two Decimal values and returns true if d < other, false otherwise.
func Lt[DS DecimalSource](d DS, other DS) bool {
	return D(d).Lt(D(other))
}

// Lt compares two Decimal values and returns true if d < other, false otherwise.
func (d *Decimal) Lt(other *Decimal) bool {
	return d.Cmp(other) == -1
}

// Lte compares two Decimal values and returns true if d <= other, false otherwise.
func Lte[DS DecimalSource](d DS, other DS) bool {
	return D(d).Lte(D(other))
}

// Lte compares two Decimal values and returns true if d <= other, false otherwise.
func (d *Decimal) Lte(other *Decimal) bool {
	return !d.Gt(other)
}

// Gt compares two Decimal values and returns true if d > other, false otherwise.
func Gt[DS DecimalSource](d DS, other DS) bool {
	return D(d).Gt(D(other))
}

// Gt compares two Decimal values and returns true if d > other, false otherwise.
func (d *Decimal) Gt(other *Decimal) bool {
	return d.Cmp(other) == 1
}

// Gte compares two Decimal values and returns true if d >= other, false otherwise.
func Gte[DS DecimalSource](d DS, other DS) bool {
	return D(d).Gte(D(other))
}

// Gte compares two Decimal values and returns true if d >= other, false otherwise.
func (d *Decimal) Gte(other *Decimal) bool {
	return !d.Lt(other)
}

// Max returns the maximum of two Decimal values.
func Max[DS DecimalSource](d DS, other DS) *Decimal {
	return D(d).Max(D(other))
}

// Max returns the maximum of two Decimal values.
func (d *Decimal) Max(other *Decimal) *Decimal {
	if d.Lt(other) {
		return d
	} else {
		return other
	}
}

// Min returns the minimum of two Decimal values.
func Min[DS DecimalSource](d DS, other DS) *Decimal {
	return D(d).Min(D(other))
}

// Min returns the minimum of two Decimal values.
func (d *Decimal) Min(other *Decimal) *Decimal {
	if d.Gt(other) {
		return d
	} else {
		return other
	}
}

// MaxAbs returns the Decimal with the maximum absolute value between d and other.
func MaxAbs[DS DecimalSource](d DS, other DS) *Decimal {
	return D(d).MaxAbs(D(other))
}

// MaxAbs returns the Decimal with the maximum absolute value between d and other.
func (d *Decimal) MaxAbs(other *Decimal) *Decimal {
	if d.CmpAbs(other) < 0 {
		return d
	} else {
		return other
	}
}

// MinAbs returns the Decimal with the minimum absolute value between d and other.
func MinAbs[DS DecimalSource](d DS, other DS) *Decimal {
	return D(d).MinAbs(D(other))
}

// MinAbs returns the Decimal with the minimum absolute value between d and other.
func (d *Decimal) MinAbs(other *Decimal) *Decimal {
	if d.CmpAbs(other) > 0 {
		return d
	} else {
		return other
	}
}

// Clamp is a combination of minimum and maximum.
// If d < min, returns min, and if d > max, returns max.
func Clamp[DS DecimalSource](d DS, min DS, max DS) *Decimal {
	return D(d).Clamp(D(min), D(max))
}

// Clamp is a combination of minimum and maximum.
// If d < min, returns min, and if d > max, returns max.
func (d *Decimal) Clamp(min *Decimal, max *Decimal) *Decimal {
	return d.Max(min).Min(max)
}

// ClampMin returns d, unless d is less than min, in which case returns min.
func ClampMin[DS DecimalSource](d DS, min DS) *Decimal {
	return D(d).Max(D(min))
}

// ClampMin returns d, unless d is less than min, in which case returns min.
func (d *Decimal) ClampMin(min *Decimal) *Decimal {
	return d.Max(min)
}

// ClampMax returns d, unless d is greater than max, in which case returns max.
func ClampMax[DS DecimalSource](d DS, max DS) *Decimal {
	return D(d).Min(D(max))
}

// ClampMax returns d, unless d is greater than max, in which case returns max.
func (d *Decimal) ClampMax(max *Decimal) *Decimal {
	return d.Min(max)
}

// EqTolerance compares two Decimal values and returns true if they are equal within a given tolerance.
// Tolerance is a relative tolerance, multiplied by the maximum of the magnitudes of the two values.
func EqTolerance[DS DecimalSource](d DS, other DS, tolerance float64) bool {
	return D(d).EqTolerance(D(other), tolerance)
}

// EqTolerance compares two Decimal values and returns true if they are equal within a given tolerance.
// Tolerance is a relative tolerance, multiplied by the maximum of the magnitudes of the two values.
func (d *Decimal) EqTolerance(other *Decimal, tolerance float64) bool {
	if tolerance == 0 {
		tolerance = 1e-7
	}
	// Numbers that are too far away are never close.
	if d.sign != other.sign {
		return false
	}
	if math.Abs(d.layer-other.layer) > 1 {
		return false
	}
	magA := d.mag
	magB := other.mag
	if d.layer > other.layer {
		magB = fMagLog10(magB)
	}
	if d.layer < other.layer {
		magA = fMagLog10(magA)
	}
	// return abs(a-b) <= tolerance * max(abs(a), abs(b))
	return math.Abs(magA-magB) <= tolerance*math.Max(math.Abs(magA), math.Abs(magB))
}

// CmpTolerance compares two Decimal values and returns -1 if d < other, 0 if d = other, 1 if d > other.
// However, the two decimals are considered equal if they are within a given tolerance.
// Tolerance is a relative tolerance, multiplied by the maximum of the magnitudes of the two values.
func CmpTolerance[DS DecimalSource](d DS, other DS, tolerance float64) int {
	return D(d).CmpTolerance(D(other), tolerance)
}

// CmpTolerance compares two Decimal values and returns -1 if d < other, 0 if d = other, 1 if d > other.
// However, the two decimals are considered equal if they are within a given tolerance.
// Tolerance is a relative tolerance, multiplied by the maximum of the magnitudes of the two values.
func (d *Decimal) CmpTolerance(other *Decimal, tolerance float64) int {
	if d.EqTolerance(other, tolerance) {
		return 0
	} else {
		return d.Cmp(other)
	}
}

// NeqTolerance compares two Decimal values and returns true if they are not equal within a given tolerance.
// Tolerance is a relative tolerance, multiplied by the maximum of the magnitudes of the two values.
func NeqTolerance[DS DecimalSource](d DS, other DS, tolerance float64) bool {
	return D(d).NeqTolerance(D(other), tolerance)
}

// NeqTolerance compares two Decimal values and returns true if they are not equal within a given tolerance.
// Tolerance is a relative tolerance, multiplied by the maximum of the magnitudes of the two values.
func (d *Decimal) NeqTolerance(other *Decimal, tolerance float64) bool {
	return !d.EqTolerance(other, tolerance)
}

// LtTolerance compares two Decimal values and returns true if d < other.
// However, the two decimals are considered equal if they are within a given tolerance.
// Tolerance is a relative tolerance, multiplied by the maximum of the magnitudes of the two values
func LtTolerance[DS DecimalSource](d DS, other DS, tolerance float64) bool {
	return D(d).LtTolerance(D(other), tolerance)
}

// LtTolerance compares two Decimal values and returns true if d < other.
// However, the two decimals are considered equal if they are within a given tolerance.
// Tolerance is a relative tolerance, multiplied by the maximum of the magnitudes of the two values
func (d *Decimal) LtTolerance(other *Decimal, tolerance float64) bool {
	return !d.EqTolerance(other, tolerance) && d.Lt(other)
}

// LtTolerance compares two Decimal values and returns true if d <= other.
// However, the two decimals are considered equal if they are within a given tolerance.
// Tolerance is a relative tolerance, multiplied by the maximum of the magnitudes of the two values
func LteTolerance[DS DecimalSource](d DS, other DS, tolerance float64) bool {
	return D(d).LteTolerance(D(other), tolerance)
}

// LtTolerance compares two Decimal values and returns true if d <= other.
// However, the two decimals are considered equal if they are within a given tolerance.
// Tolerance is a relative tolerance, multiplied by the maximum of the magnitudes of the two values
func (d *Decimal) LteTolerance(other *Decimal, tolerance float64) bool {
	return d.EqTolerance(other, tolerance) || d.Lt(other)
}

// GtTolerance compares two Decimal values and returns true if d > other.
// However, the two decimals are considered equal if they are within a given tolerance.
// Tolerance is a relative tolerance, multiplied by the maximum of the magnitudes of the two
func GtTolerance[DS DecimalSource](d DS, other DS, tolerance float64) bool {
	return D(d).GtTolerance(D(other), tolerance)
}

// GtTolerance compares two Decimal values and returns true if d > other.
// However, the two decimals are considered equal if they are within a given tolerance.
// Tolerance is a relative tolerance, multiplied by the maximum of the magnitudes of the two values
func (d *Decimal) GtTolerance(other *Decimal, tolerance float64) bool {
	return !d.EqTolerance(other, tolerance) && d.Gt(other)
}

// GteTolerance compares two Decimal values and returns true if d >= other.
// However, the two decimals are considered equal if they are within a given tolerance.
// Tolerance is a relative tolerance, multiplied by the maximum of the magnitudes of the two values
func GteTolerance[DS DecimalSource](d DS, other DS, tolerance float64) bool {
	return D(d).GteTolerance(D(other), tolerance)
}

// GteTolerance compares two Decimal values and returns true if d >= other.
// However, the two decimals are considered equal if they are within a given tolerance.
// Tolerance is a relative tolerance, multiplied by the maximum of the magnitudes of the two values
func (d *Decimal) GteTolerance(other *Decimal, tolerance float64) bool {
	return d.EqTolerance(other, tolerance) || d.Gt(other)
}

// PLog10 returns the base10 logarithm of non-negative decimals and returns 0 for negative decimals.
func PLog10[DS DecimalSource](d DS) *Decimal {
	return D(d).PLog10()
}

// PLog10 returns the base10 logarithm of non-negative decimals and returns 0 for negative decimals.
func (d *Decimal) PLog10() *Decimal {
	if d.Lt(dZero) {
		return dFC_NN(0, 0, 0)
	}
	return d.Log10()
}

// AbsLog10 returns the base10 logarithm of the absolute value of the decimal
func AbsLog10[DS DecimalSource](d DS) *Decimal {
	return D(d).AbsLog10()
}

// AbsLog10 returns the base10 logarithm of the absolute value of the decimal
func (d *Decimal) AbsLog10() *Decimal {
	if d.sign == 0 {
		return dFC_NN(math.NaN(), math.NaN(), math.NaN())
	} else if d.layer > 0 {
		return dFC(sign(d.mag), d.layer-1, math.Abs(d.mag))
	} else {
		return dFC(1, 0, math.Log10(d.mag))
	}
}

// Log10 returns the base10 logarithm of the decimal
func Log10[DS DecimalSource](d DS) *Decimal {
	return D(d).Log10()
}

// Log10 returns the base10 logarithm of the decimal
func (d *Decimal) Log10() *Decimal {
	if d.sign <= 0 {
		return dFC_NN(math.NaN(), math.NaN(), math.NaN())
	} else if d.layer > 0 {
		return dFC(sign(d.mag), d.layer-1, math.Abs(d.mag))
	} else {
		return dFC(d.sign, 0, math.Log10(d.mag))
	}
}

// Log returns the logarithm of the decimal to the given base
func Log[DS DecimalSource](d DS, base DS) *Decimal {
	return D(d).Log(D(base))
}

// Log returns the logarithm of the decimal to the given base
func (d *Decimal) Log(base *Decimal) *Decimal {
	if d.sign <= 0 {
		return dFC_NN(math.NaN(), math.NaN(), math.NaN())
	}
	if base.sign <= 0 {
		return dFC_NN(math.NaN(), math.NaN(), math.NaN())
	}
	if base.sign == 1 && base.layer == 0 && base.mag == 1 {
		return dFC_NN(math.NaN(), math.NaN(), math.NaN())
	}
	if d.layer == 0 && base.layer == 0 {
		return dFC(d.sign, 0, math.Log(d.mag)/math.Log(base.mag))
	}

	return Divide(d.Log10(), base.Log10())
}

// Ln returns the natural logarithm of the decimal
func Ln[DS DecimalSource](d DS) *Decimal {
	return D(d).Ln()
}

// Ln returns the natural logarithm of the decimal
func (d *Decimal) Ln() *Decimal {
	if d.sign <= 0 {
		return dFC_NN(math.NaN(), math.NaN(), math.NaN())
	} else if d.layer == 0 {
		return dFC(d.sign, 0, math.Log(d.mag))
	} else if d.layer == 1 {
		return dFC(sign(d.mag), 0, math.Abs(d.mag)*2.302585092994046) // ln(10)
	} else if d.layer == 2 {
		return dFC(sign(d.mag), 1, math.Abs(d.mag)+0.36221568869946325) // log10(log10(e))
	} else {
		return dFC(sign(d.mag), d.layer-1, math.Abs(d.mag))
	}
}

// Log2 returns the base2 logarithm of the decimal
func Log2[DS DecimalSource](d DS) *Decimal {
	return D(d).Log2()
}

// Log2 returns the base2 logarithm of the decimal
func (d *Decimal) Log2() *Decimal {
	if d.sign <= 0 {
		return dFC_NN(math.NaN(), math.NaN(), math.NaN())
	} else if d.layer == 0 {
		return dFC(d.sign, 0, math.Log2(d.mag))
	} else if d.layer == 1 {
		return dFC(sign(d.mag), 0, math.Abs(d.mag)*3.321928094887362) // ln(2)
	} else if d.layer == 2 {
		return dFC(sign(d.mag), 1, math.Abs(d.mag)+0.5213902276543247) // log10(log2(e))
	} else {
		return dFC(sign(d.mag), d.layer-1, math.Abs(d.mag))
	}
}

// Pow returns the decimal raised to the power of the other decimal
func Pow[DS DecimalSource](d DS, other DS) *Decimal {
	return D(d).Pow(D(other))
}

// Pow returns the decimal raised to the power of the other decimal
func (d *Decimal) Pow(other *Decimal) *Decimal {
	a := D(d)
	b := D(other)

	if a.sign == 0 {
		if b.Eq(D(0)) {
			return dFC_NN(1, 0, 1)
		} else {
			return a
		}
	}

	if a.sign == 1 && a.layer == 0 && a.mag == 1 {
		return a
	}

	if b.sign == 0 {
		return dFC_NN(1, 0, 1)
	}

	if b.sign == 1 && b.layer == 0 && b.mag == 1 {
		return a
	}

	result := a.AbsLog10().Multiply(b).PowBase10()

	if d.sign == -1 {
		if math.Mod(math.Abs(math.Mod(b.ToFloat64(), 2)), 2) == 1 {
			return result.Neg()
		} else if math.Mod(math.Abs(math.Mod(b.ToFloat64(), 2)), 2) == 0 {
			return result
		}
		return dFC_NN(math.NaN(), math.NaN(), math.NaN())
	}

	return result
}

// PowBase10 returns 10 raised to the power of the decimal
func PowBase10[DS DecimalSource](d DS) *Decimal {
	return D(d).PowBase10()
}

// PowBase10 returns 10 raised to the power of the decimal
func (d *Decimal) PowBase10() *Decimal {
	// Handle infinity cases
	if d.Eq(dInf) {
		return dFC_NN(1, math.Inf(1), math.Inf(1))
	}
	if d.Eq(dNegInf) {
		return dFC_NN(0, 0, 0)
	}
	// Handle non-finite layer or magnitude
	if math.IsInf(d.layer, 0) || math.IsInf(d.mag, 0) {
		return dFC_NN(math.NaN(), math.NaN(), math.NaN())
	}

	a := D(d)
	// Handle layer 0 case
	if a.layer == 0 {
		newMag := math.Pow(10, a.sign*a.mag)
		if !math.IsInf(newMag, 0) && math.Abs(newMag) >= 0.1 {
			return dFC(1, 0, newMag)
		} else {
			if a.sign == 0 {
				return dFC_NN(1, 0, 1)
			} else {
				return dFC_NN(a.sign, a.layer+1, math.Log10(a.mag))
			}
		}
	}

	// Handle all 4 layer 1+ cases
	if a.sign > 0 && a.mag >= 0 {
		return dFC(a.sign, a.layer+1, a.mag)
	}
	if a.sign < 0 && a.mag >= 0 {
		return dFC(-a.sign, a.layer+1, -a.mag)
	}
	// Both negative mag cases result in the same outcome
	return dFC_NN(1, 0, 1)
}

// PowBaseE returns e raised to the power of the decimal
func PowBaseE[DS DecimalSource](d DS) *Decimal {
	return D(d).PowBaseE()
}

// PowBaseE returns e raised to the power of the decimal
func (d *Decimal) PowBaseE() *Decimal {
	if d.mag < 0 {
		return dFC_NN(1, 0, 1)
	}
	if d.layer == 0 && d.mag <= 709.7 {
		return D(math.Exp(d.sign * d.mag))
	} else if d.layer == 0 {
		return dFC(1, 1, d.sign*math.Log10(math.E)*d.mag)
	} else if d.layer == 1 {
		return dFC(1, 2, d.sign*(math.Log10(0.4342944819032518)+d.mag))
	} else {
		return dFC(1, d.layer+1, d.sign*d.mag)
	}
}

// PowBaseN returns the base raised to the power of the decimal
func PowBaseN[DS DecimalSource](d DS, base DS) *Decimal {
	return D(d).PowBaseN(D(base))
}

// PowBaseN returns the base raised to the power of the decimal
func (d *Decimal) PowBaseN(base *Decimal) *Decimal {
	return d.Pow(base)
}

// Root returns the "degree"th root of the decimal
func Root[DS DecimalSource](d DS, degree DS) *Decimal {
	return D(d).Root(D(degree))
}

// Root returns the "degree"th root of the decimal
func (d *Decimal) Root(degree *Decimal) *Decimal {
	return d.Pow(degree.Recip())
}

// Sqrt returns the square root of the decimal
func Sqrt[DS DecimalSource](d DS) *Decimal {
	return D(d).Sqrt()
}

// Sqrt returns the square root of the decimal
func (d *Decimal) Sqrt() *Decimal {
	if d.layer == 0 {
		return decimalFromFloat64(math.Sqrt(d.sign * d.mag))
	} else if d.layer == 1 {
		return dFC(1, 2, math.Log10(d.mag)-0.3010299956639812)
	} else {
		result := Divide(dFC_NN(d.sign, d.layer-1, d.mag), dFC_NN(1, 0, 2))
		result.layer += 1
		result = result.Normalize()
		return result
	}
}

// Factorial returns the factorial of the decimal
// This function is extended to all real numbers via the Gamma function
func Factorial[DS DecimalSource](d DS) *Decimal {
	return D(d).Factorial()
}

// Factorial returns the factorial of the decimal
// This function is extended to all real numbers via the Gamma function
func (d *Decimal) Factorial() *Decimal {
	if d.mag < 0 {
		return d.Add(D(1)).Gamma()
	} else if d.layer == 0 {
		return d.Add(D(1)).Gamma()
	} else if d.layer == 1 {
		return d.Multiply(d.Ln().Subtract(D(1))).PowBase10()
	} else {
		return d.PowBase10()
	}
}

// Gamma returns the Gamma function of the decimal
// Gamma(x) is defined as the integral of t^(x-1) * e^-t dt from t = 0 to t = infinity
// This is equivalent to (x-1)! for nonnegative integers
func Gamma[DS DecimalSource](d DS) *Decimal {
	return D(d).Gamma()
}

// Gamma returns the Gamma function of the decimal
// Gamma(x) is defined as the integral of t^(x-1) * e^-t dt from t = 0 to t = infinity
// This is equivalent to (x-1)! for nonnegative integers
func (d *Decimal) Gamma() *Decimal {
	if d.mag < 0 {
		return d.Recip()
	} else if d.layer == 0 {
		if d.Lt(dFC_NN(1, 0, 24)) {
			return D(fGamma(d.sign * d.mag))
		}

		t := d.mag - 1
		l := 0.9189385332046727 //0.5*Math.log(2*Math.PI)
		l = l + (t+0.5)*math.Log(t)
		l = l - t

		n2 := t * t
		np := t
		lm := 12 * np
		adj := 1 / lm
		l2 := l + adj
		if l2 == l {
			return D(l).PowBaseE()
		}

		l = l2
		np = np * n2
		lm = 360 * np
		adj = 1 / lm
		l2 = l - adj
		if l2 == l {
			return D(l).PowBaseE()
		}

		l = l2
		np = np * n2
		lm = 1260 * np
		lt := 1 / lm
		l = l + lt
		np = np * n2
		lm = 1680 * np
		lt = 1 / lm
		l = l - lt
		return D(l).PowBaseE()
	} else if d.layer == 1 {
		return d.Multiply(d.Ln().Subtract(D(1))).PowBaseE()
	} else {
		return d.PowBaseE()
	}
}

// Abs returns the absolute value of the decimal
func Abs[DS DecimalSource](d DS) *Decimal {
	return D(d).Abs()
}

// Abs returns the absolute value of the decimal
func (d *Decimal) Abs() *Decimal {
	if d.sign == 0 {
		return dFC_NN(0, d.layer, d.mag)
	} else {
		return dFC_NN(1, d.layer, d.mag)
	}
}

// Neg returns the negative of the decimal
func Neg[DS DecimalSource](d DS) *Decimal {
	return D(d).Neg()
}

// Neg returns the negative of the decimal
func (d *Decimal) Neg() *Decimal {
	return dFC_NN(sign(-1*d.mag), d.layer, d.mag)
}

// Round rounds the decimal to the nearest integer
func Round[DS DecimalSource](d DS) *Decimal {
	return D(d).Round()
}

// Round rounds the decimal to the nearest integer
func (d *Decimal) Round() *Decimal {
	if d.mag < 0 {
		return dFC_NN(0, 0, 0)
	}
	if d.layer == 0 {
		return dFC(d.sign, 0, math.Round(d.mag))
	}
	return D(d)
}

// Floor rounds the decimal down to the nearest integer
func Floor[DS DecimalSource](d DS) *Decimal {
	return D(d).Floor()
}

// Floor rounds the decimal down to the nearest integer
func (d *Decimal) Floor() *Decimal {
	if d.mag < 0 {
		if d.sign == -1 {
			return dFC_NN(-1, 0, 1)
		} else {
			return dFC_NN(0, 0, 0)
		}
	}
	if d.sign == -1 {
		return d.Neg().Ceil().Neg()
	}
	if d.layer == 0 {
		return dFC(d.sign, 0, math.Floor(d.mag))
	}
	return D(d)
}

// Ceil rounds the decimal up to the nearest integer
func Ceil[DS DecimalSource](d DS) *Decimal {
	return D(d).Ceil()
}

// Ceil rounds the decimal up to the nearest integer
func (d *Decimal) Ceil() *Decimal {
	if d.mag < 0 {
		if d.sign == -1 {
			return dFC_NN(1, 0, 1)
		} else {
			return dFC_NN(0, 0, 0)
		}
	}
	if d.sign == 1 {
		return d.Neg().Floor().Neg()
	}
	if d.sign == 0 {
		return dFC(d.sign, 0, math.Ceil(d.mag))
	}
	return D(d)
}

// Trunc returns the integer part of the Decimal
// Behaves like floor on positive numbers and ceil on negative numbers
func Trunc[DS DecimalSource](d DS) *Decimal {
	return D(d).Trunc()
}

// Trunc returns the integer part of the Decimal.
// Behaves like floor on positive numbers and ceil on negative numbers
func (d *Decimal) Trunc() *Decimal {
	if d.mag < 0 {
		return dFC_NN(0, 0, 0)
	}
	if d.layer == 0 {
		return dFC(d.sign, 0, math.Trunc(d.mag))
	}
	return D(d)
}

// Recip returns the reciprocal (1/x) of the decimal
func Recip[DS DecimalSource](d DS) *Decimal {
	return D(d).Recip()
}

// Recip returns the reciprocal (1/x) of the decimal
func (d *Decimal) Recip() *Decimal {
	if d.mag == 0 {
		return dFC_NN(math.NaN(), math.NaN(), math.NaN())
	} else if d.mag == math.Inf(1) {
		return dFC_NN(0, 0, 0)
	} else if d.layer == 0 {
		return dFC(d.sign, 0, 1/d.mag)
	} else {
		return dFC(d.sign, d.layer, -d.mag)
	}
}

// Add returns the sum of the decimal and other
func Add[DS DecimalSource](d DS, other DS) *Decimal {
	return D(d).Add(D(other))
}

// Add returns the sum of the decimal and other
func (d *Decimal) Add(other *Decimal) *Decimal {
	if (d.Eq(dInf) && other.Eq(dInf)) || (d.Eq(dNegInf) && other.Eq(dNegInf)) {
		return dFC_NN(math.NaN(), math.NaN(), math.NaN())
	}

	if math.IsInf(d.layer, 0) {
		return D(d)
	}
	if math.IsInf(other.layer, 0) {
		return D(other)
	}

	if d.sign == 0 {
		return D(other)
	}
	if other.sign == 0 {
		return D(d)
	}

	if d.sign == -other.sign && d.layer == other.layer && d.mag == d.mag {
		return dFC_NN(0, 0, 0)
	}

	var a *Decimal
	var b *Decimal

	if d.layer >= 2 || other.layer >= 2 {
		return d.MaxAbs(other)
	}

	if CmpAbs(d, other) > 0 {
		a = D(d)
		b = D(other)
	} else {
		a = D(other)
		b = D(d)
	}

	if a.layer == 0 && b.layer == 0 {
		return decimalFromFloat64(a.sign*a.mag + b.sign*b.mag)
	}

	layera := a.layer * sign(a.mag)
	layerb := b.layer * sign(b.mag)

	if layera-layerb >= 2 {
		return a
	}

	if layera == 0 && layerb == -1 {
		if math.Abs(b.mag-math.Log10(a.mag)) > float64(MAX_SIGNIFICANT_DIGITS) {
			return a
		} else {
			magDiff := math.Pow(10, math.Log10(a.mag)-b.mag)
			mantissa := b.sign + a.sign*magDiff
			return dFC(sign(mantissa), 1, b.mag+math.Log10(math.Abs(mantissa)))
		}
	}

	if layera == 1 && layerb == 0 {
		if math.Abs(a.mag-math.Log10(b.mag)) > float64(MAX_SIGNIFICANT_DIGITS) {
			return a
		} else {
			magDiff := math.Pow(10, a.mag-math.Log10(b.mag))
			mantissa := b.sign + a.sign*magDiff
			return dFC(sign(mantissa), 1, math.Log10(b.mag)+math.Log10(math.Abs(mantissa)))
		}
	}

	if math.Abs(a.mag-b.mag) > float64(MAX_SIGNIFICANT_DIGITS) {
		return a
	} else {
		magDiff := math.Pow(10, a.mag-b.mag)
		mantissa := b.sign + a.sign*magDiff
		return dFC(sign(mantissa), 1, b.mag+math.Log10(math.Abs(mantissa)))
	}
}

// Subtract returns the difference between the decimal and other
func Subtract[DS DecimalSource](d DS, other DS) *Decimal {
	return D(d).Subtract(D(other))
}

// Subtract returns the difference between the decimal and other
func (d *Decimal) Subtract(other *Decimal) *Decimal {
	return d.Add(D(other).Neg())
}

// Multiply returns the product of the decimal and other
func Multiply[DS DecimalSource](d DS, other DS) *Decimal {
	return D(d).Multiply(D(other))
}

// Multiply returns the product of the decimal and other
func (d *Decimal) Multiply(other *Decimal) *Decimal {
	// infinity * -infinity = -infinity
	if (d.Eq(dInf) && other.Eq(dNegInf)) || (d.Eq(dNegInf) && other.Eq(dInf)) {
		return dFC_NN(-1, math.Inf(1), math.Inf(1))
	}

	if (d.mag == math.Inf(1) && other.Eq(dZero)) || (d.Eq(dZero) && other.mag == math.Inf(1)) {
		return dFC_NN(math.NaN(), math.NaN(), math.NaN())
	}

	if math.IsInf(d.layer, 0) {
		return D(d)
	}
	if math.IsInf(other.layer, 0) {
		return D(other)
	}

	if d.sign == 0 || other.sign == 0 {
		return dFC_NN(0, 0, 0)
	}

	if d.layer == other.layer && d.mag == -other.mag {
		return dFC_NN(d.sign*other.sign, 0, 1)
	}

	var a *Decimal
	var b *Decimal

	//Which number is bigger in terms of its multiplicative distance from 1?
	if d.layer > other.layer || (d.layer == other.layer && math.Abs(d.mag) > math.Abs(other.mag)) {
		a = D(d)
		b = D(other)
	} else {
		a = D(other)
		b = D(d)
	}

	if a.layer == 0 && b.layer == 0 {
		return decimalFromFloat64(a.sign * a.mag * b.sign * b.mag)
	}

	if a.layer >= 3 || a.layer-b.layer >= 2 {
		return dFC(a.sign*b.sign, a.layer, a.mag)
	}

	if a.layer == 1 && b.layer == 0 {
		return dFC(a.sign*b.sign, 1, a.mag+math.Log10(b.mag))
	}

	if a.layer == 1 && b.layer == 1 {
		return dFC(a.sign*b.sign, 1, a.mag+b.mag)
	}

	if a.layer == 2 && b.layer == 1 {
		newMag := dFC(sign(a.mag), a.layer-1, math.Abs(a.mag)).Add(dFC(sign(b.mag), b.layer-1, math.Abs(b.mag)))
		return dFC(a.sign*b.sign, newMag.layer+1, newMag.sign*newMag.mag)
	}

	if a.layer == 2 && b.layer == 2 {
		newMag := dFC(sign(a.mag), a.layer-1, math.Abs(a.mag)).Add(dFC(sign(b.mag), b.layer-1, math.Abs(b.mag)))
		return dFC(a.sign*b.sign, newMag.layer+1, newMag.sign*newMag.mag)
	}

	panic("Bad arguments to multiply: " + d.ToString() + ", " + other.ToString())
}

// Divide returns the quotient of the decimal and other
func Divide[DS DecimalSource](d DS, other DS) *Decimal {
	return D(d).Divide(D(other))
}

// Divide returns the quotient of the decimal and other
func (d *Decimal) Divide(other *Decimal) *Decimal {
	return d.Multiply(other.Recip())
}

// Modulo returns the remainder of d divided by other
// Uses the truncated division modulo, which is the same as Go's native modulo operator (%)
func Modulo[DS DecimalSource](d DS, other DS) *Decimal {
	return D(d).Modulo(D(other))
}

// Modulo returns the remainder of d divided by other
// Uses the truncated division modulo, which is the same as Go's native modulo operator (%)
func (d *Decimal) Modulo(other *Decimal) *Decimal {
	if other.Eq(dZero) {
		return dFC_NN(0, 0, 0)
	}

	dNum := d.ToFloat64()
	otherNum := other.ToFloat64()

	if !math.IsInf(dNum, 0) && !math.IsInf(otherNum, 0) && dNum != 0 && otherNum != 0 {
		return D(math.Mod(dNum, otherNum))
	}

	if d.Subtract(other).Eq(d) {
		return dFC_NN(0, 0, 0)
	}

	if other.Subtract(d).Eq(other) {
		return D(d)
	}

	if d.sign == -1 {
		return d.Abs().Modulo(d).Neg()
	}

	return d.Subtract(d.Divide(d).Floor().Multiply(other))
}

// IsNan returns true if the decimal is NaN
func IsNaN[DS DecimalSource](d DS) bool {
	return D(d).IsNaN()
}

// IsNaN returns true if the decimal is NaN
func (d *Decimal) IsNaN() bool {
	return math.IsNaN(d.layer) || math.IsNaN(d.mag) || math.IsNaN(d.sign)
}

// IsInf returns true if the decimal is either positive or negative infinity
func IsInf[DS DecimalSource](d DS) bool {
	return D(d).IsInf()
}

// IsInf returns true if the decimal is either positive or negative infinity
func (d *Decimal) IsInf() bool {
	return math.IsInf(d.layer, 0) || math.IsInf(d.mag, 0) || math.IsInf(d.sign, 0)
}

func dLambertW(d *Decimal, tol float64, principal bool) *Decimal {
	var w, ew, wewz, wn *Decimal
	if math.IsInf(d.mag, 0) {
		return decimalFromDecimal(d)
	}
	if principal {
		if d.Eq(dZero) {
			return dFC_NN(0, 0, 0)
		}
		if d.Eq((dOne)) {
			return decimalFromFloat64(OMEGA)
		}

		w = Ln(d)
	} else {
		if d.Eq(dZero) {
			return dFC_NN(-1, math.Inf(-1), math.Inf(-1))
		}

		w = Ln(d.Neg())
	}

	for i := 0; i < 100; i++ {
		ew = w.Neg().PowBaseE()
		wewz = w.Subtract(d.Multiply(ew))
		wn = w.Subtract(wewz.Divide(w.Add(D(1)).Subtract(w.Add(D(2)).Multiply(wewz).Divide(Multiply(D(2), w).Add(D(2))))))
		if Abs(wn.Subtract(w)).Lt(Abs(wn).Multiply(D(tol))) {
			return wn
		} else {
			w = wn
		}
	}

	panic("Iteration failed to converge")
}

// LambertW is an implementation of the Lambert W function, also called the omega function or the product logarithm.
// Solution to W(X) == x*e^x
// This is a multi-valued function in the complex plane but only two "branches" matter for real numbers. W0 (principal) and W-1 (non-principal)
// W0 works for any number >= -1/e, but W-1 only works for nonpositive numbers >= -1/e
// The principal paremeter determines which branch to use
func LambertW[DS DecimalSource](d DS, principal bool) *Decimal {
	return D(d).LambertW(principal)
}

// LambertW is an implementation of the Lambert W function, also called the omega function or the product logarithm.
// Solution to W(X) == x*e^x
// This is a multi-valued function in the complex plane but only two "branches" matter for real numbers. W0 (principal) and W-1 (non-principal)
// W0 works for any number >= -1/e, but W-1 only works for nonpositive numbers >= -1/e
// The principal paremeter determines which branch to use
func (d *Decimal) LambertW(principal bool) *Decimal {
	if d.Lt(D(-0.3678794411710499)) {
		return dFC_NN(math.NaN(), math.NaN(), math.NaN())
	} else if principal {
		if d.Abs().Lt(D("1e-300")) {
			return decimalFromDecimal(d)
		} else if d.mag < 0 {
			return decimalFromFloat64(fLambertW(d.ToFloat64(), 1e-10, true))
		} else if d.layer == 0 {
			return decimalFromFloat64(fLambertW(d.sign*d.mag, 1e-100, true))
		} else if d.Lt(D("eee15")) {
			return dLambertW(d, 1e-10, true)
		} else {
			return d.Ln()
		}
	} else {
		if d.sign == 1 {
			return dFC_NN(math.NaN(), math.NaN(), math.NaN())
		} else if d.layer == 0 {
			return decimalFromFloat64(fLambertW(d.sign*d.mag, 1e-10, false))
		} else if d.layer == 1 {
			return dLambertW(d, 1e-10, false)
		} else {
			return d.Neg().Recip().LambertW(true).Neg()
		}
	}
}

// IteratedLog returns the result of applying log(base) 'times' times
// Works with negative and positive real heights. Tetration for non-integer heights does not have a single agreed-upon definition
// So this library uses an analytic approximation for bases <= 10, but reverts to a linear approximation for bases > 10
// If you want to use the linear approximation for all bases, set linear parameter to true
func IteratedLog[DS DecimalSource](d DS, base DS, times float64, linear bool) *Decimal {
	return D(d).IteratedLog(D(base), times, linear)
}

// IteratedLog returns the result of applying log(base) 'times' times
// Works with negative and positive real heights. Tetration for non-integer heights does not have a single agreed-upon definition
// So this library uses an analytic approximation for bases <= 10, but reverts to a linear approximation for bases > 10
// If you want to use the linear approximation for all bases, set linear parameter to true
func (d *Decimal) IteratedLog(base *Decimal, times float64, linear bool) *Decimal {
	if times < 0 {
		return Tetrate(base, -times, d, linear)
	}

	result := D(d)
	fullTimes := times
	times = math.Trunc(times)
	fraction := fullTimes - times

	if result.layer-base.layer > 3 {
		layerloss := math.Min(times, result.layer-base.layer-3)
		times -= layerloss
		result.layer -= layerloss
	}

	for i := 0; i < int(times); i++ {
		result = result.Log(base)
		if math.IsInf(result.layer, 0) || math.IsInf(result.mag, 0) {
			return result.Normalize()
		}
		if i > 10000 {
			return result
		}
	}

	if fraction > 0 && fraction < 1 {
		if base.Eq(D(10)) {
			result = result.LayerAdd10(D(-fraction), linear)
		} else {
			result = result.LayerAdd(D(-fraction), base, linear)
		}
	}

	return result
}

// IteratedExp returns the result of applying exp(base) 'height' times.
// Works with negative and positive real heights. Tetration for non-integer heights does not have a single agreed-upon definition
// So this library uses an analytic approximation for bases <= 10, but reverts to a linear approximation for bases > 10
// If you want to use the linear approximation for all bases, set linear parameter to true
// Identical to Tetrate
func IteratedExp[DS DecimalSource](d DS, height float64, payload DS, linear bool) *Decimal {
	return D(d).IteratedExp(height, D(payload), linear)
}

// IteratedExp returns the result of applying exp(base) 'height' times.
// Works with negative and positive real heights. Tetration for non-integer heights does not have a single agreed-upon definition
// So this library uses an analytic approximation for bases <= 10, but reverts to a linear approximation for bases > 10
// If you want to use the linear approximation for all bases, set linear parameter to true
// Identical to Tetrate
func (d *Decimal) IteratedExp(height float64, payload *Decimal, linear bool) *Decimal {
	return d.Tetrate(height, payload, linear)
}

// LayerAdd10 adds/removes layers from a Decimal, even fractional layers. Very similar to tetrate base 10 and iterated log base 10
// Tetration for non-integer heights does not have a single agreed-upon definition
// So this library uses an analytic approximation for bases <= 10, but reverts to a linear approximation for bases > 10
// If you want to use the linear approximation for all bases, set linear parameter to true
func LayerAdd10[DS DecimalSource](d DS, diff DS, linear bool) *Decimal {
	return D(d).LayerAdd10(D(diff), linear)
}

// LayerAdd10 adds/removes layers from a Decimal, even fractional layers. Very similar to tetrate base 10 and iterated log base 10
// Tetration for non-integer heights does not have a single agreed-upon definition
// So this library uses an analytic approximation for bases <= 10, but reverts to a linear approximation for bases > 10
// If you want to use the linear approximation for all bases, set linear parameter to true
func (d *Decimal) LayerAdd10(diff *Decimal, linear bool) *Decimal {
	fDiff := D(diff).ToFloat64()
	result := decimalFromDecimal(d)

	if fDiff >= 1 {
		if result.mag < 0 && result.layer > 0 {
			result.sign = 0
			result.mag = 0
			result.layer = 0
		} else if result.sign == -1 && result.layer == 0 {
			result.sign = 1
			result.mag = -result.mag
		}
		layerAdd := math.Trunc(fDiff)
		fDiff -= layerAdd
		result.layer += layerAdd
	}

	if fDiff <= -1 {
		layerAdd := math.Trunc(fDiff)
		fDiff -= layerAdd
		result.layer += layerAdd
		if result.layer < 0 {
			for i := 0; i < 100; i++ {
				result.layer++
				result.mag = math.Log10(result.mag)
				if math.IsInf(result.mag, 0) {
					if result.sign == 0 {
						result.sign = 1
					}
					if result.layer < 0 {
						result.layer = 0
					}
					return result.Normalize()
				}
				if result.layer >= 0 {
					break
				}
			}
		}
	}

	for result.layer < 0 {
		result.layer++
		result.mag = math.Log10(result.mag)
	}

	if result.sign == 0 {
		result.sign = 1
		if result.mag == 0 && result.layer >= 1 {
			result.layer -= 1
			result.mag = 1
		}
	}
	result = result.Normalize()

	if fDiff != 0 {
		return result.LayerAdd(diff, D(10), linear)
	}

	return result
}

// LayerAdd is like adding "diff" to the number's slog(base) representation. Very similar to tetrate base 'base' and iterated log base 'base'
// Tetration for non-integer heights does not have a single agreed-upon definition
// So this library uses an analytic approximation for bases <= 10, but reverts to a linear approximation for bases > 10
// If you want to use the linear approximation for all bases, set linear parameter to true
func LayerAdd[DS DecimalSource](d DS, diff DS, base DS, linear bool) *Decimal {
	return D(d).LayerAdd(D(diff), D(base), linear)
}

// LayerAdd is like adding "diff" to the number's slog(base) representation. Very similar to tetrate base 'base' and iterated log base 'base'
// Tetration for non-integer heights does not have a single agreed-upon definition
// So this library uses an analytic approximation for bases <= 10, but reverts to a linear approximation for bases > 10
// If you want to use the linear approximation for all bases, set linear parameter to true
func (d *Decimal) LayerAdd(diff *Decimal, base *Decimal, linear bool) *Decimal {
	fDiff := diff.ToFloat64()
	if base.Gt(D(1)) && base.Lte(D(1.44466786100976613366)) {
		excessSlog, e1 := excessSlog(d, base, linear)
		slogThis := excessSlog.ToFloat64()
		range_ := e1
		slogDest := slogThis + fDiff
		negLn := base.Ln().Neg()
		lower := negLn.LambertW(true).Divide(negLn)
		upper := negLn.LambertW(false).Divide(negLn)
		slogZero := dOne
		if range_ == 1 {
			slogZero = lower.Multiply(upper).Sqrt()
		} else if range_ == 2 {
			slogZero = upper.Multiply(D(2))
		}
		slogOne := base.Pow(D(slogZero))
		wholeHeight := math.Floor(slogDest)
		fracHeight := slogDest - wholeHeight
		towerTop := slogZero.Pow(D(1 - fracHeight)).Multiply(slogOne.Pow(D(fracHeight)))
		return Tetrate(base, wholeHeight, towerTop, linear)
	}
	slogThis := d.Slog(base, 100, linear).ToFloat64()
	slogDest := slogThis + fDiff
	if slogDest >= 0 {
		return Tetrate(base, slogDest, dOne, linear)
	} else if math.IsInf(slogDest, 0) {
		return dFC_NN(math.NaN(), math.NaN(), math.NaN())
	} else if slogDest >= -1 {
		return Log(Tetrate(base, slogDest+1, dOne, linear), base)
	} else {
		return Log(Log(Tetrate(base, slogDest+2, dOne, linear), base), base)
	}
}

func (d *Decimal) slogInternal(base *Decimal, linear bool) *Decimal {
	if base.Lte(dZero) {
		return dFC_NN(math.NaN(), math.NaN(), math.NaN())
	}
	if base.Eq(dOne) {
		return dFC_NN(math.NaN(), math.NaN(), math.NaN())
	}
	if base.Lt(dOne) {
		if d.Eq(dOne) {
			return dFC_NN(0, 0, 0)
		}
		if d.Eq(dZero) {
			return dFC_NN(-1, 0, 1)
		}
		return dFC_NN(math.NaN(), math.NaN(), math.NaN())
	}

	if d.mag < 0 || d.Eq(dZero) {
		return dFC_NN(-1, 0, 1)
	}
	if base.Lt(D(1.44466786100976613366)) {
		negLn := base.Ln().Neg()
		infTower := negLn.LambertW(true).Divide(negLn)
		if d.Eq(infTower) {
			return dFC_NN(1, math.Inf(1), math.Inf(1))
		}
		if d.Gt(infTower) {
			return dFC_NN(math.NaN(), math.NaN(), math.NaN())
		}
	}

	result := 0.
	copy := decimalFromDecimal(d)
	if copy.layer-base.layer > 3 {
		layerLoss := copy.layer - base.layer - 3
		result += layerLoss
		copy.layer -= layerLoss
	}

	for i := 0; i < 100; i++ {
		if copy.Lt(dZero) {
			copy = base.Pow(copy)
			result -= 1
		} else if copy.Lte(dOne) {
			if linear {
				return decimalFromFloat64(result + copy.ToFloat64() - 1)
			} else {
				return decimalFromFloat64(result + slogCritical(base.ToFloat64(), copy.ToFloat64()))
			}
		} else {
			result += 1
			copy = copy.Log(base)
		}
	}

	return decimalFromFloat64(result)
}

// Slog is also called "super-logarithm". One of tetration's inverses, tells you what size pwoer tower you'd have to tetrate 'base' to get 'd'
// By definition, will never be higher than 1.8e308 in this lib, since a power tower 1.8e308 numbers tall is the largest representable number
// Accepts a number of iterations, and uses binary search to hone in on the true value
// Tetration for non-integer heights does not have a single agreed-upon definition
// So this library uses an analytic approximation for bases <= 10, but reverts to a linear approximation for bases > 10
// If you want to use the linear approximation for all bases, set linear parameter to true
func Slog[DS DecimalSource](d DS, base DS, iterations float64, linear bool) *Decimal {
	return D(d).Slog(D(base), iterations, linear)
}

// Slog is also called "super-logarithm". One of tetration's inverses, tells you what size pwoer tower you'd have to tetrate 'base' to get 'd'
// By definition, will never be higher than 1.8e308 in this lib, since a power tower 1.8e308 numbers tall is the largest representable number
// Accepts a number of iterations, and uses binary search to hone in on the true value
// Tetration for non-integer heights does not have a single agreed-upon definition
// So this library uses an analytic approximation for bases <= 10, but reverts to a linear approximation for bases > 10
// If you want to use the linear approximation for all bases, set linear parameter to true
func (d *Decimal) Slog(base *Decimal, iterations float64, linear bool) *Decimal {
	stepSize := 0.001
	hasChangedDirectionsOnce := false
	previouslyRose := false
	result := d.slogInternal(base, linear).ToFloat64()
	for i := 1; i < int(iterations); i++ {
		newDecimal := decimalFromDecimal(d).Tetrate(result, dOne, linear)
		currentlyRose := newDecimal.Gt(d)
		if i > 1 {
			if previouslyRose != currentlyRose {
				hasChangedDirectionsOnce = true
			}
		}
		previouslyRose = currentlyRose
		if hasChangedDirectionsOnce {
			stepSize /= 2
		} else {
			stepSize *= 2
		}
		if currentlyRose {
			stepSize = math.Abs(stepSize) * -1
		} else {
			stepSize = math.Abs(stepSize)
		}
		result += stepSize
		if stepSize == 0 {
			break
		}
	}

	return decimalFromFloat64(result)
}

// Tetrate is the result of exponentiating 'd' to 'payload' 'height' times in a row
// If payload != 1, this is the same as 'iterated exponentiation'
// Works with negative and positive real heights. Tetration for non-integer heights does not have a single agreed-upon definition
// So this library uses an analytic approximation for bases <= 10, but it reverts to the linear approximation for bases > 10
// If you want to use the linear approximation even for bases <= 10, set the linear parameter to true
func Tetrate[DS DecimalSource](d DS, height float64, payload DS, linear bool) *Decimal {
	return D(d).Tetrate(height, D(payload), linear)
}

// Tetrate is the result of exponentiating 'd' to 'payload' 'height' times in a row
// If payload != 1, this is the same as 'iterated exponentiation'
// Works with negative and positive real heights. Tetration for non-integer heights does not have a single agreed-upon definition
// So this library uses an analytic approximation for bases <= 10, but it reverts to the linear approximation for bases > 10
// If you want to use the linear approximation even for bases <= 10, set the linear parameter to true
func (d *Decimal) Tetrate(height float64, payload *Decimal, linear bool) *Decimal {
	if height == 1 {
		return Pow(d, payload)
	}

	if height == 0 {
		return D(payload)
	}

	if d.Eq(dOne) {
		return dFC_NN(1, 0, 1)
	}

	if d.Eq(dNegOne) {
		return d.Pow(payload)
	}

	if height == math.Inf(1) {
		thisNum := d.ToFloat64()
		// within the convergence range?
		if thisNum <= 1.44466786100976613366 && thisNum >= 0.06598803584531253708 {
			negLn := d.Ln().Neg()
			// For bases above 1, b^x = x has two solutions. The lower solution is a stable equilibrium, the upper solution is an unstable equilibrium.
			lower := negLn.LambertW(true).Divide(negLn)
			// However, if the base is below 1, there's only the stable equilibrium solution.
			if thisNum < 1 {
				return lower
			}
			upper := negLn.LambertW(false).Divide(negLn)
			// hotfix for the very edge of the number range not being handled properly
			if thisNum > 1.444667861009099 {
				lower = upper // Assuming D() can take a float64 and convert it to *Decimal
				upper = D(math.E)
			}
			payload = D(payload) // Ensure payload is a *Decimal, adjust as necessary
			if payload.Eq(upper) {
				return upper
			} else if payload.Lt(upper) {
				return lower
			} else {
				return dFC(1, math.Inf(1), math.Inf(1))
			}
		} else if thisNum > 1.44466786100976613366 {
			// explodes to infinity
			return dFC(1, math.Inf(1), math.Inf(1))
		} else {
			// 0.06598803584531253708 > thisNum >= 0: never converges
			// thisNum < 0: quickly becomes a complex number
			return dFC(math.NaN(), math.NaN(), math.NaN())
		}
	}

	if d.Eq(dZero) {
		result := math.Abs(math.Mod((height + 1), 2))
		if result > 1 {
			result = 2 - result
		}
		return decimalFromFloat64(result)
	}

	if height < 0 {
		return IteratedLog(payload, d, -height, false)
	}

	oldHeight := height
	height = math.Trunc(height)
	fracHeight := oldHeight - height

	if d.Gt(dZero) && (d.Lt(D(1)) || (d.Lte(D(1.44466786100976613366)) && payload.Lte(Ln(d).Neg().LambertW(false).Divide(Ln(d).Neg())))) && (oldHeight > 10000 || !linear) {
		limitHeight := math.Min(10000, height)
		if payload.Eq(dOne) {
			payload = d.Pow(D(fracHeight))
		} else if d.Lt(D(1)) {
			payload = payload.Pow(D(1 - fracHeight)).Multiply(d.Pow(payload).Pow(D(fracHeight)))
		} else {
			payload = payload.LayerAdd(D(fracHeight), d, false)
		}

		for i := 0; i < int(limitHeight); i++ {
			oldPayload := payload
			payload := d.Pow(payload)
			if oldPayload.Eq(payload) {
				return payload
			}
		}

		if oldHeight > 10000 && math.Mod(math.Ceil(oldHeight), 2) == 1 {
			return d.Pow(payload)
		}

		return payload
	}

	if fracHeight != 0 {
		if payload.Eq(dOne) {
			if d.Gt(D(10)) || linear {
				payload = d.Pow(D(fracHeight))
			} else {
				payload = decimalFromFloat64(tetrateCritical(d.ToFloat64(), fracHeight))
				if d.Lt(D(2)) {
					payload = payload.Subtract(D(1)).Multiply(d.Subtract(D(1))).Add(D(1))
				}
			}
		} else {
			if d.Eq(D(10)) {
				payload = payload.LayerAdd10(D(fracHeight), linear)
			} else if d.Lt(D(1)) {
				payload = payload.Pow(D(1 - fracHeight)).Multiply(d.Pow(payload).Pow(D(fracHeight)))
			} else {
				payload = payload.LayerAdd(D(fracHeight), d, linear)
			}
		}
	}

	for i := 0; i < int(height); i++ {
		payload = d.Pow(payload)
		if math.IsInf(payload.layer, 0) || math.IsInf(payload.mag, 0) {
			return payload.Normalize()
		}
		if payload.layer-d.layer > 3 {
			return dFC_NN(payload.sign, payload.layer+(height-float64(i)-1), payload.mag)
		}
		if i > 10000 {
			return payload
		}
	}

	return payload
}

// Pentate is the result of tetrating 'height' times in a row
func Pentate[DS DecimalSource](value DS, height float64, payload DS, linear bool) *Decimal {
	return D(value).Pentate(height, D(payload), linear)
}

// Pentate is the result of tetrating 'height' times in a row
func (d *Decimal) Pentate(height float64, payload *Decimal, linear bool) *Decimal {
	oldHeight := height
	height = math.Trunc(height)
	fracHeight := oldHeight - height

	if fracHeight != 0 {
		if payload.Eq(dOne) {
			height++
			payload = decimalFromFloat64(fracHeight)
		} else {
			if d.Eq(D(10)) {
				payload = payload.LayerAdd10(D(fracHeight), linear)
			} else {
				payload = payload.LayerAdd(D(fracHeight), d, linear)
			}
		}
	}

	for i := 0; i < int(height); i++ {
		payload = d.Tetrate(payload.ToFloat64(), dOne, linear)
		if math.IsInf(payload.layer, 0) || math.IsInf(payload.mag, 0) {
			return payload.Normalize()
		}
		if i > 10 {
			return payload
		}
	}

	return payload
}
