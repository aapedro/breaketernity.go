package breaketernitygo

import (
	"math"
)

func Cmp(d *Decimal, other *Decimal) int {
	return d.Cmp(other)
}
func (d *Decimal) Cmp(other *Decimal) int {
	if d.sign > other.sign {
		return 1
	}
	if d.sign < other.sign {
		return -1
	}
	return int(d.sign) * d.CmpAbs(other)
}

func CmpAbs(d *Decimal, other *Decimal) int {
	return d.CmpAbs(other)
}
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

func Eq(d *Decimal, other *Decimal) bool {
	return d.Eq(other)
}
func (d *Decimal) Eq(other *Decimal) bool {
	return d.sign == other.sign && d.mag == other.mag && d.layer == other.layer
}

func Neq(d *Decimal, other *Decimal) bool {
	return d.Neq(other)
}
func (d *Decimal) Neq(other *Decimal) bool {
	return !d.Eq(other)
}

func Lt(d *Decimal, other *Decimal) bool {
	return d.Lt(other)
}
func (d *Decimal) Lt(other *Decimal) bool {
	return d.Cmp(other) == -1
}

func Lte(d *Decimal, other *Decimal) bool {
	return d.Lte(other)
}
func (d *Decimal) Lte(other *Decimal) bool {
	return !d.Gt(other)
}

func Gt(d *Decimal, other *Decimal) bool {
	return d.Gt(other)
}
func (d *Decimal) Gt(other *Decimal) bool {
	return d.Cmp(other) == 1
}

func Gte(d *Decimal, other *Decimal) bool {
	return d.Gte(other)
}
func (d *Decimal) Gte(other *Decimal) bool {
	return !d.Lt(other)
}

func Max(d *Decimal, other *Decimal) *Decimal {
	return d.Max(other)
}
func (d *Decimal) Max(other *Decimal) *Decimal {
	if d.Lt(other) {
		return d
	} else {
		return other
	}
}

func Min(d *Decimal, other *Decimal) *Decimal {
	return d.Min(other)
}
func (d *Decimal) Min(other *Decimal) *Decimal {
	if d.Gt(other) {
		return d
	} else {
		return other
	}
}

func MaxAbs(d *Decimal, other *Decimal) *Decimal {
	return d.MaxAbs(other)
}
func (d *Decimal) MaxAbs(other *Decimal) *Decimal {
	if d.CmpAbs(other) < 0 {
		return d
	} else {
		return other
	}
}

func MinAbs(d *Decimal, other *Decimal) *Decimal {
	return d.MinAbs(other)
}
func (d *Decimal) MinAbs(other *Decimal) *Decimal {
	if d.CmpAbs(other) > 0 {
		return d
	} else {
		return other
	}
}

func Clamp(d *Decimal, min *Decimal, max *Decimal) *Decimal {
	return d.Clamp(min, max)
}
func (d *Decimal) Clamp(min *Decimal, max *Decimal) *Decimal {
	return d.Max(min).Min(max)
}

func ClampMin(d *Decimal, min *Decimal) *Decimal {
	return d.Max(min)
}
func (d *Decimal) ClampMin(min *Decimal) *Decimal {
	return d.Max(min)
}

func ClampMax(d *Decimal, max *Decimal) *Decimal {
	return d.Min(max)
}
func (d *Decimal) ClampMax(max *Decimal) *Decimal {
	return d.Min(max)
}

func EqTolerance(d *Decimal, other *Decimal, tolerance float64) bool {
	return d.EqTolerance(other, tolerance)
}
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

func CmpTolerance(d *Decimal, other *Decimal, tolerance float64) int {
	return d.CmpTolerance(other, tolerance)
}
func (d *Decimal) CmpTolerance(other *Decimal, tolerance float64) int {
	if d.EqTolerance(other, tolerance) {
		return 0
	} else {
		return d.Cmp(other)
	}
}

func NeqTolerance(d *Decimal, other *Decimal, tolerance float64) bool {
	return d.NeqTolerance(other, tolerance)
}
func (d *Decimal) NeqTolerance(other *Decimal, tolerance float64) bool {
	return !d.EqTolerance(other, tolerance)
}

