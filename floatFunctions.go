package breaketernitygo

import "math"

func fMagLog10(x float64) float64 {
	return math.Copysign(1, x) * math.Log10(math.Abs(x))
}

func fGamma(n float64) float64 {
	if math.IsInf(n, 0) {
		return n
	}
	if n < -50 {
		if n == math.Trunc(n) {
			return math.Inf(-1)
		}
		return 0
	}

	scal1 := 1.0
	for n < 10 {
		scal1 *= n
		n++
	}

	n -= 1
	l := 0.9189385332046727 // 0.5 * math.Log(2 * math.Pi)
	l += (n + 0.5) * math.Log(n)
	l -= n
	n2 := n * n
	np := n
	l += 1 / (12 * np)
	np *= n2
	l -= 1 / (360 * np)
	np *= n2
	l += 1 / (1260 * np)
	np *= n2
	l -= 1 / (1680 * np)
	np *= n2
	l += 1 / (1188 * np)
	np *= n2
	l -= 691 / (360360 * np)
	np *= n2
	l += 7 / (1092 * np)
	np *= n2
	l -= 3617 / (122400 * np)

	return math.Exp(l) / scal1
}

const EXPN1 = 0.36787944117144232159553 // exp(-1)
const OMEGA = 0.56714329040978387299997 // W(1, 0)
func fLambertW(z float64, tol float64, principal bool) float64 {
	var w, wn float64

	if math.IsInf(z, 0) {
		return z
	}
	if principal {
		if z == 0 {
			return z
		}
		if z == 1 {
			return OMEGA
		}

		if z < 10 {
			w = 0
		} else {
			w = math.Log(z) - math.Log(math.Log(z))
		}
	} else {
		if z == 0 {
			return math.Inf(-1)
		}

		if z <= -0.1 {
			w = -2
		} else {
			w = math.Log(-z) - math.Log(-math.Log(-z))
		}
	}

	for i := 0; i < 100; i++ {
		wn = (z*math.Exp(-w) + w*w) / (w + 1)
		if math.Abs(wn-w) < tol*math.Abs(wn) {
			return wn
		} else {
			w = wn
		}
	}

	panic("Iteration failed to converge")
}