func LtTolerance(d *Decimal, other *Decimal, tolerance float64) bool {
	return d.LtTolerance(other, tolerance)
}
func (d *Decimal) LtTolerance(other *Decimal, tolerance float64) bool {
	return !d.EqTolerance(other, tolerance) && d.Lt(other)
}

func LteTolerance(d *Decimal, other *Decimal, tolerance float64) bool {
	return d.LteTolerance(other, tolerance)
}
func (d *Decimal) LteTolerance(other *Decimal, tolerance float64) bool {
	return d.EqTolerance(other, tolerance) || d.Lt(other)
}

func GtTolerance(d *Decimal, other *Decimal, tolerance float64) bool {
	return d.GtTolerance(other, tolerance)
}
func (d *Decimal) GtTolerance(other *Decimal, tolerance float64) bool {
	return !d.EqTolerance(other, tolerance) && d.Gt(other)
}

func GteTolerance(d *Decimal, other *Decimal, tolerance float64) bool {
	return d.GteTolerance(other, tolerance)
}
func (d *Decimal) GteTolerance(other *Decimal, tolerance float64) bool {
	return d.EqTolerance(other, tolerance) || d.Gt(other)
}

func PLog10(d *Decimal) *Decimal {
	return d.PLog10()
}
func (d *Decimal) PLog10() *Decimal {
	if d.Lt(dZero) {
		return dFC_NN(0, 0, 0)
	}
	return d.Log10()
}

func AbsLog10(d *Decimal) *Decimal {
	return d.AbsLog10()
}
func (d *Decimal) AbsLog10() *Decimal {
	if d.sign == 0 {
		return dFC_NN(math.NaN(), math.NaN(), math.NaN())
	} else if d.layer > 0 {
		return dFC(sign(d.mag), d.layer-1, math.Abs(d.mag))
	} else {
		return dFC(1, 0, math.Log10(d.mag))
	}
}

func Log10(d *Decimal) *Decimal {
	return d.Log10()
}
func (d *Decimal) Log10() *Decimal {
	if d.sign <= 0 {
		return dFC_NN(math.NaN(), math.NaN(), math.NaN())
	} else if d.layer > 0 {
		return dFC(sign(d.mag), d.layer-1, math.Abs(d.mag))
	} else {
		return dFC(d.sign, 0, math.Log10(d.mag))
	}
}

func Log(d *Decimal, base *Decimal) *Decimal {
	return d.Log(base)
}
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

func Ln(d *Decimal) *Decimal {
	return d.Ln()
}
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

func Log2(d *Decimal) *Decimal {
	return d.Log2()
}
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

func Pow(d *Decimal, other *Decimal) *Decimal {
	return d.Pow(other)
}
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

func PowBase10(d *Decimal) *Decimal {
	return d.PowBase10()
}
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

func PowBaseE(d *Decimal) *Decimal {
	return d.PowBaseE()
}
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

func PowBase(d *Decimal, base *Decimal) *Decimal {
	return d.PowBase(base)
}
func (d *Decimal) PowBase(base *Decimal) *Decimal {
	return D(d).Pow(base)
}

func Root(d *Decimal, degree *Decimal) *Decimal {
	return d.Root(degree)
}
func (d *Decimal) Root(degree *Decimal) *Decimal {
	return d.Pow(degree.Recip())
}

func Sqrt(d *Decimal) *Decimal {
	return d.Sqrt()
}
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

func Factorial(d *Decimal) *Decimal {
	return d.Factorial()
}
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

func Gamma(d *Decimal) *Decimal {
	return d.Gamma()
}
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

func Abs(d *Decimal) *Decimal {
	return d.Abs()
}
func (d *Decimal) Abs() *Decimal {
	if d.sign == 0 {
		return dFC_NN(0, d.layer, d.mag)
	} else {
		return dFC_NN(1, d.layer, d.mag)
	}
}

func Neg(d *Decimal) *Decimal {
	return d.Neg()
}
func (d *Decimal) Neg() *Decimal {
	return dFC_NN(sign(-1*d.mag), d.layer, d.mag)
}

func Round(d *Decimal) *Decimal {
	return d.Round()
}
func (d *Decimal) Round() *Decimal {
	if d.mag < 0 {
		return dFC_NN(0, 0, 0)
	}
	if d.layer == 0 {
		return dFC(d.sign, 0, math.Round(d.mag))
	}
	return D(d)
}

func Floor(d *Decimal) *Decimal {
	return d.Floor()
}
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

func Ceil(d *Decimal) *Decimal {
	return d.Ceil()
}
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

func Trunc(d *Decimal) *Decimal {
	return d.Trunc()
}
func (d *Decimal) Trunc() *Decimal {
	if d.mag < 0 {
		return dFC_NN(0, 0, 0)
	}
	if d.layer == 0 {
		return dFC(d.sign, 0, math.Trunc(d.mag))
	}
	return D(d)
}

func Recip(d *Decimal) *Decimal {
	return d.Recip()
}
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

func Add(d *Decimal, other *Decimal) *Decimal {
	return d.Add(other)
}
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

func Subtract(d *Decimal, other *Decimal) *Decimal {
	return d.Subtract(other)
}
func (d *Decimal) Subtract(other *Decimal) *Decimal {
	return d.Add(D(other).Neg())
}

func Multiply(d *Decimal, other *Decimal) *Decimal {
	return d.Multiply(other)
}
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

	// TODO: fix this?
	panic("Bad arguments to multiply: " + d.ToString() + ", " + other.ToString())
}

func Divide(d *Decimal, other *Decimal) *Decimal {
	return d.Divide(other)
}
func (d *Decimal) Divide(other *Decimal) *Decimal {
	return d.Multiply(other.Recip())
}

func Modulo(d *Decimal, other *Decimal) *Decimal {
	return d.Modulo(other)
}
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

func IsNaN(d *Decimal) bool {
	return d.IsNaN()
}
func (d *Decimal) IsNaN() bool {
	return math.IsNaN(d.layer) || math.IsNaN(d.mag) || math.IsNaN(d.sign)
}

func isInf(d *Decimal) bool {
	return d.IsInf()
}
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
func LambertW(d *Decimal, principal bool) *Decimal {
	return d.LambertW(principal)
}
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

func IteratedLog(d *Decimal, base *Decimal, times float64, linear bool) *Decimal {
	return d.IteratedLog(base, times, linear)
}
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

func LayerAdd10(d *Decimal, diff *Decimal, linear bool) *Decimal {
	return d.LayerAdd10(diff, linear)
}
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

func LayerAdd(d *Decimal, diff *Decimal, base *Decimal, linear bool) *Decimal {
	return d.LayerAdd(diff, base, linear)
}
func (d *Decimal) LayerAdd(diff *Decimal, base *Decimal, linear bool) *Decimal {
	fDiff := diff.ToFloat64()
	if base.Gt(D(1)) && base.Lte(D(1.44466786100976613366)) {
		excessSlog, e1 := ExcessSlog(d, base, linear)
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

func ExcessSlog(d *Decimal, base *Decimal, linear bool) (*Decimal, int) {
	return d.ExcessSlog(base, linear)
}
func (d *Decimal) ExcessSlog(base *Decimal, linear bool) (*Decimal, int) {
	// TODO
	return d, 0
}

func slogInternal(d *Decimal, base *Decimal, linear bool) *Decimal {
	return d.slogInternal(base, linear)
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

func Slog(d *Decimal, base *Decimal, iterations float64, linear bool) *Decimal {
	return d.Slog(base, iterations, linear)
}
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

func Pentate(value *Decimal, height float64, payload *Decimal, linear bool) *Decimal {
	// TODO: not implemented yet
	return value.Pentate(height, payload, linear)
}
func (d *Decimal) Pentate(height float64, payload *Decimal, linear bool) *Decimal {
	// TODO: not implemented yet
	return d
}

func Tetrate(value *Decimal, height float64, payload *Decimal, linear bool) *Decimal {
	return value.Tetrate(height, payload, linear)
}
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